package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/lending/types"
)

// GetParams returns the parameters for the module.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return
}

// SetParams sets the parameters for the module.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetLendingAssetParams returns the current lending asset parameters.
func (k Keeper) GetLendingAssetParams(ctx sdk.Context) (assetParams []types.LendingAssetParam) {
	k.paramSpace.Get(ctx, types.KeyLendingAssetParams, &assetParams)
	return
}

// SetLendingAssetParams sets the lending asset parameters.
func (k Keeper) SetLendingAssetParams(ctx sdk.Context, assetParams []types.LendingAssetParam) {
	k.paramSpace.Set(ctx, types.KeyLendingAssetParams, assetParams)
}
