package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/farming/x/liquidity/types"
)

var tickPrec = int(types.DefaultTickPrecision)

func Test_fitTick(t *testing.T) {
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
		require.True(sdk.DecEq(t, tc.expected, types.PriceToTick(tc.price, tickPrec)))
	}
}

func TestTick(t *testing.T) {
	for _, tc := range []struct {
		i        int
		prec     int
		expected sdk.Dec
	}{
		{0, tickPrec, sdk.NewDecWithPrec(1, int64(sdk.Precision-tickPrec))},
		{1, tickPrec, sdk.MustNewDecFromStr("0.000000000000001001")},
		{8999, tickPrec, sdk.MustNewDecFromStr("0.000000000000009999")},
		{9000, tickPrec, sdk.MustNewDecFromStr("0.000000000000010000")},
		{9001, tickPrec, sdk.MustNewDecFromStr("0.000000000000010010")},
		{17999, tickPrec, sdk.MustNewDecFromStr("0.000000000000099990")},
		{18000, tickPrec, sdk.MustNewDecFromStr("0.000000000000100000")},
		{135000, tickPrec, sdk.NewDec(1)},
		{135001, tickPrec, sdk.MustNewDecFromStr("1.001")},
	} {
		t.Run("", func(t *testing.T) {
			res := types.TickFromIndex(tc.i, tc.prec)
			require.True(sdk.DecEq(t, tc.expected, res))
			require.Equal(t, tc.i, types.TickToIndex(res, tc.prec))
		})
	}
}
