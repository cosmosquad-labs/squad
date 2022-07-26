package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgLend)(nil)
)

// Message types for the lending module
const (
	TypeMsgLend = "lend"
)

// NewMsgLend returns a new MsgLend.
func NewMsgLend(lender sdk.AccAddress, coin sdk.Coin) *MsgLend {
	return &MsgLend{
		Lender: lender.String(),
		Coin:   coin,
	}
}

func (msg MsgLend) Route() string { return RouterKey }

func (msg MsgLend) Type() string { return TypeMsgLend }

func (msg MsgLend) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Lender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid lender address: %v", err)
	}
	if err := msg.Coin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid coin: %v", err)
	}
	return nil
}

func (msg MsgLend) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgLend) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Lender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// GetLender is a convenient helper for getting the lender address as
// sdk.AccAddress.
func (msg MsgLend) GetLender() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Lender)
	if err != nil {
		panic(err)
	}
	return addr
}
