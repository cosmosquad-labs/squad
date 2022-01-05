package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	oneDec = sdk.OneDec()
	tenDec = sdk.NewDec(10)
)

// log10f returns floor(log10(x)).
func log10f(x sdk.Dec) int {
	if x.IsZero() {
		panic("x must not be zero")
	}
	ret := 0
	if x.GTE(oneDec) {
		for ; x.GTE(oneDec); ret++ {
			x = x.QuoTruncate(tenDec)
		}
		return ret - 1
	}
	for ; x.LT(oneDec); ret-- {
		x = x.Mul(tenDec)
	}
	return ret
}

func pow10(power int) sdk.Dec {
	switch {
	case power == 0:
		return sdk.OneDec()
	case power > 0:
		return sdk.NewDec(10).Power(uint64(power))
	default:
		return sdk.OneDec().QuoTruncate(sdk.NewDec(10).Power(uint64(-power)))
	}
}
