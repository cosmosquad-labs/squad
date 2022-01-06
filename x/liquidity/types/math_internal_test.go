package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_log10f(t *testing.T) {
	require.Panics(t, func() {
		log10f(sdk.ZeroDec())
	})

	for _, tc := range []struct {
		x        sdk.Dec
		expected int
	}{
		{sdk.MustNewDecFromStr("999.99999999999999999"), 20},
		{sdk.MustNewDecFromStr("100"), 20},
		{sdk.MustNewDecFromStr("99.999999999999999999"), 19},
		{sdk.MustNewDecFromStr("10"), 19},
		{sdk.MustNewDecFromStr("9.999999999999999999"), 18},
		{sdk.MustNewDecFromStr("1"), 18},
		{sdk.MustNewDecFromStr("0.999999999999999999"), 17},
		{sdk.MustNewDecFromStr("0.1"), 17},
		{sdk.MustNewDecFromStr("0.099999999999999999"), 16},
		{sdk.MustNewDecFromStr("0.01"), 16},
		{sdk.MustNewDecFromStr("0.000000000000000009"), 0},
		{sdk.MustNewDecFromStr("0.000000000000000001"), 0},
	} {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tc.expected, log10f(tc.x))
		})
	}
}

func Test_pow10(t *testing.T) {
	for _, tc := range []struct {
		power    int
		expected sdk.Dec
	}{
		{18, sdk.NewDec(1)},
		{19, sdk.NewDec(10)},
		{20, sdk.NewDec(100)},
		{17, sdk.NewDecWithPrec(1, 1)},
		{16, sdk.NewDecWithPrec(1, 2)},
	} {
		t.Run("", func(t *testing.T) {
			require.True(sdk.DecEq(t, tc.expected, pow10(tc.power)))
		})
	}
}
