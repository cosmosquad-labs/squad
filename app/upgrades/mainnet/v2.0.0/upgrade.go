package v2_0_0

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	budgetkeeper "github.com/tendermint/budget/x/budget/keeper"
)

const UpgradeName = "v2.0.0"

func UpgradeHandler(mm *module.Manager, configurator module.Configurator, budgetKeeper budgetkeeper.Keeper) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		newVM, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return newVM, err
		}

		// add budget migration code related mint pool

		return newVM, err
	}
}

var StoreUpgrades store.StoreUpgrades
