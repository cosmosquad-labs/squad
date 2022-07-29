package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgFarm)(nil)
	_ sdk.Msg = (*MsgUnfarm)(nil)
	_ sdk.Msg = (*MsgUnfarmAndWithdraw)(nil)
	_ sdk.Msg = (*MsgCancelQueuedFarming)(nil)
	_ sdk.Msg = (*MsgPlaceBid)(nil)
)

// Message types for the module
const (
	TypeMsgFarm                = "farm"
	TypeMsgUnfarm              = "unfarm"
	TypeMsgUnfarmAndWithdraw   = "unfarm_and_withdraw"
	TypeMsgCancelQueuedFarming = "cancel_queued_farming"
	TypeMsgPlaceBid            = "place_bid"
	TypeMsgRefundBid           = "refund_bid"
)

// NewMsgFarm creates a new MsgFarm
func NewMsgFarm(poolId uint64, farmer string, farmingCoin sdk.Coin) *MsgFarm {
	return &MsgFarm{
		PoolId:      poolId,
		Farmer:      farmer,
		FarmingCoin: farmingCoin,
	}
}

func (msg MsgFarm) Route() string { return RouterKey }

func (msg MsgFarm) Type() string { return TypeMsgFarm }

func (msg MsgFarm) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Farmer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid farmer address: %v", err)
	}
	if msg.PoolId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid pool id")
	}
	if !msg.FarmingCoin.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "farming coin must be positive")
	}
	if err := msg.FarmingCoin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid farming coin: %v", err)
	}
	return nil
}

func (msg MsgFarm) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgFarm) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Farmer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgFarm) GetFarmer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Farmer)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgUnfarm creates a new MsgUnfarm
func NewMsgUnfarm(poolId uint64, farmer string, lfCoin sdk.Coin) *MsgUnfarm {
	return &MsgUnfarm{
		PoolId: poolId,
		Farmer: farmer,
		LFCoin: lfCoin,
	}
}

func (msg MsgUnfarm) Route() string { return RouterKey }

func (msg MsgUnfarm) Type() string { return TypeMsgUnfarm }

func (msg MsgUnfarm) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Farmer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid farmer address: %v", err)
	}
	if msg.PoolId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid pool id")
	}
	if !msg.LFCoin.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "liquid farming coin must be positive")
	}
	if err := msg.LFCoin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid liquid farming coin: %v", err)
	}
	return nil
}

func (msg MsgUnfarm) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUnfarm) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Farmer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgUnfarm) GetFarmer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Farmer)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgUnfarmAndWithdraw creates a new MsgUnfarmAndWithdraw
func NewMsgUnfarmAndWithdraw(poolId uint64, farmer string, lfCoin sdk.Coin) *MsgUnfarmAndWithdraw {
	return &MsgUnfarmAndWithdraw{
		PoolId: poolId,
		Farmer: farmer,
		LFCoin: lfCoin,
	}
}

func (msg MsgUnfarmAndWithdraw) Route() string { return RouterKey }

func (msg MsgUnfarmAndWithdraw) Type() string { return TypeMsgUnfarmAndWithdraw }

func (msg MsgUnfarmAndWithdraw) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Farmer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid farmer address: %v", err)
	}
	if msg.PoolId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid pool id")
	}
	if !msg.LFCoin.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "liquid farming coin must be positive")
	}
	if err := msg.LFCoin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid liquid farming coin: %v", err)
	}
	return nil
}

func (msg MsgUnfarmAndWithdraw) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUnfarmAndWithdraw) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Farmer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgUnfarmAndWithdraw) GetFarmer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Farmer)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgCancelQueuedFarming creates a new MsgCancelQueuedFarming
func NewMsgCancelQueuedFarming(poolId uint64, farmer string, unfarmingCoin sdk.Coin) *MsgCancelQueuedFarming {
	return &MsgCancelQueuedFarming{
		PoolId:        poolId,
		Farmer:        farmer,
		UnfarmingCoin: unfarmingCoin,
	}
}

func (msg MsgCancelQueuedFarming) Route() string { return RouterKey }

func (msg MsgCancelQueuedFarming) Type() string { return TypeMsgCancelQueuedFarming }

func (msg MsgCancelQueuedFarming) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Farmer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid farmer address: %v", err)
	}
	if msg.PoolId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid pool id")
	}
	if !msg.UnfarmingCoin.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unfarming coin must be positive")
	}
	if err := msg.UnfarmingCoin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid unfarming coin: %v", err)
	}
	return nil
}

func (msg MsgCancelQueuedFarming) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCancelQueuedFarming) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Farmer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgCancelQueuedFarming) GetFarmer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Farmer)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgPlaceBid creates a new MsgPlaceBid
func NewMsgPlaceBid(poolId uint64, bidder string, biddingCoin sdk.Coin) *MsgPlaceBid {
	return &MsgPlaceBid{
		PoolId:      poolId,
		Bidder:      bidder,
		BiddingCoin: biddingCoin,
	}
}

func (msg MsgPlaceBid) Route() string { return RouterKey }

func (msg MsgPlaceBid) Type() string { return TypeMsgPlaceBid }

func (msg MsgPlaceBid) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Bidder); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bidder address: %v", err)
	}
	if msg.PoolId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid pool id")
	}
	if !msg.BiddingCoin.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "bidding amount must be positive")
	}
	if err := msg.BiddingCoin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid bidding coin: %v", err)
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

// NewMsgRefundBid creates a new MsgRefundBid
func NewMsgRefundBid(poolId uint64, bidder string) *MsgRefundBid {
	return &MsgRefundBid{
		PoolId: poolId,
		Bidder: bidder,
	}
}

func (msg MsgRefundBid) Route() string { return RouterKey }

func (msg MsgRefundBid) Type() string { return TypeMsgRefundBid }

func (msg MsgRefundBid) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Bidder); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bidder address: %v", err)
	}
	if msg.PoolId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid pool id")
	}
	return nil
}

func (msg MsgRefundBid) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgRefundBid) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgRefundBid) GetBidder() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		panic(err)
	}
	return addr
}
