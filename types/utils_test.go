package types_test

import (
	"math/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/cosmosquad-labs/squad/types"
)

func TestGetShareValue(t *testing.T) {
	require.EqualValues(t, types.GetShareValue(sdk.NewInt(100), sdk.MustNewDecFromStr("0.9")), sdk.NewInt(90))
	require.EqualValues(t, types.GetShareValue(sdk.NewInt(100), sdk.MustNewDecFromStr("1.1")), sdk.NewInt(110))

	// truncated
	require.EqualValues(t, types.GetShareValue(sdk.NewInt(101), sdk.MustNewDecFromStr("0.9")), sdk.NewInt(90))
	require.EqualValues(t, types.GetShareValue(sdk.NewInt(101), sdk.MustNewDecFromStr("1.1")), sdk.NewInt(111))

	require.EqualValues(t, types.GetShareValue(sdk.NewInt(100), sdk.MustNewDecFromStr("0")), sdk.NewInt(0))
	require.EqualValues(t, types.GetShareValue(sdk.NewInt(0), sdk.MustNewDecFromStr("1.1")), sdk.NewInt(0))
}

func TestAddOrInit(t *testing.T) {
	strIntMap := make(types.StrIntMap)

	// Set when the key not existed on the map
	strIntMap.AddOrSet("a", sdk.NewInt(1))
	require.Equal(t, strIntMap["a"], sdk.NewInt(1))

	// Added when the key existed on the map
	strIntMap.AddOrSet("a", sdk.NewInt(1))
	require.Equal(t, strIntMap["a"], sdk.NewInt(2))
}

func TestParseTime(t *testing.T) {
	normalCase := "9999-12-31T00:00:00Z"
	normalRes, err := time.Parse(time.RFC3339, normalCase)
	require.NoError(t, err)
	errorCase := "9999-12-31T00:00:00_ErrorCase"
	_, err = time.Parse(time.RFC3339, errorCase)
	require.PanicsWithError(t, err.Error(), func() { types.ParseTime(errorCase) })
	require.Equal(t, normalRes, types.ParseTime(normalCase))
}

func TestParseDec(t *testing.T) {
	require.True(sdk.DecEq(t, sdk.NewDec(1), types.ParseDec("1.0")))
	require.True(sdk.DecEq(t, sdk.NewDecWithPrec(5, 2), types.ParseDec("0.05")))
	require.Panics(t, func() {
		types.ParseDec("1.1.1")
	})
}

func TestParseCoin(t *testing.T) {
	require.True(t, sdk.NewInt64Coin("denom1", 1000000).IsEqual(types.ParseCoin("1000000denom1")))
	require.Panics(t, func() {
		types.ParseCoin("1000000")
	})
}

func TestParseCoins(t *testing.T) {
	coins := sdk.NewCoins(
		sdk.NewInt64Coin("denom1", 1000000),
		sdk.NewInt64Coin("denom2", 2000000),
	)
	require.True(t, coins.IsEqual(types.ParseCoins("1000000denom1,2000000denom2")))
}

func TestDecApproxEqual(t *testing.T) {
	for _, tc := range []struct {
		a, b     sdk.Dec
		expected bool
	}{
		{types.ParseDec("1.0"), types.ParseDec("1.0"), true},
		{types.ParseDec("100.0"), types.ParseDec("100.1"), true},
		{types.ParseDec("100.0"), types.ParseDec("100.101"), false},
		{types.ParseDec("0.000001"), types.ParseDec("0.000001001"), true},
	} {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tc.expected, types.DecApproxEqual(tc.a, tc.b))
		})
	}
}

func TestDateRangesOverlap(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		startTimeA     time.Time
		endTimeA       time.Time
		startTimeB     time.Time
		endTimeB       time.Time
	}{
		{
			"not overlapping",
			false,
			types.ParseTime("2021-12-01T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-03T00:00:00Z"),
			types.ParseTime("2021-12-04T00:00:00Z"),
		},
		{
			"not overlapped on same end time and start time",
			false,
			types.ParseTime("2021-12-01T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-03T00:00:00Z"),
		},
		{
			"not overlapped on same end time and start time 2",
			false,
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-03T00:00:00Z"),
			types.ParseTime("2021-12-01T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
		},
		{
			"for the same time, it doesn't seem to overlap",
			false,
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
		},
		{
			"end time and start time differs by a little amount",
			true,
			types.ParseTime("2021-12-01T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00.01Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-03T00:00:00Z"),
		},
		{
			"overlap #1",
			true,
			types.ParseTime("2021-12-01T00:00:00Z"),
			types.ParseTime("2021-12-03T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-04T00:00:00Z"),
		},
		{
			"overlap #2 - same ranges",
			true,
			types.ParseTime("2021-12-01T00:00:00Z"),
			types.ParseTime("2021-12-03T00:00:00Z"),
			types.ParseTime("2021-12-01T00:00:00Z"),
			types.ParseTime("2021-12-03T00:00:00Z"),
		},
		{
			"overlap #3 - one includes another",
			true,
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-03T00:00:00Z"),
			types.ParseTime("2021-12-01T00:00:00Z"),
			types.ParseTime("2021-12-04T00:00:00Z"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expectedResult, types.DateRangesOverlap(tc.startTimeA, tc.endTimeA, tc.startTimeB, tc.endTimeB))
			require.Equal(t, tc.expectedResult, types.DateRangesOverlap(tc.startTimeB, tc.endTimeB, tc.startTimeA, tc.endTimeA))
		})
	}
}

func TestDateRangeIncludes(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		targetTime     time.Time
		startTime      time.Time
		endTime        time.Time
	}{
		{
			"not included, before started",
			false,
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:01Z"),
			types.ParseTime("2021-12-03T00:00:00Z"),
		},
		{
			"not included, after ended",
			false,
			types.ParseTime("2021-12-03T00:00:01Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-03T00:00:00Z"),
		},
		{
			"included on start time",
			true,
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-03T00:00:00Z"),
		},
		{
			"not included on end time",
			false,
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-01T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
		},
		{
			"not included on same start time and end time",
			false,
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
			types.ParseTime("2021-12-02T00:00:00Z"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expectedResult, types.DateRangeIncludes(tc.startTime, tc.endTime, tc.targetTime))
		})
	}
}

func TestRandomInt(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	require.True(sdk.IntEq(t, sdk.NewInt(67), types.RandomInt(r, sdk.NewInt(1), sdk.NewInt(100))))
	require.True(sdk.IntEq(t, sdk.NewInt(15), types.RandomInt(r, sdk.NewInt(1), sdk.NewInt(100))))
	require.True(sdk.IntEq(t, sdk.NewInt(3), types.RandomInt(r, sdk.NewInt(1), sdk.NewInt(100))))
}

func TestRandomDec(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	require.True(sdk.DecEq(t, types.ParseDec("49.573137597901179650"), types.RandomDec(r, types.ParseDec("0.01"), types.ParseDec("100"))))
	require.True(sdk.DecEq(t, types.ParseDec("97.794564449098792592"), types.RandomDec(r, types.ParseDec("0.01"), types.ParseDec("100"))))
	require.True(sdk.DecEq(t, types.ParseDec("7.031885751622762309"), types.RandomDec(r, types.ParseDec("0.01"), types.ParseDec("100"))))
}

func TestTestAddress(t *testing.T) {
	require.Equal(t, "cosmos1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqnrql8a", types.TestAddress(0).String())
	require.Equal(t, "cosmos16q8sqqqqqqqqqqqqqqqqqqqqqqqqqqqqn4c7e0", types.TestAddress(1000).String())
}
