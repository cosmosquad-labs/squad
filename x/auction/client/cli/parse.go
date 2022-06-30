package cli

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/spf13/pflag"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	FlagAuciton = "auction"
)

type auction struct {
	StartPrice      sdk.Dec
	SellingCoins    sdk.Coins
	PayingCoinDenom string
	StartTime       time.Time
	EndTime         time.Time
}

func parseCreateAuctionFlags(fs *pflag.FlagSet) (*auction, error) {
	auction := &auction{}
	auctionFile, _ := fs.GetString(FlagAuciton)

	contents, err := ioutil.ReadFile(auctionFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, auction)
	if err != nil {
		return nil, err
	}

	return auction, nil
}
