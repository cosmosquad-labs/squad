package keeper_test

import (
	_ "github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/farming/x/bearing/types"
)

func (suite *KeeperTestSuite) TestGRPCParams() {
	resp, err := suite.querier.Params(sdk.WrapSDKContext(suite.ctx), &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(suite.keeper.GetParams(suite.ctx), resp.Params)
}

func (suite *KeeperTestSuite) TestGRPCBearings() {
	bearings := []types.Bearing{
		{
			Name:               "bearing1",
			Rate:               sdk.NewDecWithPrec(5, 2),
			SourceAddress:      suite.sourceAddrs[0].String(),
			DestinationAddress: suite.destinationAddrs[0].String(),
			StartTime:          types.MustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:            types.MustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:               "bearing2",
			Rate:               sdk.NewDecWithPrec(5, 2),
			SourceAddress:      suite.sourceAddrs[0].String(),
			DestinationAddress: suite.destinationAddrs[1].String(),
			StartTime:          types.MustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:            types.MustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:               "bearing3",
			Rate:               sdk.NewDecWithPrec(5, 2),
			SourceAddress:      suite.sourceAddrs[1].String(),
			DestinationAddress: suite.destinationAddrs[0].String(),
			StartTime:          types.MustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:            types.MustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:               "bearing4",
			Rate:               sdk.NewDecWithPrec(5, 2),
			SourceAddress:      suite.sourceAddrs[1].String(),
			DestinationAddress: suite.destinationAddrs[1].String(),
			StartTime:          types.MustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:            types.MustParseRFC3339("9999-12-31T00:00:00Z"),
		},
	}

	params := suite.keeper.GetParams(suite.ctx)
	params.Bearings = bearings
	suite.keeper.SetParams(suite.ctx, params)

	balance := suite.app.BankKeeper.GetAllBalances(suite.ctx, suite.sourceAddrs[0])
	expectedCoins, _ := sdk.NewDecCoinsFromCoins(balance...).MulDec(sdk.NewDecWithPrec(5, 2)).TruncateDecimal()

	suite.ctx = suite.ctx.WithBlockTime(types.MustParseRFC3339("2021-08-31T00:00:00Z"))
	err := suite.keeper.CollectBearings(suite.ctx)
	suite.Require().NoError(err)

	for _, tc := range []struct {
		name      string
		req       *types.QueryBearingsRequest
		expectErr bool
		postRun   func(response *types.QueryBearingsResponse)
	}{
		{
			"nil request",
			nil,
			true,
			nil,
		},
		{
			"query all",
			&types.QueryBearingsRequest{},
			false,
			func(resp *types.QueryBearingsResponse) {
				suite.Require().Len(resp.Bearings, 4)
			},
		},
		{
			"query by not existing name",
			&types.QueryBearingsRequest{Name: "notfound"},
			false,
			func(resp *types.QueryBearingsResponse) {
				suite.Require().Len(resp.Bearings, 0)
			},
		},
		{
			"query by name",
			&types.QueryBearingsRequest{Name: "bearing1"},
			false,
			func(resp *types.QueryBearingsResponse) {
				suite.Require().Len(resp.Bearings, 1)
				suite.Require().Equal("bearing1", resp.Bearings[0].Bearing.Name)
			},
		},
		{
			"invalid source addr",
			&types.QueryBearingsRequest{SourceAddress: "invalid"},
			true,
			nil,
		},
		{
			"query by source addr",
			&types.QueryBearingsRequest{SourceAddress: suite.sourceAddrs[0].String()},
			false,
			func(resp *types.QueryBearingsResponse) {
				suite.Require().Len(resp.Bearings, 2)
				for _, b := range resp.Bearings {
					suite.Require().Equal(suite.sourceAddrs[0].String(), b.Bearing.SourceAddress)
				}
			},
		},
		{
			"invalid destination addr",
			&types.QueryBearingsRequest{DestinationAddress: "invalid"},
			true,
			nil,
		},
		{
			"query by destination addr",
			&types.QueryBearingsRequest{DestinationAddress: suite.destinationAddrs[0].String()},
			false,
			func(resp *types.QueryBearingsResponse) {
				suite.Require().Len(resp.Bearings, 2)
				for _, b := range resp.Bearings {
					suite.Require().Equal(suite.destinationAddrs[0].String(), b.Bearing.DestinationAddress)
				}
			},
		},
		{
			"query with multiple filters",
			&types.QueryBearingsRequest{
				SourceAddress:      suite.sourceAddrs[0].String(),
				DestinationAddress: suite.destinationAddrs[1].String(),
			},
			false,
			func(resp *types.QueryBearingsResponse) {
				suite.Require().Len(resp.Bearings, 1)
				suite.Require().Equal(suite.sourceAddrs[0].String(), resp.Bearings[0].Bearing.SourceAddress)
				suite.Require().Equal(suite.destinationAddrs[1].String(), resp.Bearings[0].Bearing.DestinationAddress)
			},
		},
		{
			"correct total collected coins",
			&types.QueryBearingsRequest{Name: "bearing1"},
			false,
			func(resp *types.QueryBearingsResponse) {
				suite.Require().Len(resp.Bearings, 1)
				suite.Require().True(coinsEq(expectedCoins, resp.Bearings[0].TotalCollectedCoins))
			},
		},
	} {
		suite.Run(tc.name, func() {
			resp, err := suite.querier.Bearings(sdk.WrapSDKContext(suite.ctx), tc.req)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				tc.postRun(resp)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGRPCAddresses() {
	for _, tc := range []struct {
		name         string
		req          *types.QueryAddressesRequest
		expectedAddr string
		expectErr    bool
	}{
		{
			"nil request",
			nil,
			"",
			true,
		},
		{
			"empty request",
			&types.QueryAddressesRequest{},
			"",
			true,
		},
		{
			"default module name and address type",
			&types.QueryAddressesRequest{Name: "testSourceAddr"},
			"cosmos1hg0v9u92ztzecpmml26206wwtghggx0flpwn5d4qc3r6dvuanxeqs4mnk5",
			false,
		},
		{
			"invalid address type",
			&types.QueryAddressesRequest{Name: "testSourceAddr", Type: 2},
			"",
			true,
		},
	} {
		suite.Run(tc.name, func() {
			resp, err := suite.querier.Addresses(sdk.WrapSDKContext(suite.ctx), tc.req)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(resp.Address, tc.expectedAddr)
			}
		})
	}
}
