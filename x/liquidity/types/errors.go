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
)
