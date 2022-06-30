package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Custom defines an interface that an auction must implement.
type Custom interface {
	GetStartPrice() sdk.Dec
	GetSellingCoins() sdk.Coins
	GetPayingCoinDenom() string
	GetStartTime() time.Time
	GetEndTime() time.Time
	AuctionRoute() string
	AuctionType() string
	ValidateBasic() error
}

type Handler func(ctx sdk.Context, custom Custom) error

func ValidateAbstract(c Custom) error {
	// TODO: not implemented yet
	return nil
}
