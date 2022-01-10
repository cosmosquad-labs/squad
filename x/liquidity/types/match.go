package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
	switch {
	case buyAmount.LT(sellDemandAmount):
		smallerOrders, biggerOrders = buyOrders, sellOrders
		smallerAmount, biggerAmount = buyAmount, sellAmount
		smallerDemandAmount = buyDemandAmount
	case buyAmount.GT(sellDemandAmount):
		smallerOrders, biggerOrders = sellOrders, buyOrders
		smallerAmount, biggerAmount = sellAmount, buyAmount
		smallerDemandAmount = sellDemandAmount
	}

	for _, order := range smallerOrders {
		proportion := order.RemainingAmount().ToDec().QuoInt(smallerAmount)
		if matchAll {
			order.SetRemainingAmount(sdk.ZeroInt())
		} else {
			out := proportion.MulInt(smallerAmount).TruncateInt()
			order.SetRemainingAmount(order.RemainingAmount().Sub(out))
		}
		in := proportion.MulInt(smallerDemandAmount).TruncateInt()
		order.SetReceivedAmount(order.ReceivedAmount().Add(in))
	}

	for _, order := range biggerOrders {
		proportion := order.RemainingAmount().ToDec().QuoInt(biggerAmount)
		if matchAll {
			order.SetRemainingAmount(sdk.ZeroInt())
		} else {
			out := proportion.MulInt(smallerDemandAmount).TruncateInt()
			order.SetRemainingAmount(order.RemainingAmount().Sub(out))
		}
		in := proportion.MulInt(smallerAmount).TruncateInt()
		order.SetReceivedAmount(order.ReceivedAmount().Add(in))
	}
}

type MatchEngine struct {
	buys  OrderSource
	sells OrderSource
	prec  int // price tick precision
}

func NewMatchEngine(buys, sells OrderSource, prec int) *MatchEngine {
	return &MatchEngine{
		buys:  buys,
		sells: sells,
		prec:  prec,
	}
}

func (eng *MatchEngine) Matchable() bool {
	highestBuyPrice, found := eng.buys.HighestTick(eng.prec)
	if !found {
		return false
	}
	return eng.sells.AmountLTE(highestBuyPrice).IsPositive()
}

func (eng *MatchEngine) EstimatedPriceDirection(lastPrice sdk.Dec) PriceDirection {
	if eng.buys.AmountGTE(lastPrice).ToDec().GTE(lastPrice.MulInt(eng.sells.AmountLTE(lastPrice))) {
		return PriceIncreasing
	}
	return PriceDecreasing
}

// SwapPrice assumes that the last price is fit in tick.
func (eng *MatchEngine) SwapPrice(lastPrice sdk.Dec) sdk.Dec {
	dir := eng.EstimatedPriceDirection(lastPrice)
	os := MergeOrderSources(eng.buys, eng.sells) // temporary order source just for ticks

	buysCache := map[int]sdk.Int{}
	buyAmountGTE := func(i int) sdk.Int {
		ba, ok := buysCache[i]
		if !ok {
			ba = eng.buys.AmountGTE(TickFromIndex(i, eng.prec))
			buysCache[i] = ba
		}
		return ba
	}
	sellsCache := map[int]sdk.Int{}
	sellAmountLTE := func(i int) sdk.Int {
		sa, ok := sellsCache[i]
		if !ok {
			sa = eng.sells.AmountLTE(TickFromIndex(i, eng.prec))
			sellsCache[i] = sa
		}
		return sa
	}

	currentPrice := lastPrice
	for {
		i := TickToIndex(currentPrice, eng.prec)
		ba := buyAmountGTE(i)
		sa := sellAmountLTE(i)
		hba := buyAmountGTE(i + 1)
		lsa := sellAmountLTE(i - 1)

		if currentPrice.MulInt(sa).TruncateInt().GTE(hba) && ba.GTE(currentPrice.MulInt(lsa).TruncateInt()) {
			return currentPrice
		}

		var nextPrice sdk.Dec
		var found bool
		switch dir {
		case PriceIncreasing:
			nextPrice, found = os.UpTick(currentPrice, eng.prec)
		case PriceDecreasing:
			nextPrice, found = os.DownTick(currentPrice, eng.prec)
		}
		if !found {
			return currentPrice
		}
		currentPrice = nextPrice
	}
}

func (eng *MatchEngine) Match(lastPrice sdk.Dec) {
	if !eng.Matchable() {
		return
	}
}
