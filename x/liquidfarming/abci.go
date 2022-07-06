package liquidfarming

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/x/liquidfarming/keeper"
	"github.com/cosmosquad-labs/squad/x/liquidfarming/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// TODO: not implemented yet
	// params := k.GetParams(ctx)
	// for _, liquidFarm := range params.LiquidFarms {
	// 	liquidFarm.
	// }
}
