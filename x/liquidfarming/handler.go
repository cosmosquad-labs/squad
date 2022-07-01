package liquidfarming

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	auctiontypes "github.com/cosmosquad-labs/squad/x/auction/types"
	"github.com/cosmosquad-labs/squad/x/liquidfarming/keeper"
	"github.com/cosmosquad-labs/squad/x/liquidfarming/types"
)

func NewCreateFixedPriceAuctionHandler(k keeper.Keeper) auctiontypes.Handler {
	return func(ctx sdk.Context, custom auctiontypes.Custom) error {
		fmt.Println(">>> NewCreateFixedPriceAuctionHandler...", custom.AuctionType())

		switch c := custom.(type) {
		case *types.FixedPriceAuction:
			return handleCreateFixedPriceAuction(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized auction custom type: %T", c)
		}
	}
}

func handleCreateFixedPriceAuction(ctx sdk.Context, k keeper.Keeper, a *types.FixedPriceAuction) error {
	fmt.Println(">>> Handling CreateFixedPriceAuction...", a)

	// TODO: not implemented yet
	// Store

	return nil
}
