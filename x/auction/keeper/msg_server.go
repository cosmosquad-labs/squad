package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/x/auction/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// CreateFixedPriceAuction defines a method to create fixed price auction.
func (m msgServer) CreateFixedPriceAuction(goCtx context.Context, msg *types.MsgCreateFixedPriceAuction) (*types.MsgCreateFixedPriceAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.CreateFixedPriceAuction(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgCreateFixedPriceAuctionResponse{}, nil
}
