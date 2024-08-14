package gummy_test

import (
	"testing"

	keepertest "gummy/testutil/keeper"
	"gummy/testutil/nullify"
	gummy "gummy/x/gummy/module"
	"gummy/x/gummy/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GummyKeeper(t)
	gummy.InitGenesis(ctx, k, genesisState)
	got := gummy.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
