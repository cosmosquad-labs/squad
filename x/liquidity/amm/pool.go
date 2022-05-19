package amm

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/cosmosquad-labs/squad/types"
)

var (
	_ Pool        = (*BasicPool)(nil)
	_ Pool        = (*RangedPool)(nil)
	_ OrderSource = (*MockPoolOrderSource)(nil)
)

// Pool is the interface of a pool.
// It also satisfies OrderView interface.
type Pool interface {
	OrderView
	Balances() (rx, ry sdk.Int)
	PoolCoinSupply() sdk.Int
	Price() sdk.Dec
	IsDepleted() bool
	Deposit(x, y sdk.Int) (ax, ay, pc sdk.Int)
	Withdraw(pc sdk.Int, feeRate sdk.Dec) (x, y sdk.Int)
	ProvidableXAmountOver(price sdk.Dec) sdk.Int
	ProvidableYAmountUnder(price sdk.Dec) sdk.Int
}

// BasicPool is the basic pool type.
type BasicPool struct {
	// rx and ry are the pool's reserve balance of each x/y coin.
	// In perspective of a pair, x coin is the quote coin and
	// y coin is the base coin.
	rx, ry sdk.Int
	// ps is the pool's pool coin supply.
	ps sdk.Int
}

// NewBasicPool returns a new BasicPool.
// It is OK to pass an empty sdk.Int to ps when ps is not going to be used.
func NewBasicPool(rx, ry, ps sdk.Int) *BasicPool {
	return &BasicPool{
		rx: rx,
		ry: ry,
		ps: ps,
	}
}

// Balances returns the balances of the pool.
func (pool *BasicPool) Balances() (rx, ry sdk.Int) {
	return pool.rx, pool.ry
}

// PoolCoinSupply returns the pool coin supply.
func (pool *BasicPool) PoolCoinSupply() sdk.Int {
	return pool.ps
}

// Price returns the pool price.
func (pool *BasicPool) Price() sdk.Dec {
	if pool.rx.IsZero() || pool.ry.IsZero() {
		panic("pool price is not defined for a depleted pool")
	}
	return pool.rx.ToDec().Quo(pool.ry.ToDec())
}

// IsDepleted returns whether the pool is depleted or not.
func (pool *BasicPool) IsDepleted() bool {
	return pool.ps.IsZero() || pool.rx.IsZero() || pool.ry.IsZero()
}

// Deposit returns accepted x and y coin amount and minted pool coin amount
// when someone deposits x and y coins.
func (pool *BasicPool) Deposit(x, y sdk.Int) (ax, ay, pc sdk.Int) {
	// Calculate accepted amount and minting amount.
	// Note that we take as many coins as possible(by ceiling numbers)
	// from depositor and mint as little coins as possible.

	utils.SafeMath(func() {
		rx, ry := pool.rx.ToDec(), pool.ry.ToDec()
		ps := pool.ps.ToDec()

		// pc = floor(ps * min(x / rx, y / ry))
		pc = ps.MulTruncate(sdk.MinDec(
			x.ToDec().QuoTruncate(rx),
			y.ToDec().QuoTruncate(ry),
		)).TruncateInt()

		mintProportion := pc.ToDec().Quo(ps)             // pc / ps
		ax = rx.Mul(mintProportion).Ceil().TruncateInt() // ceil(rx * mintProportion)
		ay = ry.Mul(mintProportion).Ceil().TruncateInt() // ceil(ry * mintProportion)
	}, func() {
		ax, ay, pc = sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()
	})

	return
}

