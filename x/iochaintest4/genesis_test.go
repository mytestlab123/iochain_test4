package iochaintest4_test

import (
	"testing"

	keepertest "github.com/mytestlab123/iochain_test4/testutil/keeper"
	"github.com/mytestlab123/iochain_test4/testutil/nullify"
	"github.com/mytestlab123/iochain_test4/x/iochaintest4"
	"github.com/mytestlab123/iochain_test4/x/iochaintest4/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Iochaintest4Keeper(t)
	iochaintest4.InitGenesis(ctx, *k, genesisState)
	got := iochaintest4.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
