package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgCreatePool)(nil)
	_ sdk.Msg = (*MsgDepositBatch)(nil)
	_ sdk.Msg = (*MsgWithdrawBatch)(nil)
	_ sdk.Msg = (*MsgSwapBatch)(nil)
	_ sdk.Msg = (*MsgCancelSwapBatch)(nil)
)

// Message types for the liquidity module
const (
	TypeMsgCreatePool      = "create_pool"
	TypeMsgDepositBatch    = "deposit_batch"
	TypeMsgWithdrawBatch   = "withdraw_batch"
	TypeMsgSwapBatch       = "swap_batch"
	TypeMsgCancelSwapBatch = "cancel_swap_batch"
)

// NewMsgCreatePool creates a new MsgCreatePool.
func NewMsgCreatePool(
	creator sdk.AccAddress,
	depositCoins sdk.Coins,
) *MsgCreatePool {
	return &MsgCreatePool{
		Creator:      creator.String(),
		DepositCoins: depositCoins,
	}
}

func (msg MsgCreatePool) Route() string { return RouterKey }

func (msg MsgCreatePool) Type() string { return TypeMsgCreatePool }

func (msg MsgCreatePool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %v", err)
	}
	if err := msg.DepositCoins.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid deposit coins: %v", err)
	}
	if len(msg.DepositCoins) != ReserveCoinNum {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid number of deposit coin")
	}
	return nil
}

func (msg MsgCreatePool) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreatePool) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgCreatePool) GetCreator() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgDepositBatch creates a new MsgDepositBatch.
func NewMsgDepositBatch(
	depositor sdk.AccAddress,
	poolID uint64,
	depositCoins sdk.Coins,
) *MsgDepositBatch {
	return &MsgDepositBatch{
		Depositor: depositor.String(),
		PoolId:    poolID,
		Coins:     depositCoins,
	}
}

func (msg MsgDepositBatch) Route() string { return RouterKey }

func (msg MsgDepositBatch) Type() string { return TypeMsgDepositBatch }

func (msg MsgDepositBatch) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Depositor); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid depositor address: %v", err)
	}
	if err := msg.Coins.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid deposit coins: %v", err)
	}
	if len(msg.Coins) != ReserveCoinNum {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid number of deposit coin")
	}
	return nil
}

func (msg MsgDepositBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDepositBatch) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgDepositBatch) GetDepositor() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgWithdrawBatch creates a new MsgWithdrawBatch.
func NewMsgWithdrawBatch(
	withdrawer sdk.AccAddress,
	poolID uint64,
	poolCoin sdk.Coin,
) *MsgWithdrawBatch {
	return &MsgWithdrawBatch{
		Withdrawer: withdrawer.String(),
		PoolId:     poolID,
		PoolCoin:   poolCoin,
	}
}

func (msg MsgWithdrawBatch) Route() string { return RouterKey }

func (msg MsgWithdrawBatch) Type() string { return TypeMsgWithdrawBatch }

func (msg MsgWithdrawBatch) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Withdrawer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid withdrawer address: %v", err)
	}
	if err := msg.PoolCoin.Validate(); err != nil {
		return err
	}
	if !msg.PoolCoin.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "pool coin must be positive")
	}
	return nil
}

func (msg MsgWithdrawBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdrawBatch) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Withdrawer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgWithdrawBatch) GetWithdrawer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Withdrawer)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgSwapBatch creates a new MsgSwapBatch.
func NewMsgSwapBatch(
	orderer sdk.AccAddress,
	coin sdk.Coin,
	demandCoinDenom string,
	price sdk.Dec,
	orderLifeSpan time.Duration,
) *MsgSwapBatch {
	return &MsgSwapBatch{
		Orderer:         orderer.String(),
		Coin:            coin,
		DemandCoinDenom: demandCoinDenom,
		Price:           price,
		OrderLifespan:   orderLifeSpan,
	}
}

func (msg MsgSwapBatch) Route() string { return RouterKey }

func (msg MsgSwapBatch) Type() string { return TypeMsgSwapBatch }

func (msg MsgSwapBatch) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Orderer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orderer address: %v", err)
	}
	if err := msg.Coin.Validate(); err != nil {
		return err
	}
	if !msg.Coin.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "offer coin must be positive")
	}
	if !msg.Price.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "price must be positive")
	}
	if !msg.Coin.Amount.GTE(MinOfferCoinAmount) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "offer coin is less than minimum offer coin amount")
	}
	// TODO: anything to validate for OrderLifeSpan?
	return nil
}

func (msg MsgSwapBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgSwapBatch) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Orderer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgSwapBatch) GetSwapOrderer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Orderer)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgCancelSwapBatch creates a new MsgCancelSwapBatch.
func NewMsgCancelSwapBatch(
	orderer sdk.AccAddress,
	swapRequestId uint64,
) *MsgCancelSwapBatch {
	return &MsgCancelSwapBatch{
		SwapRequestId: swapRequestId,
		Orderer:       orderer.String(),
	}
}

func (msg MsgCancelSwapBatch) Route() string { return RouterKey }

func (msg MsgCancelSwapBatch) Type() string { return TypeMsgCancelSwapBatch }

func (msg MsgCancelSwapBatch) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Orderer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orderer address: %v", err)
	}
	return nil
}

func (msg MsgCancelSwapBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCancelSwapBatch) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Orderer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
