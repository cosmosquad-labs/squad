package types

import (
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func tickInterval(price sdk.Dec, prec int) sdk.Dec {
	return pow10(log10f(price) - prec)
}

func fitToTick(price sdk.Dec, prec int) sdk.Dec {
	pow := prec - log10f(price)
	t := pow10(int(math.Abs(float64(pow))))
	if pow >= 0 {
		return price.MulTruncate(t).TruncateDec().QuoTruncate(t)
	}
	return price.QuoTruncate(t).TruncateDec().MulTruncate(t)
}

func upTick(price sdk.Dec, prec int) sdk.Dec {
	return fitToTick(price, prec).Add(tickInterval(price, prec))
}

// TODO: optimize it
func downTick(price sdk.Dec, prec int) sdk.Dec {
	p := fitToTick(price, prec)
	if p.Equal(price) {
		if isLowestTick(p, prec) {
			panic(fmt.Errorf("%s is the lowest possible tick", price))
		}
		var i sdk.Dec
		if pow10(log10f(p)).Equal(p) {
			i = tickInterval(p.Quo(tenDec), prec)
		} else {
			i = tickInterval(p, prec)
		}
		p = p.Sub(i)
	}
	return p
}

func isLowestTick(price sdk.Dec, prec int) bool {
	l := log10f(price)
	if -l != sdk.Precision-prec {
		return false
	}
	return price.Equal(pow10(l))
}
