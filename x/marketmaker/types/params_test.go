package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmosquad-labs/squad/v2/x/marketmaker/types"
)

func TestParams(t *testing.T) {
	require.IsType(t, paramstypes.KeyTable{}, types.ParamKeyTable())

	defaultParams := types.DefaultParams()

	paramsStr := `incentive_budget_address: cosmos1ddn66jv0sjpmck0ptegmhmqtn35qsg2vxyk2hn9sqf4qxtzqz3sqanrtcm
deposit_amount:
- denom: stake
  amount: "1000000000"
incentive_pairs: []
`

	require.Equal(t, paramsStr, defaultParams.String())
}

//func TestParamsValidate(t *testing.T) {
//	require.NoError(t, types.DefaultParams().Validate())
//
//	testCases := []struct {
//		name        string
//		configure   func(*types.Params)
//		expectedErr string
//	}{
//		{
//			"EmptyPrivatePlanCreationFee",
//			func(params *types.Params) {
//				params.PrivatePlanCreationFee = sdk.NewCoins()
//			},
//			"",
//		},
//		{
//			"ZeroNextEpochDays",
//			func(params *types.Params) {
//				params.NextEpochDays = uint32(0)
//			},
//			"next epoch days must be positive: 0",
//		},
//		{
//			"EmptyFarmingFeeCollector",
//			func(params *types.Params) {
//				params.FarmingFeeCollector = ""
//			},
//			"farming fee collector address must not be empty",
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			params := types.DefaultParams()
//			tc.configure(&params)
//			err := params.Validate()
//
//			var err2 error
//			for _, p := range params.ParamSetPairs() {
//				err := p.ValidatorFn(reflect.ValueOf(p.Value).Elem().Interface())
//				if err != nil {
//					err2 = err
//					break
//				}
//			}
//			if tc.expectedErr != "" {
//				require.EqualError(t, err, tc.expectedErr)
//				require.EqualError(t, err2, tc.expectedErr)
//			} else {
//				require.Nil(t, err)
//				require.Nil(t, err2)
//			}
//		})
//	}
//}
