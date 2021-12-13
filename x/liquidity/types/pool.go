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

func (ops PoolOperations) Deposit(x, y sdk.Int) (pc sdk.Int) {
	// TODO: implement
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
