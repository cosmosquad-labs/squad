package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper.
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Params queries the parameters of the liquidfarming module.
func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

// LiquidFarms queries all deposit requests.
func (k Keeper) LiquidFarms(c context.Context, req *types.QueryLiquidFarmsRequest) (*types.QueryLiquidFarmsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: not implemented yet

	return &types.QueryLiquidFarmsResponse{}, nil
}

// LiquidFarm queries the specific liquidfarm.
func (k Keeper) LiquidFarm(c context.Context, req *types.QueryLiquidFarmRequest) (*types.QueryLiquidFarmResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: not implemented yet

	return &types.QueryLiquidFarmResponse{}, nil
}

// QueuedFarmings queries all queued farmings.
func (k Keeper) QueuedFarmings(c context.Context, req *types.QueryQueuedFarmingsRequest) (*types.QueryQueuedFarmingsResponse, error) {
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
		dr, err := types.UnmarshalQueuedFarming(k.cdc, value)
		if err != nil {
			return false, err
		}

		if dr.PoolId != req.PoolId {
			return false, nil
		}

		if accumulate {
			qfs = append(qfs, dr)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryQueuedFarmingsResponse{QueuedFarmings: qfs, Pagination: pageRes}, nil
}

// QueuedFarming queries the specific queued farming.
func (k Keeper) QueuedFarming(c context.Context, req *types.QueryQueuedFarmingRequest) (*types.QueryQueuedFarmingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	if req.RequestId == 0 {
		return nil, status.Error(codes.InvalidArgument, "deposit request id cannot be 0")
	}

	// ctx := sdk.UnwrapSDKContext(c)

	// dq, found := k.GetQueuedFarming(ctx, req.PoolId, req.RequestId)
	// if !found {
	// 	return nil, status.Errorf(codes.NotFound, "deposit request of pool id %d and request id %d doesn't exist or deleted", req.PoolId, req.RequestId)
	// }

	return &types.QueryQueuedFarmingResponse{}, nil
}

// Bids queries all bids.
func (k Keeper) Bids(c context.Context, req *types.QueryBidsRequest) (*types.QueryBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// ctx := sdk.UnwrapSDKContext(c)
	// TODO: not implemented yet
	// sort by bid id

	return &types.QueryBidsResponse{}, nil
}

// Bid queries the specific bid.
func (k Keeper) Bid(c context.Context, req *types.QueryBidRequest) (*types.QueryBidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// ctx := sdk.UnwrapSDKContext(c)
	// TODO: not implemented yet

	return &types.QueryBidResponse{}, nil
}

// RewardsAuctions queries rewards auctions
func (k Keeper) RewardsAuctions(c context.Context, req *types.QueryRewardsAuctionsRequest) (*types.QueryRewardsAuctionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// ctx := sdk.UnwrapSDKContext(c)
	// TODO: not implemented yet

	return &types.QueryRewardsAuctionsResponse{}, nil
}

// RewardsAuction queries the specific a rewards auction.
func (k Keeper) RewardsAuction(c context.Context, req *types.QueryRewardsAuctionRequest) (*types.QueryRewardsAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// ctx := sdk.UnwrapSDKContext(c)
	// TODO: not implemented yet

	return &types.QueryRewardsAuctionResponse{}, nil
}
