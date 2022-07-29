package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/liquidfarming module sentinel errors
var (
	ErrSmallerThanMinimumAmount    = sdkerrors.Register(ModuleName, 2, "smaller than minimum amount")
	ErrInsufficientUnfarmingAmount = sdkerrors.Register(ModuleName, 3, "insufficient unfarming amount")
	ErrInvalidAuctionStatus        = sdkerrors.Register(ModuleName, 4, "invalid auction status")
)