// Withdraw returns withdrawn x and y coin amount when someone withdraws
// pc pool coin.
// Withdraw also takes care of the fee rate.
func (pool *BasicPool) Withdraw(pc sdk.Int, feeRate sdk.Dec) (x, y sdk.Int) {
	if pc.Equal(pool.ps) {
		// Redeeming the last pool coin - give all remaining rx and ry.
		x = pool.rx
		y = pool.ry
		return
	}

	utils.SafeMath(func() {
		proportion := pc.ToDec().QuoTruncate(pool.ps.ToDec())                             // pc / ps
		multiplier := sdk.OneDec().Sub(feeRate)                                           // 1 - feeRate
		x = pool.rx.ToDec().MulTruncate(proportion).MulTruncate(multiplier).TruncateInt() // floor(rx * proportion * multiplier)
		y = pool.ry.ToDec().MulTruncate(proportion).MulTruncate(multiplier).TruncateInt() // floor(ry * proportion * multiplier)
	}, func() {
		x, y = sdk.ZeroInt(), sdk.ZeroInt()
	})

	return
}

// HighestBuyPrice returns the highest buy price of the pool.
func (pool *BasicPool) HighestBuyPrice() (price sdk.Dec, found bool) {
	// The highest buy price is actually a bit lower than pool price,
	// but it's not important for our matching logic.
	return pool.Price(), true
}

// LowestSellPrice returns the lowest sell price of the pool.
func (pool *BasicPool) LowestSellPrice() (price sdk.Dec, found bool) {
	// The lowest sell price is actually a bit higher than the pool price,
	// but it's not important for our matching logic.
	return pool.Price(), true
}

// BuyAmountOver returns the amount of buy orders for price greater than
// or equal to given price.
func (pool *BasicPool) BuyAmountOver(price sdk.Dec) (amt sdk.Int) {
	if price.GTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	return pool.rx.ToDec().QuoTruncate(price).Sub(pool.ry.ToDec()).TruncateInt()
}

// SellAmountUnder returns the amount of sell orders for price less than
// or equal to given price.
func (pool *BasicPool) SellAmountUnder(price sdk.Dec) sdk.Int {
	return pool.ProvidableYAmountUnder(price)
}

// ProvidableXAmountOver returns the amount of x coin the pool would provide
// for price greater than or equal to given price.
func (pool *BasicPool) ProvidableXAmountOver(price sdk.Dec) (amt sdk.Int) {
	if price.GTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	return pool.rx.ToDec().Sub(pool.ry.ToDec().Mul(price)).TruncateInt()
}

// ProvidableYAmountUnder returns the amount of y coin the pool would provide
// for price less than or equal to given price.
func (pool *BasicPool) ProvidableYAmountUnder(price sdk.Dec) sdk.Int {
	if price.LTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	return pool.ry.ToDec().Sub(pool.rx.ToDec().QuoRoundUp(price)).TruncateInt()
}

type RangedPool struct {
	rx, ry sdk.Int
	ps     sdk.Int
	m, l   *sdk.Dec
}

// NewRangedPool returns a new RangedPool.
func NewRangedPool(rx, ry, ps sdk.Int, m, l *sdk.Dec) *RangedPool {
	return &RangedPool{
		rx: rx,
		ry: ry,
		ps: ps,
		m:  m,
		l:  l,
	}
}

