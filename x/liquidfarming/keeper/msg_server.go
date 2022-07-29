package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
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

// Farm defines a method for depositing pool coin and the module mints LFCoin at the end of epoch.
func (m msgServer) Farm(goCtx context.Context, msg *types.MsgFarm) (*types.MsgFarmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.Farm(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgFarmResponse{}, nil
}

// CancelQueuedFarming defines a method for canceling the queued farming.
func (m msgServer) CancelQueuedFarming(goCtx context.Context, msg *types.MsgCancelQueuedFarming) (*types.MsgCancelQueuedFarmingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.CancelQueuedFarming(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgCancelQueuedFarmingResponse{}, nil
}

// Unfarm defines a method for unfarming LFCoin.
func (m msgServer) Unfarm(goCtx context.Context, msg *types.MsgUnfarm) (*types.MsgUnfarmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, err := m.Keeper.Unfarm(ctx, msg.PoolId, msg.GetFarmer(), msg.LFCoin); err != nil {
		return nil, err
	}

	return &types.MsgUnfarmResponse{}, nil
}

// UnfarmAndWithdraw defines a method for unfarming LFCoin and withdraw pool coin from the pool.
func (m msgServer) UnfarmAndWithdraw(goCtx context.Context, msg *types.MsgUnfarmAndWithdraw) (*types.MsgUnfarmAndWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.UnfarmAndWithdraw(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgUnfarmAndWithdrawResponse{}, nil
}

// PlaceBid defines a method for placing a bid for a rewards auction.
func (m msgServer) PlaceBid(goCtx context.Context, msg *types.MsgPlaceBid) (*types.MsgPlaceBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, err := m.Keeper.PlaceBid(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgPlaceBidResponse{}, nil
}

// RefundBid defines a method for refunding the bid that is not winning for the auction.
func (m msgServer) RefundBid(goCtx context.Context, msg *types.MsgRefundBid) (*types.MsgRefundBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.RefundBid(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgRefundBidResponse{}, nil
}
