package types

import (
	"fmt"
	time "time"

	"github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const DefaultStartingAuctionID uint64 = 1

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
		PayingReserveAddress:  SellingReserveAddress(id).String(),
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

// TODO: proto file must be updated to getters all false
// func (a Auction) String() string {
// 	out, _ := yaml.Marshal(p)
// 	return string(out)
// }

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	AuctionTypeFixedPrice string = "FixedPrice"
)

// Implements Custom interface
var _ Custom = &FixedPriceAuction{}

func NewFixedPriceAuction(
	auctioneer string,
	startPrice sdk.Dec,
	sellingCoin sdk.Coin,
	remainingSellingCoin sdk.Coin,
	payingCoinDenom string,
	startTime time.Time,
	endTime time.Time,
	status AuctionStatus,
) Custom {
	return &FixedPriceAuction{
		Auctioneer:           auctioneer,
		StartPrice:           startPrice,
		SellingCoin:          sellingCoin,
		RemainingSellingCoin: remainingSellingCoin,
		PayingCoinDenom:      payingCoinDenom,
		StartTime:            startTime,
		EndTime:              endTime,
		Status:               status,
	}
}

func (a *FixedPriceAuction) GetAuctioneer() string {
	return a.Auctioneer
}

func (a *FixedPriceAuction) GetStartPrice() sdk.Dec {
	return a.StartPrice
}

func (a *FixedPriceAuction) GetSellingCoins() sdk.Coins {
	return sdk.NewCoins(a.SellingCoin)
}

func (a *FixedPriceAuction) GetPayingCoinDenom() string {
	return a.PayingCoinDenom
}

func (a *FixedPriceAuction) GetStartTime() time.Time {
	return a.StartTime
}

func (a *FixedPriceAuction) GetEndTime() time.Time {
	return a.EndTime
}

func (a *FixedPriceAuction) GetStatus() AuctionStatus {
	return a.Status
}

func (a *FixedPriceAuction) AuctionRoute() string { return RouterKey }

func (a *FixedPriceAuction) AuctionType() string { return AuctionTypeFixedPrice }

func (a *FixedPriceAuction) ValidateBasic() error { return ValidateAbstract(a) }
