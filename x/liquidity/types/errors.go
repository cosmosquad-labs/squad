package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DONTCOVER

// x/liquidity module sentinel errors
var (
	ErrInsufficientDepositAmount = sdkerrors.Register(ModuleName, 2, "insufficient deposit amount")
	ErrPoolAlreadyExists         = sdkerrors.Register(ModuleName, 3, "pool already exists")
	ErrPoolNotFound              = sdkerrors.Register(ModuleName, 4, "pool not found")
	ErrWrongPair                 = sdkerrors.Register(ModuleName, 5, "wrong coin denom pair")
	ErrWrongPoolCoinDenom        = sdkerrors.Register(ModuleName, 6, "wrong pool coin denom")
	ErrInvalidPriceTick          = sdkerrors.Register(ModuleName, 7, "price not fit into ticks")
	ErrPriceOutOfRange           = sdkerrors.Register(ModuleName, 8, "price out of range limit")
	ErrSwapRequestNotFound       = sdkerrors.Register(ModuleName, 9, "swap request not found") // TODO: use sdkerrors.NotFound?
)
