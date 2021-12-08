package simulation_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/farming/x/liquidstaking/simulation"
	"github.com/tendermint/farming/x/liquidstaking/types"
)

func TestDecodeBearingStore(t *testing.T) {

	cdc := simapp.MakeTestEncodingConfig()
	dec := simulation.NewDecodeStore(cdc.Marshaler)

	tc := types.TotalCollectedCoins{
		TotalCollectedCoins: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1000000))),
	}

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: types.TotalCollectedCoinsKeyPrefix, Value: cdc.Marshaler.MustMarshal(&tc)},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"totalCollectedCoins", fmt.Sprintf("%v\n%v", tc, tc)},
		{"other", ""},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
