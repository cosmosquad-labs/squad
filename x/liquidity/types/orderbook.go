package types

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OrderBook []OrderGroup

func (orderBook OrderBook) Add(order Order) OrderBook {
	var prices []string
	for _, g := range orderBook {
		prices = append(prices, g.Price.String())
	}
	i := sort.Search(len(prices), func(i int) bool {
		return prices[i] <= order.Price.String()
	})
	newOrderBook := orderBook
	if i < len(prices) {
		if prices[i] == order.Price.String() {
			switch order.Direction {
			case SwapDirectionXToY:
				newOrderBook[i].XToYOrders = append(newOrderBook[i].XToYOrders, order)
			case SwapDirectionYToX:
				newOrderBook[i].YToXOrders = append(newOrderBook[i].YToXOrders, order)
			}
		} else {
			newOrderBook = append(newOrderBook[:i], append([]OrderGroup{NewOrderGroup(order)}, newOrderBook[i:]...)...)
		}
	} else {
		newOrderBook = append(newOrderBook, NewOrderGroup(order))
	}
	return newOrderBook
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
