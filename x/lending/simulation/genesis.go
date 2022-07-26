package simulation

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/cosmosquad-labs/squad/v2/x/lending/types"
)

// RandomizedGenState generates a random GenesisState for the module.
func RandomizedGenState(simState *module.SimulationState) {
	genesis := types.DefaultGenesis()

	bz, _ := json.MarshalIndent(genesis, "", " ")
	fmt.Printf("Selected randomly generated lending parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(genesis)
}
