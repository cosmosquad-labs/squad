package simulation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/cosmosquad-labs/squad/v2/x/lending/keeper"
	"github.com/cosmosquad-labs/squad/v2/x/lending/types"
)

// Simulation operation weights constants.
const ()

// WeightedOperations returns all the operations from the module with their respective weights.
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, ak types.AccountKeeper,
	bk types.BankKeeper, k keeper.Keeper,
) simulation.WeightedOperations {

	return simulation.WeightedOperations{}
}
