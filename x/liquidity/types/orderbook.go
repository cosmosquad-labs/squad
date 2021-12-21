package types

import (
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OrderBook []OrderBookItem

func (ob OrderBook) Add(os ...Order) OrderBook {
	result := ob
	for _, order := range os {
		result = result.add(order)
	}
	return result
}

func (ob OrderBook) add(order Order) OrderBook {
	var prices []string
	for _, og := range ob {
		prices = append(prices, og.Price.String())
	}
	i := sort.Search(len(prices), func(i int) bool {
		return prices[i] <= order.Price.String()
	})
	newOrderBook := ob
	if i < len(prices) {
		if prices[i] == order.Price.String() {
			switch order.Direction {
			case SwapDirectionXToY:
				newOrderBook[i].XToYOrders = append(newOrderBook[i].XToYOrders, &order)
			case SwapDirectionYToX:
				newOrderBook[i].YToXOrders = append(newOrderBook[i].YToXOrders, &order)
			}
		} else {
			// Insert a new order group at index i.
			newOrderBook = append(newOrderBook[:i], append([]OrderBookItem{NewOrderBookItem(order)}, newOrderBook[i:]...)...)
		}
	} else {
		// Append a new order group at the end.
		newOrderBook = append(newOrderBook, NewOrderBookItem(order))
	}
	return newOrderBook
}

func (ob OrderBook) HighestPriceXToYOrderGroupIndex(start int) (idx int, found bool) {
	for i := start; i < len(ob); i++ {
		if len(ob[i].XToYOrders) > 0 {
			idx = i
			found = true
			return
		}
	}
	return
}

func (ob OrderBook) LowestPriceYToXOrderGroupIndex(start int) (idx int, found bool) {
	for i := start; i >= 0; i-- {
		if len(ob[i].YToXOrders) > 0 {
			idx = i
			found = true
			return
		}
	}
	return
}

func (ob OrderBook) MatchAtPrice(price sdk.Dec) {
	bi, found := ob.HighestPriceXToYOrderGroupIndex(0)
	if !found || ob[bi].Price.LT(price) { // no matchable buy orders
		return
	}

	si, found := ob.LowestPriceYToXOrderGroupIndex(len(ob) - 1)
	if !found || ob[si].Price.GT(price) { // no matchable sell orders
		return
	}

	for {
		lbg := false // is this the last matchable buy order group?
		nbi, found := ob.HighestPriceXToYOrderGroupIndex(bi + 1)
		if !found || ob[nbi].Price.LT(price) { // no next matchable buy orders
			lbg = true
		}

		lsg := false // is this the last matchable sell order group?
		nsi, found := ob.LowestPriceYToXOrderGroupIndex(si - 1)
		if !found || ob[nsi].Price.GT(price) { // no next matchable sell orders
			lsg = true
		}

		bg := ob[bi] // current buy order group
		sg := ob[si] // current sell order group

		MatchOrders(bg.XToYOrders, sg.YToXOrders, price)

		if lbg || lsg {
			break
		}
		bi = nbi
		si = nsi
	}
}

func (ob OrderBook) String() string {
	lines := []string{
		"+-----buy------+----------price-----------+-----sell-----+",
	}
	for _, og := range ob {
		lines = append(lines,
			fmt.Sprintf("| %12s | %24s | %-12s |",
				og.RemainingXToYAmount(), og.Price.String(), og.RemainingYToXAmount()))
	}
	lines = append(lines, "+--------------+--------------------------+--------------+")
	return strings.Join(lines, "\n")
}

type OrderBookItem struct {
	Price      sdk.Dec
	XToYOrders Orders
	YToXOrders Orders
}

func NewOrderBookItem(order Order) OrderBookItem {
	g := OrderBookItem{Price: order.Price}
	switch order.Direction {
	case SwapDirectionXToY:
		g.XToYOrders = append(g.XToYOrders, &order)
	case SwapDirectionYToX:
		g.YToXOrders = append(g.YToXOrders, &order)
	}
	return g
}

func (og OrderBookItem) RemainingXToYAmount() sdk.Int {
	amt := sdk.ZeroInt()
	for _, order := range og.XToYOrders {
		amt = amt.Add(order.RemainingAmount)
	}
	return amt
}

func (og OrderBookItem) RemainingYToXAmount() sdk.Int {
	amt := sdk.ZeroInt()
	for _, order := range og.YToXOrders {
		amt = amt.Add(order.RemainingAmount)
	}
	return amt
}

type Orders []*Order

// RemainingAmount returns total remaining amount of orders.
// Note that orders should have same SwapDirection, since
// RemainingAmount doesn't rely on SwapDirection.
func (os Orders) RemainingAmount() sdk.Int {
	amt := sdk.ZeroInt()
	for _, order := range os {
		amt = amt.Add(order.RemainingAmount)
	}
	return amt
}

// DemandingAmount returns total demanding amount of orders at given price.
// Demanding amount is the amount of coins these orders want to receive.
// Note that orders should have same SwapDirection, since
// DemandingAmount doesn't rely on SwapDirection.
// TODO: use sdk.Dec here?
func (os Orders) DemandingAmount(price sdk.Dec) sdk.Int {
	da := sdk.ZeroInt()
	for _, order := range os {
		switch order.Direction {
		case SwapDirectionXToY:
			da = da.Add(order.RemainingAmount.ToDec().QuoTruncate(price).TruncateInt())
		case SwapDirectionYToX:
			da = da.Add(order.RemainingAmount.ToDec().MulTruncate(price).TruncateInt())
		}
	}
	return da
}

// MatchAll matches orders against other orders at given price.
// It consumes all remaining amount in the source orders(os).
func (os Orders) MatchAll(others Orders, price sdk.Dec) {
	amt := os.RemainingAmount()
	da := os.DemandingAmount(price)
	oa := others.RemainingAmount()

	for _, order := range os {
		proportion := order.RemainingAmount.ToDec().QuoTruncate(amt.ToDec())
		order.RemainingAmount = sdk.ZeroInt()
		in := da.ToDec().MulTruncate(proportion).TruncateInt()
		order.ReceivedAmount = order.ReceivedAmount.Add(in)
	}

	for _, order := range others {
		proportion := order.RemainingAmount.ToDec().QuoTruncate(oa.ToDec())
		out := da.ToDec().MulTruncate(proportion).TruncateInt()
		in := amt.ToDec().MulTruncate(proportion).TruncateInt()
		order.RemainingAmount = order.RemainingAmount.Sub(out)
		order.ReceivedAmount = order.ReceivedAmount.Add(in)
	}
}

// MatchOrders matches two order groups at given price.
func MatchOrders(a, b Orders, price sdk.Dec) {
	amtA := a.RemainingAmount()
	amtB := b.RemainingAmount()
	daA := a.DemandingAmount(price)
	daB := b.DemandingAmount(price)

	// determine which orders are bigger
	if amtA.LT(daB) {
		if daA.GT(amtB) { // sanity check TODO: remove
			panic(fmt.Sprintf("%s > %s!", daA, amtB))
		}
		a.MatchAll(b, price)
	} else {
		if daB.GT(amtA) { // sanity check TODO: remove
			panic(fmt.Sprintf("%s > %s!", daB, amtA))
		}
		b.MatchAll(a, price)
	}
}

// Order represents a swap order, which is made by a user or a pool.
type Order struct {
	Orderer         sdk.AccAddress
	Direction       SwapDirection
	Price           sdk.Dec
	RemainingAmount sdk.Int
	ReceivedAmount  sdk.Int
}
