package utils

import (
	"io/ioutil"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type (
	FixedPriceAuctionJSON struct {
		StartPrice      sdk.Dec   `json:"start_price" yaml:"start_price"`
		SellingCoin     sdk.Coin  `json:"selling_coin" yaml:"selling_coin"`
		PayingCoinDenom string    `json:"paying_coin_denom" yaml:"paying_coin_denom"`
		StartTime       time.Time `json:"start_time" yaml:"start_time"`
		EndTime         time.Time `json:"end_time" yaml:"end_time"`
	}

	FixedPriceAuctionReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	}
)

func ParseFixedPriceAuctionJSON(cdc *codec.LegacyAmino, auctionFile string) (FixedPriceAuctionJSON, error) {
	auction := FixedPriceAuctionJSON{}

	custom, err := ioutil.ReadFile(auctionFile)
	if err != nil {
		return auction, err
	}

	if err := cdc.UnmarshalJSON(custom, &auction); err != nil {
		return auction, err
	}

	return auction, nil
}
