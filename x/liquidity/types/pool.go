package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (pool Pool) GetReserveAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(pool.ReserveAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

type PoolI interface {
	InitialPoolCoinSupply() sdk.Int
	PoolCoinSupply() sdk.Int
	ReserveBalance() (x, y sdk.Int)
	SwapRequests() []SwapRequest
}

type PoolOperations struct {
	Pool PoolI
}

func (ops PoolOperations) IsDepleted() bool {
	pc := ops.Pool.PoolCoinSupply()
	if pc.IsZero() {
		return true
	}
	rx, ry := ops.Pool.ReserveBalance()
	if rx.IsZero() || ry.IsZero() {
		return true
	}
	return false
}

func (ops PoolOperations) PoolPrice() sdk.Dec {
	if ops.IsDepleted() {
		return sdk.ZeroDec()
	}
	rx, ry := ops.Pool.ReserveBalance()
	return rx.ToDec().QuoInt(ry)
}

// Deposit returns accepted x amount, accepted y amount and
// minted pool coin amount.
func (ops PoolOperations) Deposit(x, y sdk.Int) (ax, ay, pc sdk.Int) {
	// If the pool is depleted, accept all coins and mint
	// pool coins as much as the initial pool coin supply.
	if ops.IsDepleted() {
		ax = x
		ay = y
		pc = ops.Pool.InitialPoolCoinSupply()
		return
	}

	// Calculate accepted amount and minting amount.
	// Note that we take as many coins as possible(by ceiling numbers)
	// from depositor and mint as little coins as possible.
	ax = x
	ay = y

	rx, ry := ops.Pool.ReserveBalance()
	cp := rx.ToDec().QuoInt(ry) // current pool price
	dp := x.ToDec().QuoInt(y)   // price of coins that are being deposited

	switch {
	case cp.LT(dp):
		ax = y.ToDec().Mul(cp).Ceil().TruncateInt()
	case cp.GT(dp):
		ay = x.ToDec().Quo(cp).Ceil().TruncateInt()
	}

	ps := ops.Pool.PoolCoinSupply().ToDec()
	pc = sdk.MinInt(
		ps.Mul(ax.ToDec()).QuoTruncate(rx.ToDec()).RoundInt(),
		ps.Mul(ay.ToDec()).QuoTruncate(ry.ToDec()).RoundInt(),
	)

	return
}

func (ops PoolOperations) Withdraw(pc sdk.Int) (x, y sdk.Int) {
	// TODO: implement
	return
}

func (ops PoolOperations) OrderBook() OrderBook {
	var orderBook OrderBook

	for _, req := range ops.Pool.SwapRequests() {
		orderBook.Add(Order{
			Orderer:   req.Requester,
			Direction: req.Direction,
			Amount:    req.RemainingAmount,
			Price:     req.Price,
		})
	}

	return orderBook
}

func (ops PoolOperations) Match(orderBook OrderBook) {
	p := ops.PoolPrice()
	_ = p
	// TODO: implement
}
