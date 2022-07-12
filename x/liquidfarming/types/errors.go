package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/liquidfarming module sentinel errors
var (
	ErrInsufficientFarmingCoinAmount = sdkerrors.Register(ModuleName, 2, "insufficient farming coin amount")
	ErrInsufficientUnfarmingAmount   = sdkerrors.Register(ModuleName, 3, "insufficient unfarming amount")
)
