package amm

import (
	"fmt"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// FindMatchPrice returns the best match price for given order sources.
// If there is no matchable orders, found will be false.
func FindMatchPrice(ov OrderView, tickPrec int) (matchPrice sdk.Dec, found bool) {
	highestBuyPrice, found := ov.HighestBuyPrice()
	if !found {
		return sdk.Dec{}, false
	}
	lowestSellPrice, found := ov.LowestSellPrice()
	if !found {
		return sdk.Dec{}, false
	}
	if highestBuyPrice.LT(lowestSellPrice) {
		return sdk.Dec{}, false
	}

	prec := TickPrecision(tickPrec)
	lowestTickIdx := prec.TickToIndex(prec.LowestTick())
	highestTickIdx := prec.TickToIndex(prec.HighestTick())
	var i, j int
	i, found = findFirstTrueCondition(lowestTickIdx, highestTickIdx, func(i int) bool {
		return ov.BuyAmountOver(prec.TickFromIndex(i + 1)).LTE(ov.SellAmountUnder(prec.TickFromIndex(i)))
	})
	if !found {
		return sdk.Dec{}, false
	}
	j, found = findFirstTrueCondition(highestTickIdx, lowestTickIdx, func(i int) bool {
		return ov.BuyAmountOver(prec.TickFromIndex(i)).GTE(ov.SellAmountUnder(prec.TickFromIndex(i - 1)))
	})
	if !found {
		return sdk.Dec{}, false
	}
	midTick := TickFromIndex(i, tickPrec).Add(TickFromIndex(j, tickPrec)).QuoInt64(2)
	return RoundPrice(midTick, tickPrec), true
}

// findFirstTrueCondition uses the binary search to find the first index
// where f(i) is true, while searching in range [start, end].
// It assumes that f(j) == false where j < i and f(j) == true where j >= i.
// start can be greater than end.
func findFirstTrueCondition(start, end int, f func(i int) bool) (i int, found bool) {
	if start < end {
		i = start + sort.Search(end-start+1, func(i int) bool {
			return f(start + i)
		})
		if i > end {
			return 0, false
		}
		return i, true
	}
	i = start - sort.Search(start-end+1, func(i int) bool {
		return f(start - i)
	})
	if i < end {
		return 0, false
	}
	return i, true
}

// MatchOrders matches orders at given matchPrice if matchable.
// Note that MatchOrders modifies the orders in-place.
// Orders should be sorted appropriately.
// quoteCoinDust is the difference between total paid quote coin and total
// received quote coin.
// quoteCoinDust can be positive because of the decimal truncation.
func MatchOrders(buyOrders, sellOrders []Order, matchPrice sdk.Dec) (quoteCoinDust sdk.Int, matched bool) {
	buyOrders = DropSmallOrders(buyOrders, matchPrice)
	sellOrders = DropSmallOrders(sellOrders, matchPrice)
	if len(buyOrders) == 0 || len(sellOrders) == 0 {
		return sdk.Int{}, false
	}

	totalBuyAmt := TotalAmount(buyOrders)
	totalSellAmt := TotalAmount(sellOrders)

	quoteCoinDust = sdk.ZeroInt()
	if totalBuyAmt.Equal(totalSellAmt) {
		for _, order := range append(buyOrders, sellOrders...) {
			matchOrder(order, order.GetAmount())
		}
	} else {
		var (
			smallOrdersAmt           sdk.Int
			smallOrders, largeOrders []Order
		)
		if totalBuyAmt.LT(totalSellAmt) {
			smallOrdersAmt = totalBuyAmt
			smallOrders, largeOrders = buyOrders, sellOrders
		} else {
			smallOrdersAmt = totalSellAmt
			smallOrders, largeOrders = sellOrders, buyOrders
		}

		bestLargeOrders, restLargeOrders := SplitOrders(largeOrders)

		for _, order := range append(smallOrders, restLargeOrders...) {
			matchOrder(order, order.GetAmount())
		}

		restLargeOrdersAmt := TotalAmount(restLargeOrders)
		remainingSmallOrdersAmt := smallOrdersAmt.Sub(restLargeOrdersAmt)
		if remainingSmallOrdersAmt.IsPositive() {
			DistributeOrderAmount(bestLargeOrders, matchPrice, remainingSmallOrdersAmt)
		}
	}

	for _, order := range append(buyOrders, sellOrders...) {
		quoteCoinDust = quoteCoinDust.Add(updateMatchedOrder(order, matchPrice))
	}

	return quoteCoinDust, true
}

// DistributeOrderAmount distributes the given order amount to the orders
// proportional to each order's amount.
// After distributing the amount based on each order's proportion,
// remaining amount due to the decimal truncation is distributed
// to the orders again, by priority.
// This time, the proportion is not considered and each order takes up
// the amount as much as possible.
func DistributeOrderAmount(orders []Order, matchPrice sdk.Dec, amt sdk.Int) {
	totalAmt := TotalAmount(orders)
	totalMatchedAmt := sdk.ZeroInt()

	for _, order := range orders {
		orderAmt := order.GetAmount().ToDec()
		proportion := orderAmt.QuoTruncate(totalAmt.ToDec())
		matchedAmt := proportion.MulInt(amt).TruncateInt()
		matchOrder(order, matchedAmt) // temporarily increment matched amount
		if matchedAmt.IsPositive() {
			totalMatchedAmt = totalMatchedAmt.Add(matchedAmt)
		}
	}

	remainingAmt := amt.Sub(totalMatchedAmt)
	if remainingAmt.IsPositive() {
		for _, order := range orders {
			matchedAmt := sdk.MinInt(remainingAmt, order.GetOpenAmount())
			matchOrder(order, matchedAmt) // temporarily increment matched amount
			remainingAmt = remainingAmt.Sub(matchedAmt)
		}
	}

	var matchedOrders, notMatchedOrders []Order
	for _, order := range orders {
		matchedAmt := order.GetAmount().Sub(order.GetOpenAmount())
		if !matchedAmt.IsZero() && (order.GetDirection() == Buy || matchPrice.MulInt(matchedAmt).IsPositive()) {
			matchedOrders = append(matchedOrders, order)
		} else {
			notMatchedOrders = append(notMatchedOrders, order)
		}
	}

	if len(notMatchedOrders) > 0 {
		// Reset matched amount
		for _, order := range orders {
			order.SetOpenAmount(order.GetAmount())
		}
		DistributeOrderAmount(matchedOrders, matchPrice, amt)
	}
}

// SplitOrders splits orders by price, by returning two order slices,
// (orders with the best price, rest orders).
// The best price is the lowest order price for buy orders and
// the highest price for sell orders.
// Note that the orders are must be sorted by price.
func SplitOrders(orders []Order) (bestOrders, restOrders []Order) {
	bestPrice := orders[len(orders)-1].GetPrice()
	i := sort.Search(len(orders), func(i int) bool {
		return orders[i].GetPrice().Equal(bestPrice)
	})
	return orders[i:], orders[:i]
}

// matchOrder increments matched amount of an order.
// Later, updateMatchedOrder should be called for paying/receiving quote coins.
func matchOrder(order Order, amt sdk.Int) {
	order.DecrOpenAmount(amt)
}

// updateMatchedOrder finalizes the matched order by updating paying/receiving
// quote coin amount and marking the order as matched.
// quoteCoinDiff is how many quote coins are paid and received during
// the matching and can be zero or a positive number.
func updateMatchedOrder(order Order, matchPrice sdk.Dec) (quoteCoinDiff sdk.Int) {
	matchedAmt := order.GetAmount().Sub(order.GetOpenAmount())
	switch order.GetDirection() {
	case Buy:
		paidQuoteCoinAmt := matchPrice.MulInt(matchedAmt).Ceil().TruncateInt()
		order.DecrRemainingOfferCoin(paidQuoteCoinAmt)
		order.IncrReceivedDemandCoin(matchedAmt)
		order.SetMatched(true)
		return paidQuoteCoinAmt
	case Sell:
		receivedQuoteCoinAmt := matchPrice.MulInt(matchedAmt).TruncateInt()
		order.DecrRemainingOfferCoin(matchedAmt)
		order.IncrReceivedDemandCoin(receivedQuoteCoinAmt)
		order.SetMatched(true)
		return receivedQuoteCoinAmt.Neg()
	default:
		panic(fmt.Errorf("invalid order direction: %s", order.GetDirection()))
	}
}

// DropSmallOrders returns filtered orders, where orders with too small amount
// are dropped.
func DropSmallOrders(orders []Order, matchPrice sdk.Dec) []Order {
	var res []Order
	for _, order := range orders {
		amt := order.GetAmount()
		// TODO: drop only when receiving coin is 0?
		if amt.GTE(MinCoinAmount) && matchPrice.MulInt(amt).TruncateInt().GTE(MinCoinAmount) {
			res = append(res, order)
		}
	}
	return res
}
