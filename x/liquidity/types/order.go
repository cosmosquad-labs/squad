package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OrderSource interface {
	AmountGTE(price sdk.Dec) sdk.Int
	AmountLTE(price sdk.Dec) sdk.Int
	Orders(tick sdk.Dec) []*Order
	UpTick(tick sdk.Dec, prec int) (res sdk.Dec, found bool)
	DownTick(tick sdk.Dec, prec int) (res sdk.Dec, found bool)
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

func (mos *mergedOrderSource) Orders(tick sdk.Dec) []*Order {
	var os []*Order
	for _, source := range mos.sources {
		os = append(os, source.Orders(tick)...)
	}
	return os
}

func (mos *mergedOrderSource) UpTick(tick sdk.Dec, prec int) (res sdk.Dec, found bool) {
	for _, source := range mos.sources {
		t, f := source.UpTick(tick, prec)
		if f && (res.IsNil() || t.LT(res)) {
			res = t
			found = true
		}
	}
	return
}

func (mos *mergedOrderSource) DownTick(tick sdk.Dec, prec int) (res sdk.Dec, found bool) {
	for _, source := range mos.sources {
		t, f := source.DownTick(tick, prec)
		if f && (res.IsNil() || t.GT(res)) {
			res = t
			found = true
		}
	}
	return
}

func MergeOrderSources(sources ...OrderSource) OrderSource {
	return &mergedOrderSource{sources: sources}
}

type OrderBook struct {
	prec  int // price tick precision
	buys  OrderSource
	sells OrderSource
}

func (ob OrderBook) EstimatedPriceDirection(lastPrice sdk.Dec) PriceDirection {
	if ob.buys.AmountGTE(lastPrice).ToDec().GTE(lastPrice.MulInt(ob.sells.AmountLTE(lastPrice))) {
		return PriceIncreasing
	}
	return PriceDecreasing
}

func (ob OrderBook) SwapPrice(lastPrice sdk.Dec) sdk.Dec {
	dir := ob.EstimatedPriceDirection(lastPrice)

	os := MergeOrderSources(ob.buys, ob.sells) // temporary order source just for ticks
	tick := PriceToTick(lastPrice, ob.prec)
	lowestTick := LowestTick(ob.prec)

	for {
		ba := ob.buys.AmountGTE(tick)
		sa := tick.MulInt(ob.sells.AmountLTE(tick)).TruncateInt()

		var next sdk.Dec
		var found bool
		switch dir {
		case PriceIncreasing:
			if ba.GTE(sa) { // TODO: check this condition(is GTE correct?)
				return tick
			}
			// TODO: check if there is no up tick?
			next, found = os.UpTick(tick, ob.prec)
		case PriceDecreasing:
			if sa.GTE(ba) { // TODO: check this condition(is LTE correct?)
				return tick
			}

			if tick.Equal(lowestTick) {
				return tick
			}

			next, found = os.DownTick(tick, ob.prec)
		}
		if !found {
			return tick
		}
		tick = next
	}
}
