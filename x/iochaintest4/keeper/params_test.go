package keeper_test

import (
	"testing"

	testkeeper "github.com/mytestlab123/iochain_test4/testutil/keeper"
	"github.com/mytestlab123/iochain_test4/x/iochaintest4/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Iochaintest4Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
