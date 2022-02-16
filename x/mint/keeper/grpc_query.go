package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmosquad-labs/squad/x/mint/types"
)

var _ types.QueryServer = Keeper{}

// Params returns params of the mint module.
func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Inflation returns params of the mint module.
func (k Keeper) Inflation(_ context.Context, _ *types.QueryInflationRequest) (*types.QueryInflationResponse, error) {
	return &types.QueryInflationResponse{InflationSchedules: k.inflationSchedules}, nil
}
