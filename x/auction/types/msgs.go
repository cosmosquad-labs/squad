package types

import (
	time "time"

	"github.com/cosmos/cosmos-sdk/codec/legacy"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgCreateFixedPriceAuction)(nil)
	_ sdk.Msg = (*MsgCancelAuction)(nil)
	_ sdk.Msg = (*MsgPlaceBid)(nil)
)

// Message types for the fundraising module.
const (
	TypeMsgCreateFixedPriceAuction = "create_fixed_price_auction"
	TypeMsgCancelAuction           = "cancel_auction"
	TypeMsgPlaceBid                = "place_bid"
)

// NewMsgCreateFixedPriceAuction creates a new MsgCreateFixedPriceAuction.
func NewMsgCreateFixedPriceAuction(
	auctioneer string,
	startPrice sdk.Dec,
	sellingCoins sdk.Coins,
	payingCoinDenom string,
	vestingSchedules []VestingSchedule,
	startTime time.Time,
	endTime time.Time,
) *MsgCreateFixedPriceAuction {
	return &MsgCreateFixedPriceAuction{
		Auctioneer:       auctioneer,
		StartPrice:       startPrice,
		SellingCoins:     sellingCoins,
		PayingCoinDenom:  payingCoinDenom,
		VestingSchedules: vestingSchedules,
		StartTime:        startTime,
		EndTime:          endTime,
	}
}

func (msg MsgCreateFixedPriceAuction) Route() string { return RouterKey }

func (msg MsgCreateFixedPriceAuction) Type() string { return TypeMsgCreateFixedPriceAuction }

func (msg MsgCreateFixedPriceAuction) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Auctioneer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid auctioneer address: %v", err)
	}
	if !msg.StartPrice.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "start price must be positive")
	}
	if err := msg.SellingCoins.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid selling coin: %v", err)
	}
	if !msg.SellingCoins.IsAnyNegative() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "selling coin amount must be positive")
	}
	// TODO: reconsider this validation
	// if msg.SellingCoin.Denom == msg.PayingCoinDenom {
	// 	return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "selling coin denom must not be the same as paying coin denom")
	// }
	if err := sdk.ValidateDenom(msg.PayingCoinDenom); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid paying coin denom: %v", err)
	}
	if !msg.EndTime.After(msg.StartTime) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "end time must be set after start time")
	}
	if err := ValidateVestingSchedules(msg.VestingSchedules, msg.EndTime); err != nil {
		return err
	}
	return nil
}

func (msg MsgCreateFixedPriceAuction) GetSignBytes() []byte {
	return sdk.MustSortJSON(legacy.Cdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateFixedPriceAuction) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Auctioneer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgCreateFixedPriceAuction) GetAuctioneer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Auctioneer)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgCancelAuction creates a new MsgCancelAuction.
func NewMsgCancelAuction(
	auctioneer string,
	auctionId uint64,
) *MsgCancelAuction {
	return &MsgCancelAuction{
		Auctioneer: auctioneer,
		AuctionId:  auctionId,
	}
}

func (msg MsgCancelAuction) Route() string { return RouterKey }

func (msg MsgCancelAuction) Type() string { return TypeMsgCancelAuction }

func (msg MsgCancelAuction) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Auctioneer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid auctioneer address %q: %v", msg.Auctioneer, err)
	}
	return nil
}

func (msg MsgCancelAuction) GetSignBytes() []byte {
	return sdk.MustSortJSON(legacy.Cdc.MustMarshalJSON(&msg))
}

func (msg MsgCancelAuction) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Auctioneer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgCancelAuction) GetAuctioneer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Auctioneer)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgPlaceBid creates a new MsgPlaceBid.
func NewMsgPlaceBid(
	auctionId uint64,
	bidder string,
	bidType BidType,
	Price sdk.Dec,
	Coin sdk.Coin,
) *MsgPlaceBid {
	return &MsgPlaceBid{
		AuctionId: auctionId,
		Bidder:    bidder,
		BidType:   bidType,
		Price:     Price,
		Coin:      Coin,
	}
}

func (msg MsgPlaceBid) Route() string { return RouterKey }

func (msg MsgPlaceBid) Type() string { return TypeMsgPlaceBid }

func (msg MsgPlaceBid) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Bidder); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bidder address: %v", err)
	}
	if !msg.Price.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "bid price must be positive value")
	}
	if err := msg.Coin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid bid coin: %v", err)
	}
	if !msg.Coin.Amount.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid coin amount: %s", msg.Coin.Amount.String())
	}
	if msg.BidType != BidTypeFixedPrice && msg.BidType != BidTypeBatchWorth &&
		msg.BidType != BidTypeBatchMany {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid bid type: %T", msg.BidType.String())
	}
	return nil
}

func (msg MsgPlaceBid) GetSignBytes() []byte {
	return sdk.MustSortJSON(legacy.Cdc.MustMarshalJSON(&msg))
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
