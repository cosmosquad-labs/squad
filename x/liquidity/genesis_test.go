package liquidity_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/tendermint/farming/testutil/keeper"
	"github.com/tendermint/farming/x/liquidity"
	"github.com/tendermint/farming/x/liquidity/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LiquidityKeeper(t)
	liquidity.InitGenesis(ctx, *k, genesisState)
	got := liquidity.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
