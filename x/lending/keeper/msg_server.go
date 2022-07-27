package keeper

import (
	"context"

	"github.com/cosmosquad-labs/squad/v2/x/lending/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// Lend defines a method for lending assets.
func (m msgServer) Lend(goCtx context.Context, msg *types.MsgLend) (*types.MsgLendResponse, error) {
	//ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: not implemented

	return &types.MsgLendResponse{}, nil
}

// Withdraw defines a method for withdrawing lent assets.
func (m msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	// TODO: not implemented

	return &types.MsgWithdrawResponse{}, nil
}
