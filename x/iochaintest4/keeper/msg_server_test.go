package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mytestlab123/iochain_test4/testutil/keeper"
	"github.com/mytestlab123/iochain_test4/x/iochaintest4/keeper"
	"github.com/mytestlab123/iochain_test4/x/iochaintest4/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Iochaintest4Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
