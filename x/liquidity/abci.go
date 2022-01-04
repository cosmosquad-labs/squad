package liquidity

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/farming/x/liquidity/keeper"
	"github.com/tendermint/farming/x/liquidity/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// logger := k.Logger(ctx)

	params := k.GetParams(ctx)
	if ctx.BlockHeight()%int64(params.BatchSize) == 0 {
		// TODO: run batch logic
	}
}
