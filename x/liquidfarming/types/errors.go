package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/liquidfarming module sentinel errors
var (
	ErrEmptyLiquidFarms = sdkerrors.Register(ModuleName, 4, "submitted parameter liquid farms are empty")
)
