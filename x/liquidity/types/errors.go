package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/liquidity module sentinel errors
var (
	ErrInvalidDenom = sdkerrors.Register(ModuleName, 4, "invalid denom")
)
