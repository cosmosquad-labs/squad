package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// InitGenesis initializes the capability module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}

	// Initialize objects to prevent from having nil slice
	if genState.Params.LiquidFarms == nil || len(genState.Params.LiquidFarms) == 0 {
		genState.Params.LiquidFarms = []types.LiquidFarm{}
	}
	if genState.RewardsAuctions == nil || len(genState.RewardsAuctions) == 0 {
		genState.RewardsAuctions = []types.RewardsAuction{}
	}

	k.SetParams(ctx, genState.Params)

	for _, record := range genState.QueuedFarmingRecords {
		farmerAddr, err := sdk.AccAddressFromBech32(record.Farmer)
		if err != nil {
			panic(err)
		}
		k.SetQueuedFarming(
			ctx,
			record.EndTime,
			record.FarmingCoinDenom,
			farmerAddr,
			record.QueuedFarming,
		)
	}
	for _, auction := range genState.RewardsAuctions {
		k.SetRewardsAuction(ctx, auction) // TODO: add test case to see if GetLastRewardsAuctionId returns the correct id
	}
	for _, bid := range genState.Bids {
		k.SetBid(ctx, bid)
	}
	for _, record := range genState.WinningBidRecords {
		k.SetWinningBid(ctx, record.WinningBid, record.AuctionId)
	}
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)

	// Initialize objects to prevent from having nil slice
	rewardsAuctions := k.GetAllRewardsAuctions(ctx)
	if len(rewardsAuctions) == 0 {
		rewardsAuctions = []types.RewardsAuction{}
	}
	if params.LiquidFarms == nil || len(params.LiquidFarms) == 0 {
		params.LiquidFarms = []types.LiquidFarm{}
	}

	queuedFarmingRecords := []types.QueuedFarmingRecord{}
	k.IterateQueuedFarmings(ctx, func(endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress, queuedFarming types.QueuedFarming) (stop bool) {
		queuedFarmingRecords = append(queuedFarmingRecords, types.QueuedFarmingRecord{
			EndTime:          endTime,
			FarmingCoinDenom: farmingCoinDenom,
			Farmer:           farmerAcc.String(),
			QueuedFarming:    queuedFarming,
		})
		return false
	})

	poolIds := []uint64{}
	for _, liquidFarm := range params.LiquidFarms {
		poolIds = append(poolIds, liquidFarm.PoolId)
	}

	bids := []types.Bid{}
	winningBidRecords := []types.WinningBidRecord{}
	for _, poolId := range poolIds {
		auctionId := k.GetLastRewardsAuctionId(ctx, poolId)
		winningBid, found := k.GetWinningBid(ctx, poolId, auctionId)
		if found {
			winningBidRecords = append(winningBidRecords, types.WinningBidRecord{
				AuctionId:  auctionId,
				WinningBid: winningBid,
			})
		}
		bids = append(bids, k.GetBidsByPoolId(ctx, poolId)...)
	}

	return &types.GenesisState{
		Params:               params,
		QueuedFarmingRecords: queuedFarmingRecords,
		RewardsAuctions:      rewardsAuctions,
		Bids:                 bids,
		WinningBidRecords:    winningBidRecords,
	}
}
