package types

import (
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OrderBook []OrderGroup

func (ob OrderBook) Add(order Order) OrderBook {
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
				newOrderBook[i].XToYOrders = append(newOrderBook[i].XToYOrders, order)
			case SwapDirectionYToX:
				newOrderBook[i].YToXOrders = append(newOrderBook[i].YToXOrders, order)
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

func (ob OrderBook) String() string {
	lines := []string{
		"+-----buy------+----------price-----------+-----sell-----+",
	}
	for _, og := range ob {
		xToYAmt, yToXAmt := sdk.ZeroInt(), sdk.ZeroInt()
		for _, order := range og.XToYOrders {
			xToYAmt = xToYAmt.Add(order.Amount)
		}
		for _, order := range og.YToXOrders {
			yToXAmt = yToXAmt.Add(order.Amount)
		}
		lines = append(lines, fmt.Sprintf("| %12s | %24s | %-12s |", xToYAmt, og.Price.String(), yToXAmt))
	}
	lines = append(lines, "+--------------+--------------------------+--------------+")
	return strings.Join(lines, "\n")
}

type OrderGroup struct {
	Price      sdk.Dec
	XToYOrders []Order
	YToXOrders []Order
}

func NewOrderGroup(order Order) OrderGroup {
	g := OrderGroup{Price: order.Price}
	switch order.Direction {
	case SwapDirectionXToY:
		g.XToYOrders = append(g.XToYOrders, order)
	case SwapDirectionYToX:
		g.YToXOrders = append(g.YToXOrders, order)
	}
	return g
}

// Order represents a swap order, which is made by a user or a pool.
// TODO: use SwapRequest instead - all fields are identical?
type Order struct {
	Orderer   string
	Direction SwapDirection
	Price     sdk.Dec
	Amount    sdk.Int
}

func NewOrder(orderer string, dir SwapDirection, price sdk.Dec, amt sdk.Int) Order {
	return Order{orderer, dir, price, amt}
}
