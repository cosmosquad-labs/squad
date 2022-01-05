package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestLog10f(t *testing.T) {
	require.Panics(t, func() {
		log10f(sdk.ZeroDec())
	})

	for _, tc := range []struct {
		x   sdk.Dec
		ret int
	}{
		{sdk.MustNewDecFromStr("999.99999999999999999"), 2},
		{sdk.MustNewDecFromStr("100"), 2},
		{sdk.MustNewDecFromStr("99.999999999999999999"), 1},
		{sdk.MustNewDecFromStr("10"), 1},
		{sdk.MustNewDecFromStr("9.999999999999999999"), 0},
		{sdk.MustNewDecFromStr("1"), 0},
		{sdk.MustNewDecFromStr("0.999999999999999999"), -1},
		{sdk.MustNewDecFromStr("0.1"), -1},
		{sdk.MustNewDecFromStr("0.099999999999999999"), -2},
		{sdk.MustNewDecFromStr("0.01"), -2},
		{sdk.MustNewDecFromStr("0.000000000000000009"), -18},
		{sdk.MustNewDecFromStr("0.000000000000000001"), -18},
	} {
		require.Equal(t, tc.ret, log10f(tc.x))
	}
}