// CreateRangedPool creates new RangedPool from given inputs, while validating
// the inputs and using only needed amount of x/y coins(the rest should be refunded).
func CreateRangedPool(x, y sdk.Int, initialPrice sdk.Dec, minPrice, maxPrice *sdk.Dec) (pool *RangedPool, err error) {
	if !x.IsPositive() && !y.IsPositive() {
		return nil, fmt.Errorf("either x or y must be positive")
	}
	if initialPrice.IsZero() {
		return nil, fmt.Errorf("initial price must not be 0")
	}
	if minPrice == nil && maxPrice == nil {
		return nil, fmt.Errorf("min price and max price must not be nil at the same time")
	}
	if minPrice != nil && maxPrice != nil && minPrice.GTE(*maxPrice) {
		return nil, fmt.Errorf("max price must be greater than min price")
	}

	switch {
	case minPrice != nil && initialPrice.LTE(*minPrice):
		// single y asset pool
		if y.IsZero() {
			return nil, fmt.Errorf("y coin amount must not be 0 for single asset pool")
		}
		ax, ay := sdk.ZeroInt(), y
		return NewRangedPool(ax, ay, InitialPoolCoinSupply(ax, ay), minPrice, maxPrice), nil
	case maxPrice != nil && initialPrice.GTE(*maxPrice):
		// single x asset pool
		if x.IsZero() {
			return nil, fmt.Errorf("x coin amount must not be 0 for single asset pool")
		}
		ax, ay := x, sdk.ZeroInt()
		return NewRangedPool(ax, ay, InitialPoolCoinSupply(ax, ay), minPrice, maxPrice), nil
	}

	m := sdk.ZeroDec()
	if minPrice != nil {
		m = *minPrice
	}

	var ax, ay sdk.Int
	if maxPrice != nil {
		l := *maxPrice
		r := l.Sub(initialPrice).Quo(l.Mul(initialPrice.Sub(m))) // r = (l-p)/(l*(p-m)))

		ay = x.ToDec().Mul(r).Ceil().TruncateInt() // ay = ceil(x * r)
		if ay.GT(y) {
			ax = y.ToDec().Quo(r).Ceil().TruncateInt() // ax = ceil(y / r)
			ay = y
		} else {
			ax = x
		}
	} else {
		r := initialPrice.Sub(m) // r = p - m

		ay = x.ToDec().QuoRoundUp(r).Ceil().TruncateInt() // ay = ceil(x / r)
		if ay.GT(y) {
			ax = y.ToDec().Mul(r).Ceil().TruncateInt() // ax = ceil(y * r)
			ay = y
		} else {
			ax = x
		}
	}

	return NewRangedPool(ax, ay, InitialPoolCoinSupply(ax, ay), minPrice, maxPrice), nil
}

// Balances returns the balances of the pool.
func (pool *RangedPool) Balances() (rx, ry sdk.Int) {
	return pool.rx, pool.ry
}

// PoolCoinSupply returns the pool coin supply.
func (pool *RangedPool) PoolCoinSupply() sdk.Int {
	return pool.ps
}

// Price returns the pool price.
func (pool *RangedPool) Price() sdk.Dec {
	a, b := pool.ab()
	if pool.rx.ToDec().Add(a).IsZero() || pool.ry.ToDec().Add(b).IsZero() {
		panic("pool price is not defined for a depleted pool")
	}
	return pool.rx.ToDec().Add(a).Quo(pool.ry.ToDec().Add(b)) // (rx + a) / (ry + b)
}

// IsDepleted returns whether the pool is depleted or not.
func (pool *RangedPool) IsDepleted() bool {
	a, b := pool.ab()
	return pool.ps.IsZero() || pool.rx.ToDec().Add(a).IsZero() || pool.ry.ToDec().Add(b).IsZero()
}

// Deposit returns accepted x and y coin amount and minted pool coin amount
// when someone deposits x and y coins.
// TODO: refactor Deposit to a separate function
func (pool *RangedPool) Deposit(x, y sdk.Int) (ax, ay, pc sdk.Int) {
	// Calculate accepted amount and minting amount.
	// Note that we take as many coins as possible(by ceiling numbers)
	// from depositor and mint as little coins as possible.

	utils.SafeMath(func() {
		rx, ry := pool.rx.ToDec(), pool.ry.ToDec()
		ps := pool.ps.ToDec()

		// pc = floor(ps * min(x / rx, y / ry))
		pc = ps.MulTruncate(sdk.MinDec(
			x.ToDec().QuoTruncate(rx),
			y.ToDec().QuoTruncate(ry),
		)).TruncateInt()

		mintProportion := pc.ToDec().Quo(ps)             // pc / ps
		ax = rx.Mul(mintProportion).Ceil().TruncateInt() // ceil(rx * mintProportion)
		ay = ry.Mul(mintProportion).Ceil().TruncateInt() // ceil(ry * mintProportion)
	}, func() {
		ax, ay, pc = sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()
	})

	return
}

