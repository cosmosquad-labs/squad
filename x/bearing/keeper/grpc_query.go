package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/farming/x/bearing/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper.
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Params queries the parameters of the bearing module.
func (k Querier) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.paramSpace.GetParamSet(ctx, &params)
	return &types.QueryParamsResponse{Params: params}, nil
}

// Bearings queries all bearings.
func (k Querier) Bearings(c context.Context, req *types.QueryBearingsRequest) (*types.QueryBearingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.SourceAddress != "" {
		if _, err := sdk.AccAddressFromBech32(req.SourceAddress); err != nil {
			return nil, err
		}
	}

	if req.DestinationAddress != "" {
		if _, err := sdk.AccAddressFromBech32(req.DestinationAddress); err != nil {
			return nil, err
		}
	}

	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.paramSpace.GetParamSet(ctx, &params)

	var bearings []types.BearingResponse
	for _, b := range params.Bearings {
		if req.Name != "" && b.Name != req.Name ||
			req.SourceAddress != "" && b.SourceAddress != req.SourceAddress ||
			req.DestinationAddress != "" && b.DestinationAddress != req.DestinationAddress {
			continue
		}

		collectedCoins := k.GetTotalCollectedCoins(ctx, b.Name)
		bearings = append(bearings, types.BearingResponse{
			Bearing:             b,
			TotalCollectedCoins: collectedCoins,
		})
	}

	return &types.QueryBearingsResponse{Bearings: bearings}, nil
}

// Addresses queries an address that can be used as source and destination is derived according to the given name, module name and address type.
func (k Querier) Addresses(_ context.Context, req *types.QueryAddressesRequest) (*types.QueryAddressesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.Name == "" && req.ModuleName == "" {
		return nil, status.Error(codes.InvalidArgument, "at least one input of name or module name is required")
	}

	if req.ModuleName == "" && req.Type == types.AddressType32Bytes {
		req.ModuleName = types.ModuleName
	}

	addr := types.DeriveAddress(req.Type, req.ModuleName, req.Name)
	if addr.Empty() {
		return nil, status.Error(codes.InvalidArgument, "invalid names with address type")
	}

	return &types.QueryAddressesResponse{Address: addr.String()}, nil
}
