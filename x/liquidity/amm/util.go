package amm

import (
	"fmt"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/cosmosquad-labs/squad/types"
)

// OfferCoinAmount returns the minimum offer coin amount for
// given order direction, price and order amount.
func OfferCoinAmount(dir OrderDirection, price sdk.Dec, amt sdk.Int) sdk.Int {
	switch dir {
	case Buy:
		return price.MulInt(amt).Ceil().TruncateInt()
	case Sell:
		return amt
	default:
		panic(fmt.Sprintf("invalid order direction: %s", dir))
	}
}

// sortTicks sorts given ticks in descending order.
func sortTicks(ticks []sdk.Dec) {
	sort.Slice(ticks, func(i, j int) bool {
		return ticks[i].GT(ticks[j])
	})
}

func CheckPoolPriceAfterDeposit(beforeX, beforeY, afterX, afterY sdk.Int) {
	beforePrice := beforeX.ToDec().Quo(beforeY.ToDec())
	afterPrice := afterX.ToDec().Quo(afterY.ToDec())
	diff := beforePrice.Sub(afterPrice).Abs()
	if diff.GT(utils.ParseDec("0.0001")) {
		panic(fmt.Errorf("before=%s/%s(%s), after=%s/%s(%s), diff=%s",
			beforeX, beforeY, beforePrice, afterX, afterY, afterPrice, diff))
	}
}