// Withdraw returns withdrawn x and y coin amount when someone withdraws
// pc pool coin.
// Withdraw also takes care of the fee rate.
// TODO: refactor Withdraw to a separate function
func (pool *RangedPool) Withdraw(pc sdk.Int, feeRate sdk.Dec) (x, y sdk.Int) {
	if pc.Equal(pool.ps) {
		// Redeeming the last pool coin - give all remaining rx and ry.
		x = pool.rx
		y = pool.ry
		return
	}

	utils.SafeMath(func() {
		proportion := pc.ToDec().QuoTruncate(pool.ps.ToDec())                             // pc / ps
		multiplier := sdk.OneDec().Sub(feeRate)                                           // 1 - feeRate
		x = pool.rx.ToDec().MulTruncate(proportion).MulTruncate(multiplier).TruncateInt() // floor(rx * proportion * multiplier)
		y = pool.ry.ToDec().MulTruncate(proportion).MulTruncate(multiplier).TruncateInt() // floor(ry * proportion * multiplier)
	}, func() {
		x, y = sdk.ZeroInt(), sdk.ZeroInt()
	})

	return
}

// HighestBuyPrice returns the highest buy price of the pool.
func (pool *RangedPool) HighestBuyPrice() (price sdk.Dec, found bool) {
	// The highest buy price is actually a bit lower than pool price,
	// but it's not important for our matching logic.
	return pool.Price(), true
}

// LowestSellPrice returns the lowest sell price of the pool.
func (pool *RangedPool) LowestSellPrice() (price sdk.Dec, found bool) {
	// The lowest sell price is actually a bit higher than the pool price,
	// but it's not important for our matching logic.
	return pool.Price(), true
}

// BuyAmountOver returns the amount of buy orders for price greater than
// or equal to given price.
func (pool *RangedPool) BuyAmountOver(price sdk.Dec) (amt sdk.Int) {
	if price.GTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	if pool.m != nil && price.LT(*pool.m) {
		price = *pool.m
	}
	a, b := pool.ab()
	utils.SafeMath(func() {
		amt = pool.rx.ToDec().Add(a).QuoTruncate(price).Sub(pool.ry.ToDec().Add(b)).TruncateInt()
		if amt.GT(MaxCoinAmount) {
			amt = MaxCoinAmount
		}
	}, func() {
		amt = MaxCoinAmount
	})
	return
}

// SellAmountUnder returns the amount of sell orders for price less than
// or equal to given price.
func (pool *RangedPool) SellAmountUnder(price sdk.Dec) sdk.Int {
	return pool.ProvidableYAmountUnder(price)
}

// ProvidableXAmountOver returns the amount of x coin the pool would provide
// for price greater than or equal to given price.
func (pool *RangedPool) ProvidableXAmountOver(price sdk.Dec) (amt sdk.Int) {
	if price.GTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	if pool.m != nil && price.LT(*pool.m) {
		price = *pool.m
	}
	a, b := pool.ab()
	return pool.rx.ToDec().Add(a).Sub(price.Mul(pool.ry.ToDec().Add(b))).TruncateInt()
}

// ProvidableYAmountUnder returns the amount of y coin the pool would provide
// for price less than or equal to given price.
func (pool *RangedPool) ProvidableYAmountUnder(price sdk.Dec) sdk.Int {
	if price.LTE(pool.Price()) {
		return sdk.ZeroInt()
	}
	if pool.l != nil && price.GT(*pool.l) {
		price = *pool.l
	}
	a, b := pool.ab()
	return pool.ry.ToDec().Add(b).Sub(pool.rx.ToDec().Add(a).QuoRoundUp(price)).TruncateInt()
}

func (pool *RangedPool) ab() (a, b sdk.Dec) {
	m := sdk.ZeroDec()
	if pool.m != nil {
		m = *pool.m
	}
	if pool.l != nil {
		// a = m * (rx + l * ry) / (l - m)
		// b = (rx + m * ry) / (l - m)
		a = m.Mul(pool.rx.ToDec().Add(pool.l.MulInt(pool.ry))).Quo(pool.l.Sub(m))
		b = pool.rx.ToDec().Add(m.MulInt(pool.ry)).Quo(pool.l.Sub(m))
	} else {
		// a = m * ry
		// b = 0
		a = m.MulInt(pool.ry)
		b = sdk.ZeroDec()
	}
	return
}

