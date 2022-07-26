package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgLend)(nil)
	_ sdk.Msg = (*MsgRedeem)(nil)
	_ sdk.Msg = (*MsgWithdraw)(nil)
)

// Message types for the lending module
const (
	TypeMsgLend     = "lend"
	TypeMsgRedeem   = "redeem"
	TypeMsgWithdraw = "withdraw"
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

// NewMsgRedeem returns a new MsgRedeem.
func NewMsgRedeem(lender sdk.AccAddress, coin sdk.Coin) *MsgRedeem {
	return &MsgRedeem{
		Lender: lender.String(),
		Coin:   coin,
	}
}

func (msg MsgRedeem) Route() string { return RouterKey }

func (msg MsgRedeem) Type() string { return TypeMsgRedeem }

func (msg MsgRedeem) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Lender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid lender address: %v", err)
	}
	if err := msg.Coin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid coin: %v", err)
	}
	return nil
}

func (msg MsgRedeem) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgRedeem) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Lender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// GetLender is a convenient helper for getting the lender address as
// sdk.AccAddress.
func (msg MsgRedeem) GetLender() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Lender)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgWithdraw returns a new MsgWithdraw.
func NewMsgWithdraw(lender sdk.AccAddress, coin sdk.Coin) *MsgWithdraw {
	return &MsgWithdraw{
		Lender: lender.String(),
		Coin:   coin,
	}
}

func (msg MsgWithdraw) Route() string { return RouterKey }

func (msg MsgWithdraw) Type() string { return TypeMsgWithdraw }

func (msg MsgWithdraw) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Lender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid lender address: %v", err)
	}
	if err := msg.Coin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid coin: %v", err)
	}
	return nil
}

func (msg MsgWithdraw) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdraw) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Lender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// GetLender is a convenient helper for getting the lender address as
// sdk.AccAddress.
func (msg MsgWithdraw) GetLender() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Lender)
	if err != nil {
		panic(err)
	}
	return addr
}
