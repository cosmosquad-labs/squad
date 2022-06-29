package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/x/auction/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper.
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

// Auctions queries all auctions.
func (k Querier) Auctions(c context.Context, req *types.QueryAuctionsRequest) (*types.QueryAuctionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// TODO: not implemented yet

	return &types.QueryAuctionsResponse{}, nil
}

// Auction queries the specific auction.
func (k Querier) Auction(c context.Context, req *types.QueryAuctionRequest) (*types.QueryAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// TODO: not implemented yet

	return &types.QueryAuctionResponse{}, nil
}

// Bids queries all bids for the auction.
func (k Querier) Bids(c context.Context, req *types.QueryBidsRequest) (*types.QueryBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// TODO: not implemented yet

	return &types.QueryBidsResponse{}, nil
}

// Bid queries the specific bid from the auction id and bid id.
func (k Querier) Bid(c context.Context, req *types.QueryBidRequest) (*types.QueryBidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// TODO: not implemented yet

	return &types.QueryBidResponse{}, nil
}

// Vestings queries all vesting queues for the auction.
func (k Querier) Vestings(c context.Context, req *types.QueryVestingsRequest) (*types.QueryVestingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// TODO: not implemented yet

	return &types.QueryVestingsResponse{}, nil
}
