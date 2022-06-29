package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/fundraising module sentinel errors
var (
	ErrUnknownAuctionType = sdkerrors.Register(ModuleName, 2, "unknown auction type")
)
