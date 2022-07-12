package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/liquidfarming module sentinel errors
var (
	ErrLiquidFarmNotFound            = sdkerrors.Register(ModuleName, 1, "liquid farm not found")
	ErrPoolNotFound                  = sdkerrors.Register(ModuleName, 2, "pool not found")
	ErrQueuedFarmingNotFound         = sdkerrors.Register(ModuleName, 3, "queued farming not found")
	ErrInsufficientFarmingCoinAmount = sdkerrors.Register(ModuleName, 4, "insufficient farming coin amount")
	ErrInsufficientWithdrawingAmount = sdkerrors.Register(ModuleName, 5, "insufficient withdrawing amount")
)
