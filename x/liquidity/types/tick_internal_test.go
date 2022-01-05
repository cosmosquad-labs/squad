package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestFitToTick(t *testing.T) {
	for _, tc := range []struct {
		price sdk.Dec
		ret   sdk.Dec
	}{
		{sdk.MustNewDecFromStr("0.000000000000099999"), sdk.MustNewDecFromStr("0.00000000000009999")},
		{sdk.MustNewDecFromStr("1.999999999999999999"), sdk.MustNewDecFromStr("1.999")},
		{sdk.MustNewDecFromStr("99.999999999999999999"), sdk.MustNewDecFromStr("99.99")},
		{sdk.MustNewDecFromStr("100.999999999999999999"), sdk.MustNewDecFromStr("100.9")},
		{sdk.MustNewDecFromStr("9999.999999999999999999"), sdk.MustNewDecFromStr("9999")},
		{sdk.MustNewDecFromStr("10019"), sdk.MustNewDecFromStr("10010")},
		{sdk.MustNewDecFromStr("1000100005"), sdk.MustNewDecFromStr("1000000000")},
	} {
		require.True(sdk.DecEq(t, tc.ret, fitToTick(tc.price, int(DefaultTickPrecision))))
	}
}
