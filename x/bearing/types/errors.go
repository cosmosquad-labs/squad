package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Sentinel errors for the bearing module.
var (
	ErrInvalidBearingName      = sdkerrors.Register(ModuleName, 2, "bearing name only allows letters, digits, and dash(-) without spaces and the maximum length is 50")
	ErrInvalidStartEndTime     = sdkerrors.Register(ModuleName, 3, "bearing end time must be after the start time")
	ErrInvalidBearingRate      = sdkerrors.Register(ModuleName, 4, "invalid bearing rate")
	ErrInvalidTotalBearingRate = sdkerrors.Register(ModuleName, 5, "invalid total rate of the bearings with the same source address")
	ErrDuplicateBearingName    = sdkerrors.Register(ModuleName, 6, "duplicate bearing name")
)
