package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MatchEngine struct {
	BuyOrderSource  OrderSource
	SellOrderSource OrderSource
	TickPrecision   int // price tick precision
}

func NewMatchEngine(buys, sells OrderSource, prec int) *MatchEngine {
	return &MatchEngine{
		BuyOrderSource:  buys,
		SellOrderSource: sells,
		TickPrecision:   prec,
	}
}

func (engine *MatchEngine) Matchable() bool {
	highestBuyPrice, found := engine.BuyOrderSource.HighestTick(engine.TickPrecision)
	if !found {
		return false
	}
	return engine.SellOrderSource.AmountLTE(highestBuyPrice).IsPositive()
}

func (engine *MatchEngine) EstimatedPriceDirection(lastPrice sdk.Dec) PriceDirection {
	buyAmount := engine.BuyOrderSource.AmountGTE(lastPrice)
	sellAmount := engine.SellOrderSource.AmountLTE(lastPrice)
	if buyAmount.ToDec().GTE(lastPrice.MulInt(sellAmount)) {
		return PriceIncreasing
	}
	return PriceDecreasing
}

// SwapPrice assumes that the last price is fit in tick.
func (engine *MatchEngine) SwapPrice(lastPrice sdk.Dec) sdk.Dec {
	dir := engine.EstimatedPriceDirection(lastPrice)
	tickSource := MergeOrderSources(engine.BuyOrderSource, engine.SellOrderSource) // temporary order source just for ticks

	buysCache := map[int]sdk.Int{}
	buyAmountGTE := func(i int) sdk.Int {
		ba, ok := buysCache[i]
		if !ok {
			ba = engine.BuyOrderSource.AmountGTE(TickFromIndex(i, engine.TickPrecision))
			buysCache[i] = ba
		}
		return ba
	}
	sellsCache := map[int]sdk.Int{}
	sellAmountLTE := func(i int) sdk.Int {
		sa, ok := sellsCache[i]
		if !ok {
			sa = engine.SellOrderSource.AmountLTE(TickFromIndex(i, engine.TickPrecision))
			sellsCache[i] = sa
		}
		return sa
	}

	currentPrice := lastPrice
	for {
		i := TickToIndex(currentPrice, engine.TickPrecision)
		ba := buyAmountGTE(i)
		sa := sellAmountLTE(i)
		hba := buyAmountGTE(i + 1)
		lsa := sellAmountLTE(i - 1)

		if currentPrice.MulInt(sa).TruncateInt().GTE(hba) && ba.GTE(currentPrice.MulInt(lsa).TruncateInt()) {
			return currentPrice
		}

		if dir == PriceIncreasing && hba.IsZero() || dir == PriceDecreasing && lsa.IsZero() {
			return currentPrice
		}

		var nextPrice sdk.Dec
		var found bool
		switch dir {
		case PriceIncreasing:
			nextPrice, found = tickSource.UpTick(currentPrice, engine.TickPrecision)
		case PriceDecreasing:
			nextPrice, found = tickSource.DownTick(currentPrice, engine.TickPrecision)
		}
		if !found {
			return currentPrice
		}
		currentPrice = nextPrice
	}
}

func (engine *MatchEngine) Match(lastPrice sdk.Dec) (orderBook *OrderBook, swapPrice sdk.Dec, matched bool) {
	if !engine.Matchable() {
		return
	}
	matched = true

	swapPrice = engine.SwapPrice(lastPrice)
	buyPrice, _ := engine.BuyOrderSource.HighestTick(engine.TickPrecision)
	sellPrice, _ := engine.SellOrderSource.LowestTick(engine.TickPrecision)

	orderBook = NewOrderBook()

	for {
		buyOrders := orderBook.BuyTicks.Orders(buyPrice)
		if len(buyOrders) == 0 {
			orderBook.AddOrders(engine.BuyOrderSource.Orders(buyPrice)...)
			buyOrders = orderBook.BuyTicks.Orders(buyPrice)
		}
		sellOrders := orderBook.SellTicks.Orders(sellPrice)
		if len(sellOrders) == 0 {
			orderBook.AddOrders(engine.SellOrderSource.Orders(sellPrice)...)
			sellOrders = orderBook.SellTicks.Orders(sellPrice)
		}

		MatchOrders(buyOrders, sellOrders, swapPrice)

		if (buyPrice.Equal(swapPrice) && sellPrice.Equal(swapPrice)) ||
			buyPrice.LT(swapPrice) || sellPrice.GT(swapPrice) {
			break
		}

		if buyOrders.RemainingAmount().IsZero() {
			var found bool
			buyPrice, found = engine.BuyOrderSource.DownTick(buyPrice, engine.TickPrecision)
			if !found {
				break
			}
		}
		if sellOrders.RemainingAmount().IsZero() {
			var found bool
			sellPrice, found = engine.SellOrderSource.UpTick(sellPrice, engine.TickPrecision)
			if !found {
				break
			}
		}
	}

	return
}

// MatchOrders matches two order groups at given price.
func MatchOrders(buyOrders, sellOrders Orders, price sdk.Dec) {
	buyAmount := buyOrders.RemainingAmount()
	sellAmount := sellOrders.RemainingAmount()

	if buyAmount.IsZero() || sellAmount.IsZero() {
		return
	}

	matchAll := false
	sellDemandAmount := price.MulInt(sellAmount).TruncateInt()
	if buyAmount.Equal(sellDemandAmount) {
		matchAll = true
	}
	buyDemandAmount := buyAmount.ToDec().QuoTruncate(price).TruncateInt()

	var smallerOrders, biggerOrders Orders
	var smallerAmount, biggerAmount sdk.Int
	var smallerDemandAmount sdk.Int
	if buyAmount.LTE(sellDemandAmount) { // Note that we use LTE here.
		smallerOrders, biggerOrders = buyOrders, sellOrders
		smallerAmount, biggerAmount = buyAmount, sellAmount
		smallerDemandAmount = buyDemandAmount
	} else {
		smallerOrders, biggerOrders = sellOrders, buyOrders
		smallerAmount, biggerAmount = sellAmount, buyAmount
		smallerDemandAmount = sellDemandAmount
	}

	for _, order := range smallerOrders {
		proportion := order.RemainingAmount.ToDec().QuoInt(smallerAmount)
		order.RemainingAmount = sdk.ZeroInt()
		in := proportion.MulInt(smallerDemandAmount).TruncateInt()
		order.ReceivedAmount = order.ReceivedAmount.Add(in)
	}

	for _, order := range biggerOrders {
		proportion := order.RemainingAmount.ToDec().QuoInt(biggerAmount)
		if matchAll {
			order.RemainingAmount = sdk.ZeroInt()
		} else {
			out := proportion.MulInt(smallerDemandAmount).TruncateInt()
			order.RemainingAmount = order.RemainingAmount.Sub(out)
		}
		in := proportion.MulInt(smallerAmount).TruncateInt()
		order.ReceivedAmount = order.ReceivedAmount.Add(in)
	}
}
