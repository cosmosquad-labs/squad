package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	farmingtypes "github.com/cosmosquad-labs/squad/v2/x/farming/types"
)

// Wrapper struct
type Hooks struct {
	k Keeper
}

var _ farmingtypes.FarmingHooks = Hooks{}

// Create new liquidfarming hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

func (h Hooks) AfterStaked(ctx sdk.Context, farmerAcc sdk.AccAddress, stakingCoinDenom string, newStakingAmt sdk.Int) {
	// TODO: not implemented yet
	// Get deposit requests, Mint
}

func (h Hooks) AfterAllocateRewards(ctx sdk.Context) {
	// TODO: not implemented yet
	// Select winner -> Distribute rewards -> Refund (?)
	// Stake with the pool coin (auto-compounding role)
	// Create RewardsAuction
}
