package types

import (
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ OrderSource = (*OrderBookTicks)(nil)
	_ OrderSource = (*mergedOrderSource)(nil)
)

type PriceDirection int

const (
	PriceIncreasing PriceDirection = iota + 1
	PriceDecreasing
)

type Order struct {
	Orderer         sdk.AccAddress
	Direction       SwapDirection
	Price           sdk.Dec
	OrderAmount     sdk.Int
	RemainingAmount sdk.Int
	ReceivedAmount  sdk.Int
}

func NewOrder(orderer sdk.AccAddress, dir SwapDirection, price sdk.Dec, amount sdk.Int) *Order {
	return &Order{
		Orderer:         orderer,
		Direction:       dir,
		Price:           price,
		OrderAmount:     amount,
		RemainingAmount: amount,
		ReceivedAmount:  sdk.ZeroInt(),
	}
}

type Orders []*Order

func (orders Orders) RemainingAmount() sdk.Int {
	amount := sdk.ZeroInt()
	for _, order := range orders {
		amount = amount.Add(order.RemainingAmount)
	}
	return amount
}

type OrderBookTick struct {
	price  sdk.Dec
	orders Orders
}

func NewOrderBookTick(order *Order) *OrderBookTick {
	return &OrderBookTick{
		price:  order.Price,
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
		return prices[i].LTE(price)
	})
	if i < len(prices) && prices[i].Equal(price) {
		exact = true
	}
	return
}

func (ticks *OrderBookTicks) AddOrder(order *Order) {
	i, exact := ticks.FindPrice(order.Price)
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

func (ticks OrderBookTicks) Orders(price sdk.Dec) Orders {
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

func NewOrderBook() *OrderBook {
	return &OrderBook{
		buys:  OrderBookTicks{},
		sells: OrderBookTicks{},
	}
}

func (ob *OrderBook) AddOrder(order *Order) {
	switch order.Direction {
	case SwapDirectionBuy:
		ob.buys.AddOrder(order)
	case SwapDirectionSell:
		ob.sells.AddOrder(order)
	}
}

func (ob *OrderBook) AddOrders(orders ...*Order) {
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

func (ob OrderBook) String() string {
	os := MergeOrderSources(ob.buys, ob.sells)
	price, found := os.HighestTick(0)
	if !found {
		return "<nil>"
	}
	lines := []string{
		"+-----buy------+----------price-----------+-----sell-----+",
	}
	for {
		lines = append(lines,
			fmt.Sprintf("| %12s | %24s | %-12s |",
				ob.buys.Orders(price).RemainingAmount(), price.String(), ob.sells.Orders(price).RemainingAmount()))

		price, found = os.DownTick(price, 0)
		if !found {
			break
		}
	}
	lines = append(lines, "+--------------+--------------------------+--------------+")
	return strings.Join(lines, "\n")
}

// OrderSource defines a source of orders which can be an order book or
// a pool.
// TODO: omit prec parameter?
type OrderSource interface {
	AmountGTE(price sdk.Dec) sdk.Int
	AmountLTE(price sdk.Dec) sdk.Int
	Orders(price sdk.Dec) Orders
	UpTick(price sdk.Dec, prec int) (tick sdk.Dec, found bool)
	DownTick(price sdk.Dec, prec int) (tick sdk.Dec, found bool)
	HighestTick(prec int) (tick sdk.Dec, found bool)
	LowestTick(prec int) (tick sdk.Dec, found bool)
}

type mergedOrderSource struct {
	sources []OrderSource
}

func (mos *mergedOrderSource) AmountGTE(price sdk.Dec) sdk.Int {
	amt := sdk.ZeroInt()
	for _, source := range mos.sources {
		amt = amt.Add(source.AmountGTE(price))
	}
	return amt
}

func (mos *mergedOrderSource) AmountLTE(price sdk.Dec) sdk.Int {
	amt := sdk.ZeroInt()
	for _, source := range mos.sources {
		amt = amt.Add(source.AmountLTE(price))
	}
	return amt
}

func (mos *mergedOrderSource) Orders(price sdk.Dec) Orders {
	var os Orders
	for _, source := range mos.sources {
		os = append(os, source.Orders(price)...)
	}
	return os
}

func (mos *mergedOrderSource) UpTick(price sdk.Dec, prec int) (tick sdk.Dec, found bool) {
	for _, source := range mos.sources {
		t, f := source.UpTick(price, prec)
		if f && (tick.IsNil() || t.LT(tick)) {
			tick = t
			found = true
		}
	}
	return
}

func (mos *mergedOrderSource) DownTick(price sdk.Dec, prec int) (tick sdk.Dec, found bool) {
	for _, source := range mos.sources {
		t, f := source.DownTick(price, prec)
		if f && (tick.IsNil() || t.GT(tick)) {
			tick = t
			found = true
		}
	}
	return
}

func (mos *mergedOrderSource) HighestTick(prec int) (tick sdk.Dec, found bool) {
	for _, source := range mos.sources {
		t, f := source.HighestTick(prec)
		if f && (tick.IsNil() || t.GT(tick)) {
			tick = t
			found = true
		}
	}
	return
}

func (mos *mergedOrderSource) LowestTick(prec int) (tick sdk.Dec, found bool) {
	for _, source := range mos.sources {
		t, f := source.LowestTick(prec)
		if f && (tick.IsNil() || t.LT(tick)) {
			tick = t
			found = true
		}
	}
	return
}

func MergeOrderSources(sources ...OrderSource) OrderSource {
	return &mergedOrderSource{sources: sources}
}
