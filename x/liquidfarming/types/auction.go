package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	auctiontypes "github.com/cosmosquad-labs/squad/x/auction/types"
)

const (
	AuctionTypeFixedPrice = "FixedPriceAuction"
)

var _ auctiontypes.Custom = &FixedPriceAuction{}

func init() {
	auctiontypes.RegisterAuctionType(AuctionTypeFixedPrice)
	auctiontypes.RegisterAuctionTypeCodec(&FixedPriceAuction{}, "squad/FixedPriceAuction")
}

func NewFixedPriceAuction(
	auctioneer string,
	startPrice sdk.Dec,
	sellingCoin sdk.Coin,
	remainingSellingCoin sdk.Coin,
	payingCoinDenom string,
	startTime time.Time,
	endTime time.Time,
	status AuctionStatus,
) auctiontypes.Custom {
	return &FixedPriceAuction{
		Auctioneer:      auctioneer,
		StartPrice:      startPrice,
		SellingCoin:     sellingCoin,
		PayingCoinDenom: payingCoinDenom,
		StartTime:       startTime,
		EndTime:         endTime,
		Status:          status,
	}
}

func (a *FixedPriceAuction) GetAuctioneer() string      { return a.Auctioneer }
func (a *FixedPriceAuction) GetStartPrice() sdk.Dec     { return a.StartPrice }
func (a *FixedPriceAuction) GetSellingCoins() sdk.Coins { return sdk.NewCoins(a.SellingCoin) }
func (a *FixedPriceAuction) GetPayingCoinDenom() string { return a.PayingCoinDenom }
func (a *FixedPriceAuction) GetStartTime() time.Time    { return a.StartTime }
func (a *FixedPriceAuction) GetEndTime() time.Time      { return a.EndTime }
func (a *FixedPriceAuction) AuctionRoute() string       { return RouterKey }
func (a *FixedPriceAuction) AuctionType() string        { return AuctionTypeFixedPrice }
func (a *FixedPriceAuction) ValidateBasic() error       { return nil }
