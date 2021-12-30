package keeper

import (
	"context"

	"github.com/tendermint/farming/x/liquidity/types"
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

// DepositWithinBatch defines a method to deposit coins to the pool.
// The first deposit of the pool creates a pool and mints an initial pool coin.
func (k msgServer) DepositWithinBatch(goCtx context.Context, msg *types.MsgDepositWithinBatch) (*types.MsgDepositWithinBatchResponse, error) {
	// TODO: not implemented yet
	return &types.MsgDepositWithinBatchResponse{}, nil
}

// WithdrawWithinBatch defines a method ...
func (k msgServer) WithdrawWithinBatch(goCtx context.Context, msg *types.MsgWithdrawWithinBatch) (*types.MsgWithdrawWithinBatchResponse, error) {
	// TODO: not implemented yet
	return &types.MsgWithdrawWithinBatchResponse{}, nil
}

// SwapWithinBatch defines a method ...
func (k msgServer) SwapWithinBatch(goCtx context.Context, msg *types.MsgSwapWithinBatch) (*types.MsgSwapWithinBatchResponse, error) {
	// TODO: not implemented yet
	return &types.MsgSwapWithinBatchResponse{}, nil
}
