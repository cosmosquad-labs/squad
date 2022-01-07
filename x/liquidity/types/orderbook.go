package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PriceDirection int

const (
	PriceIncreasing PriceDirection = iota + 1
	PriceDecreasing
)

//type OrderBook []OrderBookItem
//
//func (ob OrderBook) Add(os ...Order) OrderBook {
//	result := ob
//	for _, order := range os {
//		result = result.add(order)
//	}
//	return result
//}
//
//func (ob OrderBook) add(order Order) OrderBook {
//	var prices []string
//	for _, og := range ob {
//		prices = append(prices, og.Price.String())
//	}
//	i := sort.Search(len(prices), func(i int) bool {
//		return prices[i] <= order.Price.String()
//	})
//	newOrderBook := ob
//	if i < len(prices) {
//		if prices[i] == order.Price.String() {
//			switch order.Direction {
//			case SwapDirectionXToY:
//				newOrderBook[i].XToYOrders = append(newOrderBook[i].XToYOrders, &order)
//			case SwapDirectionYToX:
//				newOrderBook[i].YToXOrders = append(newOrderBook[i].YToXOrders, &order)
//			}
//		} else {
//			// Insert a new order group at index i.
//			newOrderBook = append(newOrderBook[:i], append([]OrderBookItem{NewOrderBookItem(order)}, newOrderBook[i:]...)...)
//		}
//	} else {
//		// Append a new order group at the end.
//		newOrderBook = append(newOrderBook, NewOrderBookItem(order))
//	}
//	return newOrderBook
//}
//
//func (ob OrderBook) HighestPriceXToYItemIndex(start int) (idx int, found bool) {
//	for i := start; i < len(ob); i++ {
//		if len(ob[i].XToYOrders) > 0 {
//			idx = i
//			found = true
//			return
//		}
//	}
//	return
//}
//
//func (ob OrderBook) LowestPriceYToXItemIndex(start int) (idx int, found bool) {
//	for i := start; i >= 0; i-- {
//		if len(ob[i].YToXOrders) > 0 {
//			idx = i
//			found = true
//			return
//		}
//	}
//	return
//}
//
//func (ob OrderBook) PriceDirection(lastPrice sdk.Dec) PriceDirection {
//	mb := sdk.ZeroInt()  // buy order amount with price higher than the last price
//	ms := sdk.ZeroInt()  // sell order amount with price lower than the last price
//	lpb := sdk.ZeroInt() // buy order amount with price equal to the last price
//	lps := sdk.ZeroInt() // buy order amount with price equal to the last price
//
//	for _, og := range ob {
//		switch {
//		case og.Price.GT(lastPrice):
//			mb = mb.Add(og.XToYOrders.RemainingAmount())
//		case og.Price.LT(lastPrice):
//			ms = ms.Add(og.YToXOrders.RemainingAmount())
//		default:
//			lpb = lpb.Add(og.XToYOrders.RemainingAmount())
//			lps = lps.Add(og.YToXOrders.RemainingAmount())
//		}
//	}
//
//	switch {
//	case mb.ToDec().GT(ms.Add(lps).ToDec().Mul(lastPrice)):
//		return PriceIncreasing
//	case ms.ToDec().Mul(lastPrice).GT(mb.Add(lpb).ToDec()):
//		return PriceDecreasing
//	default:
//		return PriceStaying
//	}
//}
//
//// lastPrice is the last order book price.
//func (ob OrderBook) Match(lastPrice sdk.Dec) {
//	pd := ob.PriceDirection(lastPrice) // price direction
//
//	bi, found := ob.HighestPriceXToYItemIndex(0)
//	if !found {
//		return
//	}
//
//	si, found := ob.LowestPriceYToXItemIndex(len(ob) - 1)
//	if !found {
//		return
//	}
//
//	for {
//		bg := ob[bi] // current buy order group
//		sg := ob[si] // current sell order group
//
//		if bg.Price.LT(sg.Price) {
//			break
//		}
//
//		var matchPrice sdk.Dec
//		switch pd {
//		case PriceIncreasing:
//			matchPrice = sg.Price
//		case PriceDecreasing:
//			matchPrice = bg.Price
//		case PriceStaying:
//			matchPrice = lastPrice
//		}
//		MatchOrders(bg.XToYOrders, sg.YToXOrders, matchPrice)
//
//		if bg.XToYOrders.RemainingAmount().IsZero() {
//			nbi, found := ob.HighestPriceXToYItemIndex(bi + 1)
//			if !found {
//				break
//			}
//			bi = nbi
//		}
//
//		if bg.YToXOrders.RemainingAmount().IsZero() {
//			nsi, found := ob.LowestPriceYToXItemIndex(si - 1)
//			if !found {
//				break
//			}
//			si = nsi
//		}
//	}
//}
//
//func (ob OrderBook) String() string {
//	lines := []string{
//		"+-----buy------+----------price-----------+-----sell-----+",
//	}
//	for _, og := range ob {
//		lines = append(lines,
//			fmt.Sprintf("| %12s | %24s | %-12s |",
//				og.XToYOrders.RemainingAmount(), og.Price.String(), og.YToXOrders.RemainingAmount()))
//	}
//	lines = append(lines, "+--------------+--------------------------+--------------+")
//	return strings.Join(lines, "\n")
//}
//
//type OrderBookItem struct {
//	Price      sdk.Dec
//	XToYOrders Orders
//	YToXOrders Orders
//}
//
//func NewOrderBookItem(order Order) OrderBookItem {
//	g := OrderBookItem{Price: order.Price}
//	switch order.Direction {
//	case SwapDirectionXToY:
//		g.XToYOrders = append(g.XToYOrders, &order)
//	case SwapDirectionYToX:
//		g.YToXOrders = append(g.YToXOrders, &order)
//	}
//	return g
//}
//
//type Orders []*Order
//
//// RemainingAmount returns total remaining amount of orders.
//// Note that orders should have same SwapDirection, since
//// RemainingAmount doesn't rely on SwapDirection.
//func (os Orders) RemainingAmount() sdk.Int {
//	amt := sdk.ZeroInt()
//	for _, order := range os {
//		amt = amt.Add(order.RemainingAmount)
//	}
//	return amt
//}
//
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

// Order represents a swap order, which is made by a user or a pool.
type Order struct {
	Orderer         sdk.AccAddress
	Direction       SwapDirection
	Price           sdk.Dec
	RemainingAmount sdk.Int
	ReceivedAmount  sdk.Int
}
