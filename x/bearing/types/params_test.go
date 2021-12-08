package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/farming/x/bearing/types"
)

var (
	dAddr1   = sdk.AccAddress(address.Module(types.ModuleName, []byte("destinationAddr1")))
	dAddr2   = sdk.AccAddress(address.Module(types.ModuleName, []byte("destinationAddr2")))
	sAddr1   = sdk.AccAddress(address.Module(types.ModuleName, []byte("sourceAddr1")))
	sAddr2   = sdk.AccAddress(address.Module(types.ModuleName, []byte("sourceAddr2")))
	bearings = []types.Bearing{
		{
			Name:               "test",
			Rate:               sdk.OneDec(),
			SourceAddress:      sAddr1.String(),
			DestinationAddress: dAddr1.String(),
			StartTime:          types.MustParseRFC3339("2021-08-01T00:00:00Z"),
			EndTime:            types.MustParseRFC3339("2021-08-03T00:00:00Z"),
		},
		{
			Name:               "test1",
			Rate:               sdk.OneDec(),
			SourceAddress:      sAddr2.String(),
			DestinationAddress: dAddr2.String(),
			StartTime:          types.MustParseRFC3339("2021-07-01T00:00:00Z"),
			EndTime:            types.MustParseRFC3339("2021-07-10T00:00:00Z"),
		},
		{
			Name:               "test2",
			Rate:               sdk.MustNewDecFromStr("0.1"),
			SourceAddress:      sAddr2.String(),
			DestinationAddress: dAddr2.String(),
			StartTime:          types.MustParseRFC3339("2021-07-01T00:00:00Z"),
			EndTime:            types.MustParseRFC3339("2021-07-10T00:00:00Z"),
		},
		{
			Name:               "test3",
			Rate:               sdk.MustNewDecFromStr("0.1"),
			SourceAddress:      sAddr2.String(),
			DestinationAddress: dAddr2.String(),
			StartTime:          types.MustParseRFC3339("2021-08-01T00:00:00Z"),
			EndTime:            types.MustParseRFC3339("2021-08-10T00:00:00Z"),
		},
		{
			Name:               "test4",
			Rate:               sdk.OneDec(),
			SourceAddress:      sAddr2.String(),
			DestinationAddress: dAddr2.String(),
			StartTime:          types.MustParseRFC3339("2021-08-01T00:00:00Z"),
			EndTime:            types.MustParseRFC3339("2021-08-20T00:00:00Z"),
		},
		{
			Name:               "test5",
			Rate:               sdk.MustNewDecFromStr("0.1"),
			SourceAddress:      sAddr2.String(),
			DestinationAddress: dAddr2.String(),
			StartTime:          types.MustParseRFC3339("2021-08-19T00:00:00Z"),
			EndTime:            types.MustParseRFC3339("2021-08-25T00:00:00Z"),
		},
	}
)

func TestParams(t *testing.T) {
	require.IsType(t, paramstypes.KeyTable{}, types.ParamKeyTable())

	defaultParams := types.DefaultParams()

	paramsStr := `epoch_blocks: 1
bearings: []
`
	require.Equal(t, paramsStr, defaultParams.String())
}

func TestValidateBearings(t *testing.T) {
	err := types.ValidateBearings([]types.Bearing{bearings[0], bearings[1]})
	require.NoError(t, err)

	err = types.ValidateBearings([]types.Bearing{bearings[0], bearings[1], bearings[2]})
	require.ErrorIs(t, err, types.ErrInvalidTotalBearingRate)

	err = types.ValidateBearings([]types.Bearing{bearings[1], bearings[4]})
	require.NoError(t, err)

	err = types.ValidateBearings([]types.Bearing{bearings[4], bearings[5]})
	require.ErrorIs(t, err, types.ErrInvalidTotalBearingRate)

	err = types.ValidateBearings([]types.Bearing{bearings[3], bearings[3]})
	require.ErrorIs(t, err, types.ErrDuplicateBearingName)
}

func TestCollectibleBearings(t *testing.T) {
	collectibleBearings := types.CollectibleBearings([]types.Bearing{bearings[0], bearings[1]}, types.MustParseRFC3339("2021-07-05T00:00:00Z"))
	require.Len(t, collectibleBearings, 1)

	collectibleBearings = types.CollectibleBearings([]types.Bearing{bearings[0], bearings[1], bearings[2]}, types.MustParseRFC3339("2021-07-05T00:00:00Z"))
	require.Len(t, collectibleBearings, 2)

	collectibleBearings = types.CollectibleBearings([]types.Bearing{bearings[4], bearings[5]}, types.MustParseRFC3339("2021-08-18T00:00:00Z"))
	require.Len(t, collectibleBearings, 1)

	collectibleBearings = types.CollectibleBearings([]types.Bearing{bearings[4], bearings[5]}, types.MustParseRFC3339("2021-08-19T00:00:00Z"))
	require.Len(t, collectibleBearings, 2)

	collectibleBearings = types.CollectibleBearings([]types.Bearing{bearings[4], bearings[5]}, types.MustParseRFC3339("2021-08-20T00:00:00Z"))
	require.Len(t, collectibleBearings, 1)
}

func TestValidateEpochBlocks(t *testing.T) {
	err := types.ValidateEpochBlocks(uint32(0))
	require.NoError(t, err)

	err = types.ValidateEpochBlocks(nil)
	require.EqualError(t, err, "invalid parameter type: <nil>")

	err = types.ValidateEpochBlocks(types.DefaultEpochBlocks)
	require.NoError(t, err)

	err = types.ValidateEpochBlocks(10000000000000000)
	require.EqualError(t, err, "invalid parameter type: int")
}
