package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ OrderSource = (*mergedOrderSource)(nil)

type OrderSource interface {
	AmountGTE(price sdk.Dec) sdk.Int
	AmountLTE(price sdk.Dec) sdk.Int
	Orders(price sdk.Dec) []Order
	UpTick(price sdk.Dec, prec int) (tick sdk.Dec, found bool)
	DownTick(price sdk.Dec, prec int) (tick sdk.Dec, found bool)
	//HighestTick(prec int) (res sdk.Dec, found bool)
	//LowestTick(prec int) (res sdk.Dec, found bool)
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

func (mos *mergedOrderSource) Orders(price sdk.Dec) []Order {
	var os []Order
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

func MergeOrderSources(sources ...OrderSource) OrderSource {
	return &mergedOrderSource{sources: sources}
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

func (eng *MatchEngine) EstimatedPriceDirection(lastPrice sdk.Dec) PriceDirection {
	if eng.buys.AmountGTE(lastPrice).ToDec().GTE(lastPrice.MulInt(eng.sells.AmountLTE(lastPrice))) {
		return PriceIncreasing
	}
	return PriceDecreasing
}

func (eng *MatchEngine) SwapPrice(lastPrice sdk.Dec) sdk.Dec {
	dir := eng.EstimatedPriceDirection(lastPrice)

	os := MergeOrderSources(eng.buys, eng.sells) // temporary order source just for ticks
	curPrice := PriceToTick(lastPrice, eng.prec)
	if dir == PriceIncreasing {
		if curPrice.LT(lastPrice) {
			curPrice = UpTick(curPrice, eng.prec)
		}
	}
	lowestPrice := LowestTick(eng.prec)

	swapPrice := curPrice
	for {
		ba := eng.buys.AmountGTE(curPrice)
		sa := curPrice.MulInt(eng.sells.AmountLTE(curPrice)).TruncateInt()

		var next sdk.Dec
		var found bool
		switch dir {
		case PriceIncreasing:
			if sa.GT(ba) {
				return swapPrice
			}
			// TODO: check if there is no uptick?
			next, found = os.UpTick(curPrice, eng.prec)
		case PriceDecreasing:
			if ba.GT(sa) {
				return swapPrice
			}

			if curPrice.Equal(lowestPrice) {
				return curPrice
			}

			next, found = os.DownTick(curPrice, eng.prec)
		}
		if !found {
			return curPrice
		}
		swapPrice = curPrice
		curPrice = next
	}
}

func (eng *MatchEngine) Match(lastPrice sdk.Dec) {

}
