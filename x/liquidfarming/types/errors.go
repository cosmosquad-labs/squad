package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/liquidfarming module sentinel errors
var (
	ErrInsufficientFarmingAmount   = sdkerrors.Register(ModuleName, 2, "insufficient farming coin amount")
	ErrInsufficientUnfarmingAmount = sdkerrors.Register(ModuleName, 3, "insufficient unfarming amount")
	ErrInsufficientBidAmount       = sdkerrors.Register(ModuleName, 4, "insufficient bid amount")
	ErrInvalidAuctionStatus        = sdkerrors.Register(ModuleName, 5, "invalid auction status")
)
