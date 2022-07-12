package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

func (k Keeper) PlaceBid(ctx sdk.Context, msg *types.MsgPlaceBid) error {
	// TODO: not implemented yet
	// Check for minimum deposit amount
	// Check if the bidder has enough amount of coin to place a bid
	// Check if the auction exists
	// Check if the bid amount is greater than the currently winning bid amount

	return nil
}

func (k Keeper) RefundBid(ctx sdk.Context, msg *types.MsgRefundBid) error {
	// TODO: not implemented yet
	// Winning bid can't be refunded

	return nil
}
