package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/liquidfarming module sentinel errors
var (
	ErrEmptyLiquidFarmRequests = sdkerrors.Register(ModuleName, 4, "submitted parameter liquid farm requests are empty")
)
