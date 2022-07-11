package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// GetLastDepositRequestId returns the last deposit request id for the pool id.
func (k Keeper) GetLastDepositRequestId(ctx sdk.Context, poolId uint64) uint64 {
	var id uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLastDepositRequestIdKey(poolId))
	if bz == nil {
		id = 0 // initialize the deposit request id
	} else {
		val := gogotypes.UInt64Value{}
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return id
}

// SetDepositRequestId sets the deposit request id with the given pool id.
func (k Keeper) SetDepositRequestId(ctx sdk.Context, poolId uint64, reqId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: reqId})
	store.Set(types.GetLastDepositRequestIdKey(poolId), bz)
}

// GetLastBidId returns the last bid id for the bid.
func (k Keeper) GetLastBidId(ctx sdk.Context, auctionId uint64) uint64 {
	var id uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLastBidIdKey(auctionId))
	if bz == nil {
		id = 0 // initialize the bid id
	} else {
		val := gogotypes.UInt64Value{}
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return id
}

// SetBidId sets the bid id for the auction.
func (k Keeper) SetBidId(ctx sdk.Context, auctionId uint64, bidId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: bidId})
	store.Set(types.GetLastBidIdKey(auctionId), bz)
}

// GetRewardsAuctionId returns the last auction id.
func (k Keeper) GetRewardsAuctionId(ctx sdk.Context) uint64 {
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

// SetRewardsAuctionId stores the last auction id.
func (k Keeper) SetRewardsAuctionId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.LastRewardsAuctionIdKey, bz)
}

// GetDepositRequest returns the particular deposit request.
func (k Keeper) GetDepositRequest(ctx sdk.Context, poolId, reqId uint64) (req types.DepositRequest, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetDepositRequestKey(poolId, reqId))
	if bz == nil {
		return
	}
	req = types.MustUnmarshalDepositRequest(k.cdc, bz)
	return req, true
}

// GetDepositRequestsByDepositor returns deposit requests by the depositor.
func (k Keeper) GetDepositRequestsByDepositor(ctx sdk.Context, depositor sdk.AccAddress) (reqs []types.DepositRequest) {
	_ = k.IterateDepositRequestsByDepositor(ctx, depositor, func(req types.DepositRequest) (stop bool, err error) {
		reqs = append(reqs, req)
		return false, nil
	})
	return
}

// SetDepositRequest stores deposit request for the batch execution.
func (k Keeper) SetDepositRequest(ctx sdk.Context, req types.DepositRequest) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalDepositRequest(k.cdc, req)
	store.Set(types.GetDepositRequestKey(req.PoolId, req.Id), bz)
}

// SetDepositRequestIndex stores the deposit request index.
func (k Keeper) SetDepositRequestIndex(ctx sdk.Context, req types.DepositRequest) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetDepositRequestIndexKey(req.GetDepositor(), req.PoolId, req.Id), []byte{})
}

// DeleteDepositRequest deletes deposit request and its index.
func (k Keeper) DeleteDepositRequest(ctx sdk.Context, req types.DepositRequest) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetDepositRequestKey(req.PoolId, req.Id))
	k.DeleteDepositRequestIndex(ctx, req)
}

// DeleteDepositRequestIndex deletes deposit request index.
func (k Keeper) DeleteDepositRequestIndex(ctx sdk.Context, req types.DepositRequest) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetDepositRequestIndexKey(req.GetDepositor(), req.PoolId, req.Id))
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

// IterateDepositRequestsByDepositor iterates through deposit requests in the
// store by a depositor and call cb on each order.
func (k Keeper) IterateDepositRequestsByDepositor(ctx sdk.Context, depositor sdk.AccAddress, cb func(req types.DepositRequest) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetDepositRequestIndexKeyPrefix(depositor))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		_, poolId, reqId := types.ParseDepositRequestIndexKey(iter.Key())
		req, _ := k.GetDepositRequest(ctx, poolId, reqId)
		stop, err := cb(req)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}
	return nil
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