// PoolsOrderBook returns an order book with orders made by pools.
// Use Ticks or EvenTicks to generate ticks where pools put orders.
// ticks should have more than 1 element.
// The orders in the order book are just mocks, so the rest fields
// other than Direction, Price and Amount of BaseOrder will have no
// special meaning.
func PoolsOrderBook(pools []Pool, ticks []sdk.Dec) *OrderBook {
	sortTicks(ticks)
	highestTick := ticks[0]
	lowestTick := ticks[len(ticks)-1]
	gap := ticks[0].Sub(ticks[1])
	ob := NewOrderBook()
	for _, pool := range pools {
		poolPrice := pool.Price()
		if poolPrice.GT(lowestTick) { // Buy orders
			accAmtX := pool.ProvidableXAmountOver(highestTick.Add(gap))
			for _, tick := range ticks {
				amtX := pool.ProvidableXAmountOver(tick).Sub(accAmtX)
				if amtX.IsPositive() {
					amt := amtX.ToDec().QuoTruncate(tick).TruncateInt()
					if amt.IsPositive() {
						ob.Add(NewBaseOrder(
							Buy, tick, amt, sdk.NewCoin("quote", OfferCoinAmount(Buy, tick, amt)), "base"))
					}
					accAmtX = accAmtX.Add(amtX)
				}
			}
		}
		if poolPrice.LT(highestTick) { // Sell orders
			accAmt := pool.SellAmountUnder(lowestTick.Sub(gap))
			for i := len(ticks) - 1; i >= 0; i-- {
				tick := ticks[i]
				amt := pool.SellAmountUnder(tick).Sub(accAmt)
				if amt.IsPositive() {
					ob.Add(NewBaseOrder(Sell, tick, amt, sdk.NewCoin("base", amt), "quote"))
					accAmt = accAmt.Add(amt)
				}
			}
		}
	}
	return ob
}

// InitialPoolCoinSupply returns ideal initial pool coin minting amount.
func InitialPoolCoinSupply(x, y sdk.Int) sdk.Int {
	cx := len(x.BigInt().Text(10)) - 1 // characteristic of x
	cy := len(y.BigInt().Text(10)) - 1 // characteristic of y
	c := ((cx + 1) + (cy + 1) + 1) / 2 // ceil(((cx + 1) + (cy + 1)) / 2)
	res := big.NewInt(10)
	res.Exp(res, big.NewInt(int64(c)), nil) // 10^c
	return sdk.NewIntFromBigInt(res)
}

// MockPoolOrderSource demonstrates how to implement a pool OrderSource.
type MockPoolOrderSource struct {
	Pool
	baseCoinDenom, quoteCoinDenom string
}

// NewMockPoolOrderSource returns a new MockPoolOrderSource for testing.
func NewMockPoolOrderSource(pool Pool, baseCoinDenom, quoteCoinDenom string) *MockPoolOrderSource {
	return &MockPoolOrderSource{
		Pool:           pool,
		baseCoinDenom:  baseCoinDenom,
		quoteCoinDenom: quoteCoinDenom,
	}
}

// BuyOrdersOver returns buy orders for price greater or equal than given price.
func (os *MockPoolOrderSource) BuyOrdersOver(price sdk.Dec) []Order {
	amt := os.BuyAmountOver(price)
	if amt.IsZero() {
		return nil
	}
	quoteCoin := sdk.NewCoin(os.quoteCoinDenom, OfferCoinAmount(Buy, price, amt))
	return []Order{NewBaseOrder(Buy, price, amt, quoteCoin, os.baseCoinDenom)}
}

// SellOrdersUnder returns sell orders for price less or equal than given price.
func (os *MockPoolOrderSource) SellOrdersUnder(price sdk.Dec) []Order {
	amt := os.SellAmountUnder(price)
	if amt.IsZero() {
		return nil
	}
	baseCoin := sdk.NewCoin(os.baseCoinDenom, amt)
	return []Order{NewBaseOrder(Sell, price, amt, baseCoin, os.quoteCoinDenom)}
}
