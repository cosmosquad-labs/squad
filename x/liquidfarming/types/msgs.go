package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgDeposit)(nil)
	_ sdk.Msg = (*MsgCancel)(nil)
	_ sdk.Msg = (*MsgWithdraw)(nil)
	_ sdk.Msg = (*MsgPlaceBid)(nil)
)

// Message types for the farming module
const (
	TypeMsgDeposit  = "deposit"
	TypeMsgCancel   = "cancel"
	TypeMsgWithdraw = "withdraw"
	TypeMsgPlaceBid = "place_bid"
)

// NewMsgDeposit returns a new MsgDeposit.
func NewMsgDeposit(poolId uint64, depositor string, depositCoin sdk.Coin) *MsgDeposit {
	return &MsgDeposit{
		PoolId:      poolId,
		Depositor:   depositor,
		DepositCoin: depositCoin,
	}
}

func (msg MsgDeposit) Route() string { return RouterKey }

func (msg MsgDeposit) Type() string { return TypeMsgDeposit }

func (msg MsgDeposit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Depositor); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid depositor address: %v", err)
	}
	if !msg.DepositCoin.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "deposit coin must be positive")
	}
	if err := msg.DepositCoin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid deposit coin: %v", err)
	}
	return nil
}

func (msg MsgDeposit) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDeposit) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgDeposit) GetDepositor() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgCancel returns a new MsgCancel.
func NewMsgCancel(depositId uint64) *MsgCancel {
	return &MsgCancel{
		DepositId: depositId,
	}
}

func (msg MsgCancel) Route() string { return RouterKey }

func (msg MsgCancel) Type() string { return TypeMsgCancel }

func (msg MsgCancel) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Depositor); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid depositor address: %v", err)
	}
	if msg.DepositId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid deposit id")
	}
	return nil
}

func (msg MsgCancel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCancel) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgCancel) GetDepositor() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgWithdraw returns a new MsgWithdraw.
func NewMsgWithdraw(depositId uint64, withdrawer string, lfCoin sdk.Coin) *MsgWithdraw {
	return &MsgWithdraw{
		DepositId:  depositId,
		Withdrawer: withdrawer,
		LFCoin:     lfCoin,
	}
}

func (msg MsgWithdraw) Route() string { return RouterKey }

func (msg MsgWithdraw) Type() string { return TypeMsgWithdraw }

func (msg MsgWithdraw) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Withdrawer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid withdrawer address: %v", err)
	}
	if msg.DepositId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid deposit id")
	}
	if !msg.LFCoin.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "withdrawing coin must be positive")
	}
	if err := msg.LFCoin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid LFCoin: %v", err)
	}
	return nil
}

func (msg MsgWithdraw) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdraw) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Withdrawer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgWithdraw) GetWithdrawer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Withdrawer)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgPlaceBid returns a new MsgPlaceBid.
func NewMsgPlaceBid(auctionId uint64, bidder string, amount sdk.Coin) *MsgPlaceBid {
	return &MsgPlaceBid{
		AuctionId: auctionId,
		Bidder:    bidder,
		Amount:    amount,
	}
}

func (msg MsgPlaceBid) Route() string { return RouterKey }

func (msg MsgPlaceBid) Type() string { return TypeMsgPlaceBid }

func (msg MsgPlaceBid) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Bidder); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bidder address: %v", err)
	}
	if msg.AuctionId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid auction id")
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be positive")
	}
	if err := msg.Amount.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount: %v", err)
	}
	return nil
}

func (msg MsgPlaceBid) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgPlaceBid) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgPlaceBid) GetBidder() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		panic(err)
	}
	return addr
}
