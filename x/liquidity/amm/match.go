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

// FindLastMatchableOrders returns the last matchable order indexes for
// each buy/sell side.
// lastBuyPartialMatchAmt and lastSellPartialMatchAmt are
// the amount of partially matched portion of the last orders.
// FindLastMatchableOrders drops(ignores) an order if the orderer
// receives zero demand coin after truncation when the order is either
// fully matched or partially matched.
func FindLastMatchableOrders(buyOrders, sellOrders []Order, matchPrice sdk.Dec) (lastBuyIdx, lastSellIdx int, lastBuyPartialMatchAmt, lastSellPartialMatchAmt sdk.Int, found bool) {
	if len(buyOrders) == 0 || len(sellOrders) == 0 {
		return 0, 0, sdk.Int{}, sdk.Int{}, false
	}
	type Side struct {
		orders          []Order
		totalOpenAmt    sdk.Int
		i               int
		partialMatchAmt sdk.Int
	}
	buySide := &Side{buyOrders, TotalOpenAmount(buyOrders), len(buyOrders) - 1, sdk.Int{}}
	sellSide := &Side{sellOrders, TotalOpenAmount(sellOrders), len(sellOrders) - 1, sdk.Int{}}
	sides := map[OrderDirection]*Side{
		Buy:  buySide,
		Sell: sellSide,
	}
	// Repeatedly check both buy/sell side to see if there is an order to drop.
	// If there is not, then the loop is finished.
	for {
		ok := true
		for _, dir := range []OrderDirection{Buy, Sell} {
			side := sides[dir]
			i := side.i
			order := side.orders[i]
			matchAmt := sdk.MinInt(buySide.totalOpenAmt, sellSide.totalOpenAmt)
			otherOrdersAmt := side.totalOpenAmt.Sub(order.GetOpenAmount())
			// side.partialMatchAmt can be negative at this moment, but
			// FindLastMatchableOrders won't return a negative amount because
			// the if-block below would set ok = false if otherOrdersAmt >= matchAmt
			// and the loop would be continued.
			side.partialMatchAmt = matchAmt.Sub(otherOrdersAmt)
			if otherOrdersAmt.GTE(matchAmt) ||
				(dir == Sell && matchPrice.MulInt(side.partialMatchAmt).TruncateInt().IsZero()) {
				if i == 0 { // There's no orders left, which means orders are not matchable.
					return 0, 0, sdk.Int{}, sdk.Int{}, false
				}
				side.totalOpenAmt = side.totalOpenAmt.Sub(order.GetOpenAmount())
				side.i--
				ok = false
			}
		}
		if ok {
			return buySide.i, sellSide.i, buySide.partialMatchAmt, sellSide.partialMatchAmt, true
		}
	}
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

	totalBuyAmt := TotalOpenAmount(buyOrders)
	totalSellAmt := TotalOpenAmount(sellOrders)

	quoteCoinDust = sdk.ZeroInt()
	if totalBuyAmt.Equal(totalSellAmt) {
		for _, order := range append(buyOrders, sellOrders...) {
			MatchOrder(order, order.GetAmount())
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
			MatchOrder(order, order.GetAmount())
		}

		restLargeOrdersAmt := TotalOpenAmount(restLargeOrders)
		remainingSmallOrdersAmt := smallOrdersAmt.Sub(restLargeOrdersAmt)
		if remainingSmallOrdersAmt.IsPositive() {
			DistributeAmount(bestLargeOrders, matchPrice, remainingSmallOrdersAmt)
		}
	}

	for _, order := range append(buyOrders, sellOrders...) {
		quoteCoinDust = quoteCoinDust.Add(UpdateMatchedOrder(order, matchPrice))
	}

	return quoteCoinDust, true
}

func DistributeAmount(orders []Order, matchPrice sdk.Dec, amt sdk.Int) {
	totalAmt := TotalOpenAmount(orders)
	totalMatchedAmt := sdk.ZeroInt()

	for _, order := range orders {
		orderAmt := order.GetAmount().ToDec()
		proportion := orderAmt.QuoTruncate(totalAmt.ToDec())
		matchedAmt := proportion.MulInt(amt).TruncateInt()
		MatchOrder(order, matchedAmt) // temporarily increment matched amount
		if matchedAmt.IsPositive() {
			totalMatchedAmt = totalMatchedAmt.Add(matchedAmt)
		}
	}

	remainingAmt := amt.Sub(totalMatchedAmt)
	if remainingAmt.IsPositive() {
		for _, order := range orders {
			matchedAmt := sdk.MinInt(remainingAmt, order.GetOpenAmount())
			MatchOrder(order, matchedAmt) // temporarily increment matched amount
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
		DistributeAmount(matchedOrders, matchPrice, amt)
	}
}

// Note that the orders are must be sorted by price.
func SplitOrders(orders []Order) (rest, target []Order) {
	targetPrice := orders[len(orders)-1].GetPrice()
	i := sort.Search(len(orders), func(i int) bool {
		return orders[i].GetPrice().Equal(targetPrice)
	})
	return orders[i:], orders[:i]
}

func MatchOrder(order Order, amt sdk.Int) {
	order.DecrOpenAmount(amt)
}

func UpdateMatchedOrder(order Order, matchPrice sdk.Dec) (quoteCoinDiff sdk.Int) {
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
		openAmt := order.GetOpenAmount()
		if openAmt.GTE(MinCoinAmount) && matchPrice.MulInt(openAmt).TruncateInt().GTE(MinCoinAmount) {
			res = append(res, order)
		}
	}
	return res
}
