package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmosquad-labs/squad/x/mint/types"
)

// RandomizedGenState generates a random GenesisState for mint
func RandomizedGenState(simState *module.SimulationState) {
	mintGenesis := types.DefaultGenesisState()
	bz, err := json.MarshalIndent(&mintGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated mint parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(mintGenesis)
}
