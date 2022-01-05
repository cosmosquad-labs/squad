package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_tickInterval(t *testing.T) {
	defTickPrec := int(DefaultTickPrecision)

	require.Panics(t, func() {
		tickInterval(sdk.ZeroDec(), int(DefaultTickPrecision))
	})

	for _, tc := range []struct {
		price    sdk.Dec
		prec     int
		expected sdk.Dec
	}{
		{sdk.MustNewDecFromStr("10000"), defTickPrec, sdk.MustNewDecFromStr("10")},
		{sdk.MustNewDecFromStr("9999"), defTickPrec, sdk.MustNewDecFromStr("1")},
		{sdk.MustNewDecFromStr("10"), defTickPrec, sdk.MustNewDecFromStr("0.01")},
		{sdk.MustNewDecFromStr("9.999999999999999999"), defTickPrec, sdk.MustNewDecFromStr("0.001")},
		{sdk.MustNewDecFromStr("9"), defTickPrec, sdk.MustNewDecFromStr("0.001")},
		{sdk.MustNewDecFromStr("1"), defTickPrec, sdk.MustNewDecFromStr("0.001")},
		{sdk.MustNewDecFromStr("0.9"), defTickPrec, sdk.MustNewDecFromStr("0.0001")},
	} {
		require.True(sdk.DecEq(t, tc.expected, tickInterval(tc.price, tc.prec)))
	}
}

func Test_fitToTick(t *testing.T) {
	for _, tc := range []struct {
		price    sdk.Dec
		expected sdk.Dec
	}{
		{sdk.MustNewDecFromStr("0.000000000000099999"), sdk.MustNewDecFromStr("0.00000000000009999")},
		{sdk.MustNewDecFromStr("1.999999999999999999"), sdk.MustNewDecFromStr("1.999")},
		{sdk.MustNewDecFromStr("99.999999999999999999"), sdk.MustNewDecFromStr("99.99")},
		{sdk.MustNewDecFromStr("100.999999999999999999"), sdk.MustNewDecFromStr("100.9")},
		{sdk.MustNewDecFromStr("9999.999999999999999999"), sdk.MustNewDecFromStr("9999")},
		{sdk.MustNewDecFromStr("10019"), sdk.MustNewDecFromStr("10010")},
		{sdk.MustNewDecFromStr("1000100005"), sdk.MustNewDecFromStr("1000000000")},
	} {
		require.True(sdk.DecEq(t, tc.expected, fitToTick(tc.price, int(DefaultTickPrecision))))
	}
}

func Test_upTick(t *testing.T) {
	defTickPrec := int(DefaultTickPrecision)

	for _, tc := range []struct {
		price    sdk.Dec
		prec     int
		expected sdk.Dec
	}{
		{sdk.MustNewDecFromStr("1000"), defTickPrec, sdk.MustNewDecFromStr("1001")},
		{sdk.MustNewDecFromStr("999.9"), defTickPrec, sdk.MustNewDecFromStr("1000")},
		{sdk.MustNewDecFromStr("999"), defTickPrec, sdk.MustNewDecFromStr("999.1")},
		{sdk.MustNewDecFromStr("1.1"), defTickPrec, sdk.MustNewDecFromStr("1.101")},
		{sdk.MustNewDecFromStr("1"), defTickPrec, sdk.MustNewDecFromStr("1.001")},
		{sdk.MustNewDecFromStr("0.999999999999999999"), defTickPrec, sdk.MustNewDecFromStr("1")},
		{sdk.MustNewDecFromStr("0.1"), defTickPrec, sdk.MustNewDecFromStr("0.1001")},
		{sdk.MustNewDecFromStr("0.09999"), defTickPrec, sdk.MustNewDecFromStr("0.1")},
		{sdk.MustNewDecFromStr("0.09997"), defTickPrec, sdk.MustNewDecFromStr("0.09998")},
	} {
		require.True(sdk.DecEq(t, tc.expected, upTick(tc.price, tc.prec)))
	}
}

func Test_downTick(t *testing.T) {
	defTickPrec := int(DefaultTickPrecision)

	require.Panics(t, func() {
		downTick(sdk.MustNewDecFromStr("0.000000000000001000"), defTickPrec)
	})

	for _, tc := range []struct {
		price    sdk.Dec
		prec     int
		expected sdk.Dec
	}{
		{sdk.MustNewDecFromStr("10010"), defTickPrec, sdk.MustNewDecFromStr("10000")},
		{sdk.MustNewDecFromStr("100"), defTickPrec, sdk.MustNewDecFromStr("99.99")},
		{sdk.MustNewDecFromStr("99.99"), defTickPrec, sdk.MustNewDecFromStr("99.98")},
		{sdk.MustNewDecFromStr("1"), defTickPrec, sdk.MustNewDecFromStr("0.9999")},
		{sdk.MustNewDecFromStr("0.999"), defTickPrec, sdk.MustNewDecFromStr("0.9989")},
		{sdk.MustNewDecFromStr("0.99999"), defTickPrec, sdk.MustNewDecFromStr("0.9999")},
		{sdk.MustNewDecFromStr("0.1"), defTickPrec, sdk.MustNewDecFromStr("0.09999")},
		{sdk.MustNewDecFromStr("0.00000000000001"), defTickPrec, sdk.MustNewDecFromStr("0.000000000000009999")},
		{sdk.MustNewDecFromStr("0.000000000000001001"), defTickPrec, sdk.MustNewDecFromStr("0.000000000000001000")},
	} {
		require.True(sdk.DecEq(t, tc.expected, downTick(tc.price, tc.prec)))
	}
}
