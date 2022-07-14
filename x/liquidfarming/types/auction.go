package types

import (
	"strconv"
	"strings"
	time "time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	farmingtypes "github.com/cosmosquad-labs/squad/v2/x/farming/types"
)

const (
	SellingReserveAddressPrefix string = "SellingReserveAddress"
	PayingReserveAddressPrefix  string = "PayingReserveAddress"
	ModuleAddressNameSplitter   string = "|"

	ReserveAddressType = farmingtypes.AddressType32Bytes
)

func NewRewardsAuction(
	id uint64,
	poolId uint64,
	sellingRewards sdk.Coins,
	biddingCoinDenom string,
	startTime time.Time,
	endTime time.Time,
) RewardsAuction {
	return RewardsAuction{
		Id:                    id,
		PoolId:                poolId,
		SellingRewards:        sellingRewards,
		BiddingCoinDenom:      biddingCoinDenom,
		SellingReserveAddress: SellingReserveAddress(poolId).String(),
		PayingReserveAddress:  PayingReserveAddress(poolId).String(),
		StartTime:             startTime,
		EndTime:               endTime,
		Status:                AuctionStatusStarted,
		WinnerBidId:           0,
	}
}

func (ra *RewardsAuction) Validate() error {
	if err := ra.SellingRewards.Validate(); err != nil {
		return sdkerrors.Wrapf(err, "invalid rewards")
	}
	if ra.BiddingCoinDenom == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "bidding coin denom cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(ra.SellingReserveAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid selling reserve address: %v", err)
	}
	if _, err := sdk.AccAddressFromBech32(ra.PayingReserveAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid paying reserve address: %v", err)
	}
	if !ra.EndTime.After(ra.StartTime) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "end time must be set after than start time")
	}
	if ra.Status != AuctionStatusStarted {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "auction status must be set correctly")
	}
	return nil
}

func MustMarshalRewardsAuction(cdc codec.BinaryCodec, auction RewardsAuction) []byte {
	return cdc.MustMarshal(&auction)
}

func MustUnmarshalRewardsAuction(cdc codec.BinaryCodec, value []byte) RewardsAuction {
	pair, err := UnmarshalRewardsAuction(cdc, value)
	if err != nil {
		panic(err)
	}
	return pair
}

func UnmarshalRewardsAuction(cdc codec.BinaryCodec, value []byte) (auction RewardsAuction, err error) {
	err = cdc.UnmarshalInterface(value, &auction)
	return auction, err
}

// SellingReserveAddress returns the selling reserve address with the given pool id.
func SellingReserveAddress(poolId uint64) sdk.AccAddress {
	return farmingtypes.DeriveAddress(
		ReserveAddressType,
		ModuleName,
		strings.Join([]string{SellingReserveAddressPrefix, strconv.FormatUint(poolId, 10)}, ModuleAddressNameSplitter),
	)
}

// PayingReserveAddress returns the paying reserve address with the given pool id.
func PayingReserveAddress(poolId uint64) sdk.AccAddress {
	return farmingtypes.DeriveAddress(
		ReserveAddressType,
		ModuleName,
		strings.Join([]string{PayingReserveAddressPrefix, strconv.FormatUint(poolId, 10)}, ModuleAddressNameSplitter),
	)
}
