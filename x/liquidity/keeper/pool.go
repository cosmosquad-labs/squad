package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/crescent-network/crescent/x/liquidity/types"
)

// GetPoolId returns the global pool id counter.
func (k Keeper) GetPoolId(ctx sdk.Context) uint64 {
	var id uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.PoolIdKey)
	if bz == nil {
		id = 0 // initialize the pool id
	} else {
		val := gogotypes.UInt64Value{}
		err := k.cdc.Unmarshal(bz, &val)
		if err != nil {
			panic(err)
		}
		id = val.GetValue()
	}
	return id
}

// GetPool returns pool object for the given pool id.
func (k Keeper) GetPool(ctx sdk.Context, id uint64) (pool types.Pool, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPoolKey(id)

	value := store.Get(key)
	if value == nil {
		return pool, false
	}

	pool = types.MustUnmarshalPool(k.cdc, value)

	return pool, true
}

// GetAllPools returns all pairs in the store.
func (k Keeper) GetAllPools(ctx sdk.Context) (pools []types.Pool) {
	k.IterateAllPools(ctx, func(pool types.Pool) (stop bool) {
		pools = append(pools, pool)
		return false
	})

	return pools
}

// GetNextPoolIdWithUpdate increments pool id by one and set it.
func (k Keeper) GetNextPoolIdWithUpdate(ctx sdk.Context) uint64 {
	id := k.GetPoolId(ctx) + 1
	k.SetPoolId(ctx, id)
	return id
}

// SetPoolId stores the global pool id counter.
func (k Keeper) SetPoolId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.PoolIdKey, bz)
}

// SetPool stores the particular pool.
func (k Keeper) SetPool(ctx sdk.Context, pool types.Pool) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalPool(k.cdc, pool)
	store.Set(types.GetPoolKey(pool.Id), b)
}

// SetPoolByPair stores a pool by pair.
func (k Keeper) SetPoolByPair(ctx sdk.Context, pairId uint64, poolId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetPoolsByPairIndexKey(pairId, poolId), []byte{})
}

// IterateAllPools iterates over all the stored pools and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAllPools(ctx sdk.Context, cb func(pool types.Pool) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PoolKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		pool := types.MustUnmarshalPool(k.cdc, iterator.Value())
		if cb(pool) {
			break
		}
	}
}

// IterateAllPools iterates over all the stored pools by the pair and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IteratePoolsByPair(ctx sdk.Context, pairId uint64, cb func(pool types.Pool) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.GetPoolsByPairKey(pairId))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		poolId := types.ParsePoolsByPairIndexKey(iterator.Key())
		pool, _ := k.GetPool(ctx, poolId)
		if cb(pool) {
			break
		}
	}
}
