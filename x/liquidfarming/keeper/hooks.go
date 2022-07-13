package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	farmingtypes "github.com/cosmosquad-labs/squad/v2/x/farming/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v2/x/liquidity/types"
)

// Wrapper struct
type Hooks struct {
	k Keeper
}

var _ farmingtypes.FarmingHooks = Hooks{}

// Create new liquidfarming hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

func (h Hooks) AfterStaked(ctx sdk.Context, reserveAcc sdk.AccAddress, stakingCoinDenom string, newStakingAmt sdk.Int) {
	farmer := sdk.AccAddress{}
	mintingCoin := sdk.Coin{}
	h.k.IterateMatureQueuedFarmings(ctx, ctx.BlockTime(), func(endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress, queuedFarming types.QueuedFarming) (stop bool) {
		poolCoinDenom := liquiditytypes.PoolCoinDenom(queuedFarming.PoolId)
		lfCoinDenom := types.LFCoinDenom(queuedFarming.PoolId)

		if poolCoinDenom == farmingCoinDenom && newStakingAmt.Equal(queuedFarming.Amount) {
			lfCoinTotalSupply := h.k.bankKeeper.GetSupply(ctx, lfCoinDenom).Amount
			poolCoinBalance := h.k.bankKeeper.SpendableCoins(ctx, reserveAcc).AmountOf(poolCoinDenom)
			farmedAmt := queuedFarming.Amount

			if poolCoinBalance.IsZero() || lfCoinTotalSupply.IsZero() { // first case
				mintingCoin = sdk.NewCoin(lfCoinDenom, farmedAmt)
			} else {
				// MintingAmount = TotalLFCoinSupply / TotalLPCoinStaked * LPCoinDeposit
				mintingAmt := lfCoinTotalSupply.Quo(poolCoinBalance).Mul(farmedAmt)
				mintingCoin = sdk.NewCoin(lfCoinDenom, mintingAmt)
			}
			fmt.Println("poolCoinBalance: ", poolCoinBalance)
			fmt.Println("lfCoinTotalSupply: ", lfCoinTotalSupply)
			fmt.Println("mintingCoin: ", mintingCoin)
			fmt.Println("")

			farmer = farmerAcc

			return true
		}
		return false
	})

	// Mint liquid farming coin
	if err := h.k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(mintingCoin)); err != nil {
		panic(err)
	}

	// Send LFCoin to the farmer
	if err := h.k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, farmer, sdk.NewCoins(mintingCoin)); err != nil {
		panic(err)
	}

	// TODO: emit events?
}

func (h Hooks) AfterAllocateRewards(ctx sdk.Context) {
	// TODO: not implemented yet
	// Select winner -> Distribute rewards -> Refund (?)
	// Stake with the pool coin (auto-compounding role)
	// Create RewardsAuction
}
