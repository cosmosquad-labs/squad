package types

import (
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OrderBook []OrderGroup

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
			newOrderBook = append(newOrderBook[:i], append([]OrderGroup{NewOrderGroup(order)}, newOrderBook[i:]...)...)
		}
	} else {
		// Append a new order group at the end.
		newOrderBook = append(newOrderBook, NewOrderGroup(order))
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

type OrderGroup struct {
	Price      sdk.Dec
	XToYOrders Orders
	YToXOrders Orders
}

func NewOrderGroup(order Order) OrderGroup {
	g := OrderGroup{Price: order.Price}
	switch order.Direction {
	case SwapDirectionXToY:
		g.XToYOrders = append(g.XToYOrders, &order)
	case SwapDirectionYToX:
		g.YToXOrders = append(g.YToXOrders, &order)
	}
	return g
}

func (og OrderGroup) RemainingXToYAmount() sdk.Int {
	amt := sdk.ZeroInt()
	for _, order := range og.XToYOrders {
		amt = amt.Add(order.RemainingAmount)
	}
	return amt
}

func (og OrderGroup) RemainingYToXAmount() sdk.Int {
	amt := sdk.ZeroInt()
	for _, order := range og.YToXOrders {
		amt = amt.Add(order.RemainingAmount)
	}
	return amt
}

type Orders []*Order

func (os Orders) Amount() sdk.Int {
	amt := sdk.ZeroInt()
	for _, order := range os {
		amt = amt.Add(order.RemainingAmount)
	}
	return amt
}

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

// Match matches orders against other orders at given price.
// It consumes all remaining amount in source(the receiver, os) orders.
func (os Orders) Match(others Orders, price sdk.Dec) {
	amt := os.Amount()
	da := os.DemandingAmount(price)
	oa := others.Amount()

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

func MatchOrders(buys, sells Orders, price sdk.Dec) {
	bx := buys.Amount()                 // buy orders total amount
	sy := sells.Amount()                // sell orders total amount
	bdy := buys.DemandingAmount(price)  // buy orders demanding amount
	sdx := sells.DemandingAmount(price) // sell orders demanding amount

	// determine which orders are bigger
	if bx.LT(sdx) { // sell orders are bigger than buy orders
		if bdy.GT(sy) { // sanity check TODO: remove
			panic(fmt.Sprintf("%s > %s!", bdy, sy))
		}
		buys.Match(sells, price)
	} else { // sell orders are bigger than(or equal to) buy orders
		if sdx.GT(bx) { // sanity check TODO: remove
			panic(fmt.Sprintf("%s > %s!", sdx, bx))
		}
		sells.Match(buys, price)
	}
}

// Order represents a swap order, which is made by a user or a pool.
// TODO: use SwapRequest instead - all fields are identical?
type Order struct {
	Orderer         sdk.AccAddress // where the ReceivedAmount coin goes
	Direction       SwapDirection
	Price           sdk.Dec
	RemainingAmount sdk.Int
	ReceivedAmount  sdk.Int
}
