package types

import (
	"fmt"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ OrderSource = (*OrderBookTicks)(nil)

type PriceDirection int

const (
	PriceIncreasing PriceDirection = iota + 1
	PriceDecreasing
)

type Orders []Order

func (orders Orders) RemainingAmount() sdk.Int {
	amount := sdk.ZeroInt()
	for _, order := range orders {
		amount = amount.Add(order.RemainingAmount())
	}
	return amount
}

type OrderBookTick struct {
	price  sdk.Dec
	orders Orders
}

func NewOrderBookTick(order Order) *OrderBookTick {
	return &OrderBookTick{
		price:  order.Price(),
		orders: Orders{order},
	}
}

type OrderBookTicks []*OrderBookTick

func (ticks OrderBookTicks) FindPrice(price sdk.Dec) (i int, exact bool) {
	var prices []sdk.Dec
	for _, tick := range ticks {
		prices = append(prices, tick.price)
	}
	i = sort.Search(len(prices), func(i int) bool {
		return prices[i].LT(price)
	})
	if i < len(prices) && prices[i].Equal(price) {
		exact = true
	}
	return
}

func (ticks *OrderBookTicks) AddOrder(order Order) {
	i, exact := ticks.FindPrice(order.Price())
	if exact {
		(*ticks)[i].orders = append((*ticks)[i].orders, order)
	} else {
		if i < len(*ticks) {
			// Insert a new order book tick at index i.
			*ticks = append((*ticks)[:i], append([]*OrderBookTick{NewOrderBookTick(order)}, (*ticks)[i:]...)...)
		} else {
			// Append a new order group at the end.
			*ticks = append(*ticks, NewOrderBookTick(order))
		}
	}
}

func (ticks OrderBookTicks) AmountGTE(price sdk.Dec) sdk.Int {
	i, exact := ticks.FindPrice(price)
	if !exact {
		i--
	}
	amount := sdk.ZeroInt()
	for ; i >= 0; i-- {
		amount = amount.Add(ticks[i].orders.RemainingAmount())
	}
	return amount
}

func (ticks OrderBookTicks) AmountLTE(price sdk.Dec) sdk.Int {
	i, _ := ticks.FindPrice(price)
	amount := sdk.ZeroInt()
	for ; i < len(ticks); i++ {
		amount = amount.Add(ticks[i].orders.RemainingAmount())
	}
	return amount
}

func (ticks OrderBookTicks) Orders(price sdk.Dec) []Order {
	i, exact := ticks.FindPrice(price)
	if !exact {
		return nil
	}
	return ticks[i].orders
}

func (ticks OrderBookTicks) UpTick(price sdk.Dec, _ int) (tick sdk.Dec, found bool) {
	i, _ := ticks.FindPrice(price)
	if i == 0 {
		return
	}
	return ticks[i-1].price, true
}

func (ticks OrderBookTicks) DownTick(price sdk.Dec, _ int) (tick sdk.Dec, found bool) {
	i, exact := ticks.FindPrice(price)
	if !exact {
		i--
	}
	if i >= len(ticks)-1 {
		return
	}
	return ticks[i+1].price, true
}

func (ticks OrderBookTicks) HighestTick(_ int) (tick sdk.Dec, found bool) {
	if len(ticks) == 0 {
		return
	}
	return ticks[0].price, true
}

func (ticks OrderBookTicks) LowestTick(_ int) (tick sdk.Dec, found bool) {
	if len(ticks) == 0 {
		return
	}
	return ticks[len(ticks)-1].price, true
}

type OrderBook struct {
	buys  OrderBookTicks
	sells OrderBookTicks
}

func (ob OrderBook) AddOrder(order Order) {
	switch order.Direction() {
	case SwapDirectionBuy:
		ob.buys.AddOrder(order)
	case SwapDirectionSell:
		ob.sells.AddOrder(order)
	}
}

func (ob OrderBook) AddOrders(orders ...Order) {
	for _, order := range orders {
		ob.AddOrder(order)
	}
}

func (ob OrderBook) OrderSource(dir SwapDirection) OrderSource {
	switch dir {
	case SwapDirectionBuy:
		return ob.buys
	case SwapDirectionSell:
		return ob.sells
	default:
		panic(fmt.Sprintf("unknown swap direction: %v", dir))
	}
}

//func (ob OrderBook) String() string {
//	lines := []string{
//		"+-----buy------+----------price-----------+-----sell-----+",
//	}
//	for _, tick := range ob.ticks {
//		lines = append(lines,
//			fmt.Sprintf("| %12s | %24s | %-12s |",
//				tick.buys.RemainingAmount(), tick.price.String(), tick.sells.RemainingAmount()))
//	}
//	lines = append(lines, "+--------------+--------------------------+--------------+")
//	return strings.Join(lines, "\n")
//}

//// DemandingAmount returns total demanding amount of orders at given price.
//// Demanding amount is the amount of coins these orders want to receive.
//// Note that orders should have same SwapDirection, since
//// DemandingAmount doesn't rely on SwapDirection.
//// TODO: use sdk.Dec here?
//func (os Orders) DemandingAmount(price sdk.Dec) sdk.Int {
//	da := sdk.ZeroInt()
//	for _, order := range os {
//		switch order.Direction {
//		case SwapDirectionXToY:
//			da = da.Add(order.RemainingAmount.ToDec().QuoTruncate(price).TruncateInt())
//		case SwapDirectionYToX:
//			da = da.Add(order.RemainingAmount.ToDec().MulTruncate(price).TruncateInt())
//		}
//	}
//	return da
//}
//
//// MatchOrders matches two order groups at given price.
//func MatchOrders(a, b Orders, price sdk.Dec) {
//	amtA := a.RemainingAmount()
//	amtB := b.RemainingAmount()
//	daA := a.DemandingAmount(price)
//	daB := b.DemandingAmount(price)
//
//	var sos, bos Orders // smaller orders, bigger orders
//	// Remaining amount and demanding amount of smaller orders and bigger orders.
//	var sa, ba, sda sdk.Int
//	if amtA.LTE(daB) { // a is smaller than(or equal to) b
//		if daA.GT(amtB) { // sanity check TODO: remove
//			panic(fmt.Sprintf("%s > %s!", daA, amtB))
//		}
//		sos, bos = a, b
//		sa, ba = amtA, amtB
//		sda = daA
//	} else { // b is smaller than a
//		if daB.GT(amtA) { // sanity check TODO: remove
//			panic(fmt.Sprintf("%s > %s!", daB, amtA))
//		}
//		sos, bos = b, a
//		sa, ba = amtB, amtA
//		sda = daB
//	}
//
//	if sa.IsZero() || ba.IsZero() { // TODO: need more zero value checks?
//		return
//	}
//
//	for _, order := range sos {
//		proportion := order.RemainingAmount.ToDec().QuoTruncate(sa.ToDec()) // RemainingAmount / sa
//		order.RemainingAmount = sdk.ZeroInt()
//		var in sdk.Int
//		if sa.Equal(ba) {
//			in = ba.ToDec().MulTruncate(proportion).TruncateInt() // ba * proportion
//		} else {
//			in = sda.ToDec().MulTruncate(proportion).TruncateInt() // sda * proportion
//		}
//		order.ReceivedAmount = order.ReceivedAmount.Add(in)
//	}
//
//	for _, order := range bos {
//		proportion := order.RemainingAmount.ToDec().QuoTruncate(ba.ToDec()) // RemainingAmount / ba
//		if sa.Equal(ba) {
//			order.RemainingAmount = sdk.ZeroInt()
//		} else {
//			out := sda.ToDec().MulTruncate(proportion).TruncateInt() // sda * proportion
//			order.RemainingAmount = order.RemainingAmount.Sub(out)
//		}
//		in := sa.ToDec().MulTruncate(proportion).TruncateInt() // sa * proportion
//		order.ReceivedAmount = order.ReceivedAmount.Add(in)
//	}
//}

//// Order represents a swap order, which is made by a user or a pool.
//type Order struct {
//	Orderer         sdk.AccAddress
//	Direction       SwapDirection
//	Price           sdk.Dec
//	RemainingAmount sdk.Int
//	ReceivedAmount  sdk.Int
//}

type Order interface {
	Orderer() sdk.AccAddress
	Direction() SwapDirection
	Price() sdk.Dec
	RemainingAmount() sdk.Int
	SetRemainingAmount(amount sdk.Int)
	ReceivedAmount() sdk.Int
	SetReceivedAmount(amount sdk.Int)
}
