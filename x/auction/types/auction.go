package types

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewAuction(custom Custom, id uint64, auctioneer string) (Auction, error) {
	msg, ok := custom.(proto.Message)
	if !ok {
		return Auction{}, fmt.Errorf("%T does not implement proto.Message", custom)
	}

	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return Auction{}, err
	}

	auction := Auction{
		Custom:                any,
		Id:                    id,
		Auctioneer:            auctioneer,
		SellingReserveAddress: SellingReserveAddress(id).String(),
		PayingReserveAddress:  PayingReserveAddress(id).String(),
	}

	return auction, nil
}

// GetCustom returns the auction Custom
func (a Auction) GetCustom() Custom {
	custom, ok := a.Custom.GetCachedValue().(Custom)
	if !ok {
		return nil
	}
	return custom
}

func (a Auction) AuctionId() uint64 {
	return a.Id
}

func (a Auction) AuctionType() string {
	custom := a.GetCustom()
	if custom == nil {
		return ""
	}
	return custom.AuctionType()
}

func (a Auction) GetAuctioneer() string {
	addr, err := sdk.AccAddressFromBech32(a.Auctioneer)
	if err != nil {
		panic(err)
	}
	return addr.String()
}

var validAuctionTypes = map[string]struct{}{}

func RegisterAuctionType(ty string) {
	if _, ok := validAuctionTypes[ty]; ok {
		panic(fmt.Sprintf("already registered auction type: %s", ty))
	}

	validAuctionTypes[ty] = struct{}{}
}

// AuctionHandler implements the Handler interface for auction module-based
// auctions (ie. FixedPriceAuction ). Since these are
// merely signaling mechanisms at the moment and do not affect state, it
// performs a no-op.
func AuctionHandler(ctx sdk.Context, c Custom) error {
	ctx.Logger().Debug(">>> AuctionHandler...")
	return nil
}
