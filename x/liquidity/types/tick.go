package types

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func fitToTick(price sdk.Dec, prec int) sdk.Dec {
	pow := prec - log10f(price)
	t := sdk.NewDec(10).Power(uint64(math.Abs(float64(pow))))
	if pow >= 0 {
		return price.MulTruncate(t).TruncateDec().QuoTruncate(t)
	}
	return price.QuoTruncate(t).TruncateDec().MulTruncate(t)
}
