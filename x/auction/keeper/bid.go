package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/x/auction/types"
)

// PlaceBid places a bid for the selling coin of the auction.
func (k Keeper) PlaceBid(ctx sdk.Context, msg *types.MsgPlaceBid) (types.Bid, error) {
	// TODO: not implemented yet
	return types.Bid{}, nil
}
