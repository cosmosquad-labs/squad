package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmosquad-labs/squad/x/liquidfarming/types"
    "github.com/cosmosquad-labs/squad/x/liquidfarming/keeper"
    keepertest "github.com/cosmosquad-labs/squad/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LiquidfarmingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
