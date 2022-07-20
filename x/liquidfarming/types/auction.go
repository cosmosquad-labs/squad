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

	// The module uses the address type of 32 bytes length, but it can always be changed depending on Cosmos SDK's direction.
	ReserveAddressType = farmingtypes.AddressType32Bytes
)

// NewRewardsAuction creates a new RewardsAuction.
func NewRewardsAuction(
	id uint64,
	poolId uint64,
	biddingCoinDenom string,
	startTime time.Time,
	endTime time.Time,
) RewardsAuction {
	return RewardsAuction{
		Id:                    id,
		PoolId:                poolId,
		BiddingCoinDenom:      biddingCoinDenom,
		SellingReserveAddress: SellingReserveAddress(poolId).String(),
		PayingReserveAddress:  PayingReserveAddress(poolId).String(),
		StartTime:             startTime,
		EndTime:               endTime,
		Status:                AuctionStatusStarted,
		Winner:                "",
		Rewards:               sdk.Coins{},
	}
}

// Validate validates RewardsAuction.
func (a *RewardsAuction) Validate() error {
	if a.BiddingCoinDenom == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "bidding coin denom cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(a.SellingReserveAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid selling reserve address: %v", err)
	}
	if _, err := sdk.AccAddressFromBech32(a.PayingReserveAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid paying reserve address: %v", err)
	}
	if !a.EndTime.After(a.StartTime) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "end time must be set after than start time")
	}
	if a.Status != AuctionStatusStarted {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "auction status must be set correctly")
	}
	return nil
}

// SetStatus sets rewards auction status.
func (a *RewardsAuction) SetStatus(status AuctionStatus) {
	a.Status = status
}

// SetWinner sets winner address.
func (a *RewardsAuction) SetWinner(winner string) {
	a.Winner = winner
}

// SetRewards sets auction rewards.
func (a *RewardsAuction) SetRewards(rewards sdk.Coins) {
	a.Rewards = rewards
}

// GetSellingReserveAddress returns the selling reserve address in the form of sdk.AccAddress.
func (a RewardsAuction) GetSellingReserveAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(a.SellingReserveAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// GetPayingReserveAddress returns the paying reserve address in the form of sdk.AccAddress.
func (a RewardsAuction) GetPayingReserveAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(a.PayingReserveAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewBid creates a new Bid.
func NewBid(poolId uint64, bidder string, amount sdk.Coin) Bid {
	return Bid{
		PoolId: poolId,
		Bidder: bidder,
		Amount: amount,
	}
}

// GetBidder returns the bidder address in the form of sdk.AccAddress.
func (b Bid) GetBidder() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(b.Bidder)
	if err != nil {
		panic(err)
	}
	return addr
}

// MustMarshalRewardsAuction marshals RewardsAuction and
// it panics upon failure.
func MustMarshalRewardsAuction(cdc codec.BinaryCodec, auction RewardsAuction) []byte {
	return cdc.MustMarshal(&auction)
}

// MustUnmarshalRewardsAuction unmarshals RewardsAuction and
// it panics upon failure.
func MustUnmarshalRewardsAuction(cdc codec.BinaryCodec, value []byte) RewardsAuction {
	pair, err := UnmarshalRewardsAuction(cdc, value)
	if err != nil {
		panic(err)
	}
	return pair
}

// UnmarshalRewardsAuction unmarshals RewardsAuction.
func UnmarshalRewardsAuction(cdc codec.BinaryCodec, value []byte) (auction RewardsAuction, err error) {
	err = cdc.Unmarshal(value, &auction)
	return auction, err
}

// SellingReserveAddress creates new selling reserve address in the form of sdk.AccAddress
// with the given pool id.
func SellingReserveAddress(poolId uint64) sdk.AccAddress {
	return farmingtypes.DeriveAddress(
		ReserveAddressType,
		ModuleName,
		strings.Join([]string{SellingReserveAddressPrefix, strconv.FormatUint(poolId, 10)}, ModuleAddressNameSplitter),
	)
}

// PayingReserveAddress creates the paying reserve address in the form of sdk.AccAddress
// with the given pool id.
func PayingReserveAddress(poolId uint64) sdk.AccAddress {
	return farmingtypes.DeriveAddress(
		ReserveAddressType,
		ModuleName,
		strings.Join([]string{PayingReserveAddressPrefix, strconv.FormatUint(poolId, 10)}, ModuleAddressNameSplitter),
	)
}
