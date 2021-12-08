package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/tendermint/farming/x/liquidstaking/types"
)

// DONTCOVER

// Simulation parameter constants
const (
	EpochBlocks = "epoch_blocks"
	Bearings    = "bearings"
)

// GenEpochBlocks returns randomized epoch blocks.
func GenEpochBlocks(r *rand.Rand) uint32 {
	return uint32(simtypes.RandIntBetween(r, int(types.DefaultEpochBlocks), 10))
}

// GenBearings returns randomized bearings.
func GenBearings(r *rand.Rand) []types.Bearing {
	ranBearings := make([]types.Bearing, 0)

	for i := 0; i < simtypes.RandIntBetween(r, 1, 3); i++ {
		bearing := types.Bearing{
			Name:               "simulation-test-" + simtypes.RandStringOfLength(r, 5),
			Rate:               sdk.NewDecFromIntWithPrec(sdk.NewInt(int64(simtypes.RandIntBetween(r, 1, 4))), 1), // 10~30%
			SourceAddress:      "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",                                   // Cosmos Hub's FeeCollector module account
			DestinationAddress: sdk.AccAddress(address.Module(types.ModuleName, []byte("GravityDEXFarmingBearing"))).String(),
			StartTime:          types.MustParseRFC3339("2000-01-01T00:00:00Z"),
			EndTime:            types.MustParseRFC3339("9999-12-31T00:00:00Z"),
		}
		ranBearings = append(ranBearings, bearing)
	}

	return ranBearings
}

// RandomizedGenState generates a random GenesisState for liquidstaking.
func RandomizedGenState(simState *module.SimulationState) {
	var epochBlocks uint32
	var bearings []types.Bearing
	simState.AppParams.GetOrGenerate(
		simState.Cdc, EpochBlocks, &epochBlocks, simState.Rand,
		func(r *rand.Rand) { epochBlocks = GenEpochBlocks(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, Bearings, &bearings, simState.Rand,
		func(r *rand.Rand) { bearings = GenBearings(r) },
	)

	bearingGenesis := types.GenesisState{
		Params: types.Params{
			EpochBlocks: epochBlocks,
			Bearings:    bearings,
		},
	}

	bz, _ := json.MarshalIndent(&bearingGenesis, "", " ")
	fmt.Printf("Selected randomly generated liquidstaking parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bearingGenesis)
}
