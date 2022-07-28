package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmosquad-labs/squad/v2/x/lending/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// Lend defines a method for lending assets.
func (k msgServer) Lend(goCtx context.Context, msg *types.MsgLend) (*types.MsgLendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var assetParam types.LendingAssetParam
	found := false
	for _, assetParam = range k.GetLendingAssetParams(ctx) {
		if assetParam.Denom == msg.Coin.Denom {
			found = true
			break
		}
	}
	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrNotFound, "lending asset for denom %s not found", msg.Coin.Denom)
	}

	asset, found := k.GetLendingAsset(ctx, msg.Coin.Denom)
	if !found {
		asset = types.NewLendingAsset(msg.Coin.Denom)
	}

	bondSupply := k.bankKeeper.GetSupply(ctx, asset.BondDenom)
	mintingBondAmt := types.CalculateMintingBondAmount(
		msg.Coin.Amount, asset.TotalLentAmount, asset.AccruedInterestAmount, bondSupply.Amount)
	mintingCoins := sdk.NewCoins(sdk.NewCoin(asset.BondDenom, mintingBondAmt))

	lender := msg.GetLender()
	reserveAddr := asset.GetReserveAddress()
	if err := k.bankKeeper.SendCoins(ctx, lender, reserveAddr, sdk.NewCoins(msg.Coin)); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintingCoins); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lender, mintingCoins); err != nil {
		return nil, err
	}

	asset.TotalLentAmount = asset.TotalLentAmount.Add(msg.Coin.Amount)
	k.SetLendingAsset(ctx, asset)

	// TODO: emit an event

	return &types.MsgLendResponse{}, nil
}

// Withdraw defines a method for withdrawing lent assets.
func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	// TODO: not implemented

	return &types.MsgWithdrawResponse{}, nil
}
