package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/x/liquidfarming/types"
)

// GetLastLiquidfarmId returns the last liquid farm id.
func (k Keeper) GetLastLiquidfarmId(ctx sdk.Context) uint64 {
	var id uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastLiquidfarmIdKey)
	if bz == nil {
		id = 0 // initialize the id
	} else {
		val := gogotypes.UInt64Value{}
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return id
}

// SetLastLiquidfarmId stores the last liquid farm id.
func (k Keeper) SetLastLiquidfarmId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.LastLiquidfarmIdKey, bz)
}

// GetLastAuctionId returns the last auction id.
func (k Keeper) GetLastAuctionId(ctx sdk.Context) uint64 {
	var id uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastAuctionIdKey)
	if bz == nil {
		id = 0 // initialize the auction id
	} else {
		val := gogotypes.UInt64Value{}
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return id
}

// SetLastAuctionId stores the last auction id.
func (k Keeper) SetLastAuctionId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.LastAuctionIdKey, bz)
}

// GetDepositRequest returns the particular deposit request.
func (k Keeper) GetDepositRequest(ctx sdk.Context, liquidfarmId, id uint64) (req types.DepositRequest, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetDepositRequestKey(liquidfarmId, id))
	if bz == nil {
		return
	}
	req = types.MustUnmarshalDepositRequest(k.cdc, bz)
	return req, true
}

// SetDepositRequest stores deposit request for the batch execution.
func (k Keeper) SetDepositRequest(ctx sdk.Context, req types.DepositRequest) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalDepositRequest(k.cdc, req)
	store.Set(types.GetDepositRequestKey(req.LiquidFarmId, req.Id), bz)
}

// SetDepositRequestIndex stores the deposit request index.
func (k Keeper) SetDepositRequestIndex(ctx sdk.Context, req types.DepositRequest) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetDepositRequestIndexKey(req.GetDepositor(), req.LiquidFarmId, req.Id), []byte{})
}

func (k Keeper) GetLiquidFarm(ctx sdk.Context, id uint64) (liquidfarm types.LiquidFarm, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLiquidFarmKey(id))
	if bz == nil {
		return
	}
	liquidfarm = types.MustUnmarshalLiquidFarm(k.cdc, bz)
	return liquidfarm, true
}

func (k Keeper) SetLiquidFarm(ctx sdk.Context, liquidfarm types.LiquidFarm) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalLiquidFarm(k.cdc, liquidfarm)
	store.Set(types.GetLiquidFarmKey(liquidfarm.Id), bz)
}

// IterateAllLiquidFarms iterates over all the stored liquid farms and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAllLiquidFarms(ctx sdk.Context, cb func(liquidfarm types.LiquidFarm) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.LiquidFarmKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		pair := types.MustUnmarshalLiquidFarm(k.cdc, iter.Value())
		stop, err := cb(pair)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
}
