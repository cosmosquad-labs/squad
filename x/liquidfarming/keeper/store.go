package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// GetLastQueuedFarmingId returns the last queued farming id for the pool id.
func (k Keeper) GetLastQueuedFarmingId(ctx sdk.Context, poolId uint64) uint64 {
	var id uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLastQueuedFarmingIdKey(poolId))
	if bz == nil {
		id = 0 // initialize the deposit request id
	} else {
		val := gogotypes.UInt64Value{}
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return id
}

// SetQueuedFarmingId sets the deposit request id with the given pool id.
func (k Keeper) SetQueuedFarmingId(ctx sdk.Context, poolId uint64, reqId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: reqId})
	store.Set(types.GetLastQueuedFarmingIdKey(poolId), bz)
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

// GetQueuedFarming returns the particular queued farming.
func (k Keeper) GetQueuedFarming(ctx sdk.Context, poolId, reqId uint64) (qf types.QueuedFarming, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetQueuedFarmingKey(poolId, reqId))
	if bz == nil {
		return
	}
	qf = types.MustUnmarshalQueuedFarming(k.cdc, bz)
	return qf, true
}

// GetQueuedFarmingsByDepositor returns queued farmings by the depositor.
func (k Keeper) GetQueuedFarmingsByDepositor(ctx sdk.Context, depositor sdk.AccAddress) (qfs []types.QueuedFarming) {
	_ = k.IterateQueuedFarmingsByDepositor(ctx, depositor, func(req types.QueuedFarming) (stop bool, err error) {
		qfs = append(qfs, req)
		return false, nil
	})
	return
}

// SetQueuedFarming stores queued farming for the batch execution.
func (k Keeper) SetQueuedFarming(ctx sdk.Context, qf types.QueuedFarming) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalQueuedFarming(k.cdc, qf)
	store.Set(types.GetQueuedFarmingKey(qf.PoolId, qf.Id), bz)
}

// SetQueuedFarmingIndex stores the queued farming index.
func (k Keeper) SetQueuedFarmingIndex(ctx sdk.Context, req types.QueuedFarming) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetQueuedFarmingIndexKey(req.GetFarmer(), req.PoolId, req.Id), []byte{})
}

// DeleteQueuedFarming deletes deposit request and its index.
func (k Keeper) DeleteQueuedFarming(ctx sdk.Context, req types.QueuedFarming) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetQueuedFarmingKey(req.PoolId, req.Id))
	k.DeleteQueuedFarmingIndex(ctx, req)
}

// DeleteQueuedFarmingIndex deletes deposit request index.
func (k Keeper) DeleteQueuedFarmingIndex(ctx sdk.Context, req types.QueuedFarming) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetQueuedFarmingIndexKey(req.GetFarmer(), req.PoolId, req.Id))
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

// IterateQueuedFarmingsByDepositor iterates through deposit requests in the
// store by a depositor and call cb on each order.
func (k Keeper) IterateQueuedFarmingsByDepositor(ctx sdk.Context, depositor sdk.AccAddress, cb func(req types.QueuedFarming) (stop bool, err error)) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetQueuedFarmingIndexKeyPrefix(depositor))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		_, poolId, reqId := types.ParseQueuedFarmingIndexKey(iter.Key())
		req, _ := k.GetQueuedFarming(ctx, poolId, reqId)
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

type MsgRefundBid struct {
	AuctionId uint64
	BidId     string
	Bidder    sdk.Coin
}
