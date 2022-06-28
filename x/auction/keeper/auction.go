package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/x/auction/types"
)

func (k Keeper) CreateFixedPriceAuction(ctx sdk.Context, msg *types.MsgCreateFixedPriceAuction) (types.AuctionI, error) {
	// TODO: not implemented yet
	return nil, nil
}

func (k Keeper) CancelAuction(ctx sdk.Context, msg *types.MsgCancelAuction) error {
	// TODO: not implemented yet
	return nil
}
