package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/lending/types"
)

// GetLendingAsset returns a specific lending asset by its denom.
func (k Keeper) GetLendingAsset(ctx sdk.Context, denom string) (asset types.LendingAsset, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLendingAssetKey(denom))
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshal(bz, &asset)
	return asset, true
}

// SetLendingAsset stores a lending asset.
func (k Keeper) SetLendingAsset(ctx sdk.Context, asset types.LendingAsset) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&asset)
	store.Set(types.GetLendingAssetKey(asset.Denom), bz)
}
