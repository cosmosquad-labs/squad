package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/x/liquidfarming/types"
)

// GetLastRewardsAuctionId returns the last auction id.
func (k Keeper) GetLastRewardsAuctionId(ctx sdk.Context) uint64 {
	var id uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastRewardsAuctionIdKey)
	if bz == nil {
		id = 0 // initialize the auction id
	} else {
		val := gogotypes.UInt64Value{}
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return id
}

// SetRewardsLastAuctionId stores the last auction id.
func (k Keeper) SetRewardsLastAuctionId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.LastRewardsAuctionIdKey, bz)
}

// GetDepositRequest returns the particular deposit request.
func (k Keeper) GetDepositRequest(ctx sdk.Context, reqId, poolId uint64) (req types.DepositRequest, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetDepositRequestKey(reqId, poolId))
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
	store.Set(types.GetDepositRequestKey(req.Id, req.PoolId), bz)
}

// SetDepositRequestIndex stores the deposit request index.
func (k Keeper) SetDepositRequestIndex(ctx sdk.Context, req types.DepositRequest) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetDepositRequestIndexKey(req.GetDepositor(), req.PoolId, req.Id), []byte{})
}

func (k Keeper) GetRewardsAuction(ctx sdk.Context, poolId, auctionId uint64) (auction types.RewardsAuction, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetRewardsAuctionKey(poolId, auctionId))
	if bz == nil {
		return auction, false
	}

	auction = types.MustUnmarshalRewardsAuction(k.cdc, bz)

	return auction, true
}

func (k Keeper) SetRewardsAuction(ctx sdk.Context, auction types.RewardsAuction) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalRewardsAuction(k.cdc, auction)
	store.Set(types.GetRewardsAuctionKey(auction.PoolId, auction.Id), bz)
}

// GetRewardsAuctions returns all auctions in the store.
func (k Keeper) GetRewardsAuctions(ctx sdk.Context) (auctions []types.RewardsAuction) {
	k.IterateRewardsAuctions(ctx, func(auction types.RewardsAuction) (stop bool) {
		auctions = append(auctions, auction)
		return false
	})
	return auctions
}

// IterateRewardsAuctions iterates over all the stored auctions and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateRewardsAuctions(ctx sdk.Context, cb func(auction types.RewardsAuction) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AuctionKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		auction := types.MustUnmarshalRewardsAuction(k.cdc, iterator.Value())

		if cb(auction) {
			break
		}
	}
}
