package keeper // noalias

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmosquad-labs/squad/x/liquidstaking/types"
)

// RandomValidator returns a random validator given access to the keeper and ctx
func RandomValidator(r *rand.Rand, keeper types.StakingKeeper, ctx sdk.Context) (val stakingtypes.Validator, ok bool) {
	vals := keeper.GetAllValidators(ctx)
	if len(vals) == 0 {
		return stakingtypes.Validator{}, false
	}

	i := r.Intn(len(vals))

	return vals[i], true
}
