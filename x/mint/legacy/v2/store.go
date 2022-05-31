package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmosquad-labs/squad/x/mint/types"
)

func MigrateStore(ctx sdk.Context, paramSpace paramtypes.Subspace) error {
	migrateParamsStore(ctx, paramSpace)
	return nil
}

func migrateParamsStore(ctx sdk.Context, paramSpace paramtypes.Subspace) {
	paramSpace.Set(ctx, types.KeyMintPoolAddress, types.DefaultMintPoolAddress.String())
}
