package keeper

import (
	"time"

	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// GetBidId returns the last bid id for the bid.
func (k Keeper) GetBidId(ctx sdk.Context, auctionId uint64) uint64 {
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

// GetQueuedFarming returns a queued farming for given farming coin denom
// and farmer.
func (k Keeper) GetQueuedFarming(ctx sdk.Context, endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress) (queuedFarming types.QueuedFarming, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetQueuedFarmingKey(endTime, farmingCoinDenom, farmerAcc))
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshal(bz, &queuedFarming)
	found = true
	return
}

// SetQueuedFarming sets a queued farming for given farming coin denom
// and farmer.
func (k Keeper) SetQueuedFarming(ctx sdk.Context, endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress, queuedFarming types.QueuedFarming) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&queuedFarming)
	store.Set(types.GetQueuedFarmingKey(endTime, farmingCoinDenom, farmerAcc), bz)
	store.Set(types.GetQueuedFarmingIndexKey(farmerAcc, farmingCoinDenom, endTime), []byte{})
}

// DeleteQueuedFarming deletes a queued farming for given farming coin denom
// and farmer.
func (k Keeper) DeleteQueuedFarming(ctx sdk.Context, endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetQueuedFarmingKey(endTime, farmingCoinDenom, farmerAcc))
	store.Delete(types.GetQueuedFarmingIndexKey(farmerAcc, farmingCoinDenom, endTime))
}

// GetQueuedFarmingsByFarmer returns all queued farmings
// by a farmer.
func (k Keeper) GetQueuedFarmingsByFarmer(ctx sdk.Context, farmerAcc sdk.AccAddress) []types.QueuedFarming {
	queuedFarmings := []types.QueuedFarming{}
	k.IterateQueuedFarmingsByFarmer(ctx, farmerAcc, func(farmingCoinDenom string, endTime time.Time, queuedFarming types.QueuedFarming) (stop bool) {
		queuedFarmings = append(queuedFarmings, queuedFarming)
		return false
	})
	return queuedFarmings
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

// IterateQueuedFarmingsByFarmer iterates through all queued farmings
// by farmer stored in the store and invokes callback function for each item.
// Stops the iteration when the callback function returns true.
func (k Keeper) IterateQueuedFarmingsByFarmer(ctx sdk.Context, farmerAcc sdk.AccAddress, cb func(stakingCoinDenom string, endTime time.Time, queuedFarming types.QueuedFarming) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetQueuedFarmingsByFarmerPrefix(farmerAcc))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		_, farmingCoinDenom, endTime := types.ParseQueuedFarmingIndexKey(iter.Key())
		queuedFarming, _ := k.GetQueuedFarming(ctx, endTime, farmingCoinDenom, farmerAcc)
		if cb(farmingCoinDenom, endTime, queuedFarming) {
			break
		}
	}
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
