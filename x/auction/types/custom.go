package types

import (
	"strings"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Custom defines an interface that an auction must implement.
type Custom interface {
	GetId() uint64
	GetAuctioneer() string
	AuctionType() string
	GetStatus() AuctionStatus
	GetStartTime() time.Time
	GetEndTime() time.Time
	// ProposalRoute() string
}

func ValidateAbstract(c Custom) error {
	if c.GetId() == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid auction id")
	}
	if len(strings.TrimSpace(c.GetAuctioneer())) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "auctioneer cannot be blank")
	}
	if c.GetStatus() == AuctionStatusStandBy || c.GetStatus() == AuctionStatusStarted {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid auction status")
	}
	return nil
}
