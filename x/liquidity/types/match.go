package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
