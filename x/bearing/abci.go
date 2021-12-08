package bearing

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/farming/x/bearing/keeper"
	"github.com/tendermint/farming/x/bearing/types"
)

// BeginBlocker collects bearings for the current block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	err := k.CollectBearings(ctx)
	if err != nil {
		panic(err)
	}
}
