package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/fundraising module sentinel errors
var (
	ErrUnknownAuctionType     = sdkerrors.Register(ModuleName, 2, "unknown auction type")
	ErrNoAuctionHandlerExists = sdkerrors.Register(ModuleName, 3, "no handler exists for auction type")
	ErrInvalidAuctionCustom   = sdkerrors.Register(ModuleName, 4, "invalid auction content")
	ErrInvalidAuctionType     = sdkerrors.Register(ModuleName, 5, "invalid auction type")
)
