package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/liquidfarming module sentinel errors
var (
	ErrLiquidFarmNotFound        = sdkerrors.Register(ModuleName, 1, "liquid farm not found")
	ErrPoolNotFound              = sdkerrors.Register(ModuleName, 2, "pool not found")
	ErrDepositRequestNotFound    = sdkerrors.Register(ModuleName, 3, "deposit request not found")
	ErrInsufficientDepositAmount = sdkerrors.Register(ModuleName, 4, "insufficient deposit amount")
)
