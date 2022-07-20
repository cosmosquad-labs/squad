package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v2/x/liquidity/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper.
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Params queries the parameters of the module.
func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}

// LiquidFarms queries all liquidfarms.
func (k Querier) LiquidFarms(c context.Context, req *types.QueryLiquidFarmsRequest) (*types.QueryLiquidFarmsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	liquidFarmsRes := []types.LiquidFarmResponse{}
	for _, lf := range k.GetParams(ctx).LiquidFarms {
		reserveAcc := types.LiquidFarmReserveAddress(lf.PoolId)
		poolCoinDenom := liquiditytypes.PoolCoinDenom(lf.PoolId)
		queuedAmt := k.farmingKeeper.GetAllQueuedCoinsByFarmer(ctx, reserveAcc).AmountOf(poolCoinDenom)
		stakedAmt := k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAcc).AmountOf(poolCoinDenom)
		liquidFarmsRes = append(liquidFarmsRes, types.LiquidFarmResponse{
			PoolId:                   lf.PoolId,
			LiquidFarmReserveAddress: reserveAcc.String(),
			LFCoinDenom:              types.LiquidFarmCoinDenom(lf.PoolId),
			MinimumFarmAmount:        lf.MinimumFarmAmount,
			MinimumBidAmount:         lf.MinimumBidAmount,
			QueuedCoin:               sdk.NewCoin(poolCoinDenom, queuedAmt),
			StakedCoin:               sdk.NewCoin(poolCoinDenom, stakedAmt),
		})
	}

	return &types.QueryLiquidFarmsResponse{LiquidFarms: liquidFarmsRes}, nil
}

// LiquidFarm queries the specific liquidfarm.
func (k Querier) LiquidFarm(c context.Context, req *types.QueryLiquidFarmRequest) (*types.QueryLiquidFarmResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)

	liquidFarmRes := types.LiquidFarmResponse{}
	for _, lf := range k.GetParams(ctx).LiquidFarms {
		if lf.PoolId == req.PoolId {
			reserveAcc := types.LiquidFarmReserveAddress(lf.PoolId)
			poolCoinDenom := liquiditytypes.PoolCoinDenom(lf.PoolId)
			queuedAmt := k.farmingKeeper.GetAllQueuedCoinsByFarmer(ctx, reserveAcc).AmountOf(poolCoinDenom)
			stakedAmt := k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAcc).AmountOf(poolCoinDenom)
			liquidFarmRes = types.LiquidFarmResponse{
				PoolId:                   lf.PoolId,
				LiquidFarmReserveAddress: reserveAcc.String(),
				LFCoinDenom:              types.LiquidFarmCoinDenom(lf.PoolId),
				MinimumFarmAmount:        lf.MinimumFarmAmount,
				MinimumBidAmount:         lf.MinimumBidAmount,
				QueuedCoin:               sdk.NewCoin(poolCoinDenom, queuedAmt),
				StakedCoin:               sdk.NewCoin(poolCoinDenom, stakedAmt),
			}
		}
	}

	return &types.QueryLiquidFarmResponse{LiquidFarm: liquidFarmRes}, nil
}

// QueuedFarmings queries all queued farmings.
func (k Querier) QueuedFarmings(c context.Context, req *types.QueryQueuedFarmingsRequest) (*types.QueryQueuedFarmingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	qfsStore := prefix.NewStore(store, types.QueuedFarmingKeyPrefix)

	var qfs []types.QueuedFarming
	pageRes, err := query.FilteredPaginate(qfsStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		qf, err := types.UnmarshalQueuedFarming(k.cdc, value)
		if err != nil {
			return false, err
		}

		if qf.PoolId != req.PoolId {
			return false, nil
		}

		if accumulate {
			qfs = append(qfs, qf)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryQueuedFarmingsResponse{QueuedFarmings: qfs, Pagination: pageRes}, nil
}

// QueuedFarmingsByFarmer queries all queued farmings by the given farmer.
func (k Querier) QueuedFarmingsByFarmer(c context.Context, req *types.QueryQueuedFarmingsByFarmerRequest) (*types.QueryQueuedFarmingsByFarmerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	if req.FarmerAddress != "" {
		if _, err := sdk.AccAddressFromBech32(req.FarmerAddress); err != nil {
			return nil, err
		}
	}

	// TODO: not implemented yet
	// Consider combining this query with QueuedFarmings
	// ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryQueuedFarmingsByFarmerResponse{}, nil
}

// RewardsAuction queries rewards auction
func (k Querier) RewardsAuction(c context.Context, req *types.QueryRewardsAuctionRequest) (*types.QueryRewardsAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// ctx := sdk.UnwrapSDKContext(c)
	// TODO: not implemented yet

	return &types.QueryRewardsAuctionResponse{}, nil
}

// Bids queries all bids.
func (k Querier) Bids(c context.Context, req *types.QueryBidsRequest) (*types.QueryBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)
	auctionId := k.GetLastRewardsAuctionId(ctx, req.PoolId)
	_, found := k.GetRewardsAuction(ctx, req.PoolId, auctionId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "auction that corresponds to pool %d not found", req.PoolId)
	}

	var bids []types.Bid
	// store := ctx.KVStore(k.storeKey)

	var pageRes *query.PageResponse
	// var err error

	return &types.QueryBidsResponse{Bids: bids, Pagination: pageRes}, nil
}

// BidByBidder queries the specific bid.
func (k Querier) BidByBidder(c context.Context, req *types.QueryBidByBidderRequest) (*types.QueryBidByBidderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// ctx := sdk.UnwrapSDKContext(c)
	// TODO: not implemented yet

	return &types.QueryBidByBidderResponse{}, nil
}
