package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmosquad-labs/squad/x/auction/types"
)

// GetNextAuctionIdWithUpdate increments auction id by one and store it.
func (k Keeper) GetNextAuctionIdWithUpdate(ctx sdk.Context) uint64 {
	id := k.GetLastAuctionId(ctx) + 1
	k.SetAuctionId(ctx, id)
	return id
}

func (k Keeper) CreateAuction(ctx sdk.Context, msg *types.MsgCreateAuction) (types.Auction, error) {
	custom := msg.GetCustom()

	if !k.router.HasRoute(custom.AuctionRoute()) {
		return types.Auction{}, sdkerrors.Wrap(types.ErrNoAuctionHandlerExists, custom.AuctionRoute())
	}

	// Execute the auction content in a new context branch (with branched store)
	// to validate the actual parameter changes before the auction creation.
	// State is not persisted.
	cacheCtx, _ := ctx.CacheContext()
	handler := k.router.GetRoute(custom.AuctionRoute())
	if err := handler(cacheCtx, custom); err != nil {
		return types.Auction{}, sdkerrors.Wrap(types.ErrInvalidAuctionCustom, err.Error())
	}

	if ctx.BlockTime().After(msg.EndTime) { // EndTime < CurrentTime
		return types.Auction{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "end time must be set after the current time")
	}

	auction, err := types.NewAuction(
		custom,
		k.GetNextAuctionIdWithUpdate(ctx),
		msg.Auctioneer,
		types.AuctionStatusStandBy,
		msg.StartTime,
		msg.EndTime,
	)
	if err != nil {
		return types.Auction{}, err
	}

	// Update status if the start time is already passed over the current time
	if !auction.GetStartTime().After(ctx.BlockTime()) {
		auction.Status = types.AuctionStatusStarted
	}

	// Store auction
	k.SetAuctionId(ctx, auction.Id)
	k.SetAuction(ctx, auction)

	return types.Auction{}, nil
}

func (k Keeper) MarshalAuction(auction types.Auction) ([]byte, error) {
	bz, err := k.cdc.Marshal(&auction)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

func (k Keeper) UnmarshalAuction(bz []byte, auction *types.Auction) error {
	err := k.cdc.Unmarshal(bz, auction)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) MustMarshalAuction(auction types.Auction) []byte {
	bz, err := k.MarshalAuction(auction)
	if err != nil {
		panic(err)
	}
	return bz
}

func (k Keeper) MustUnmarshalAuction(bz []byte, auction *types.Auction) {
	err := k.UnmarshalAuction(bz, auction)
	if err != nil {
		panic(err)
	}
}
