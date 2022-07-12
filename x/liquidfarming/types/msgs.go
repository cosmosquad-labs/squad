package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgFarm)(nil)
	_ sdk.Msg = (*MsgCancelQueuedFarming)(nil)
	_ sdk.Msg = (*MsgUnfarm)(nil)
	_ sdk.Msg = (*MsgPlaceBid)(nil)
)

// Message types for the farming module
const (
	TypeMsgFarm                             = "farm"
	TypeMsgCancelQueuedFarmingQueuedFarming = "cancel_queued_farming"
	TypeMsgUnfarm                           = "unfarm"
	TypeMsgPlaceBid                         = "place_bid"
	TypeMsgRefundBid                        = "refund_bid"
)

// NewMsgFarm returns a new MsgFarm.
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

// NewMsgCancelQueuedFarming returns a new MsgCancelQueuedFarming.
func NewMsgCancelQueuedFarming(farmer string, poolId, queuedFarmingId uint64) *MsgCancelQueuedFarming {
	return &MsgCancelQueuedFarming{
		Farmer:          farmer,
		PoolId:          poolId,
		QueuedFarmingId: queuedFarmingId,
	}
}

func (msg MsgCancelQueuedFarming) Route() string { return RouterKey }

func (msg MsgCancelQueuedFarming) Type() string { return TypeMsgCancelQueuedFarmingQueuedFarming }

func (msg MsgCancelQueuedFarming) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Farmer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid farmer address: %v", err)
	}
	if msg.PoolId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid pool id")
	}
	if msg.QueuedFarmingId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid queued farming id")
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

// NewMsgUnfarm returns a new MsgUnfarm.
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

func NewMsgRefundBid(auctionId, bidId uint64, bidder string) *MsgRefundBid {
	return &MsgRefundBid{
		AuctionId: auctionId,
		BidId:     bidId,
		Bidder:    bidder,
	}
}

func (msg MsgRefundBid) Route() string { return RouterKey }

func (msg MsgRefundBid) Type() string { return TypeMsgRefundBid }

func (msg MsgRefundBid) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Bidder); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bidder address: %v", err)
	}
	if msg.AuctionId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid auction id")
	}
	if msg.BidId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid bid id")
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