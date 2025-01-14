package memoir_test

import (
	"testing"

	keepertest "memoir/testutil/keeper"
	"memoir/testutil/nullify"
	memoir "memoir/x/memoir/module"
	"memoir/x/memoir/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MemoirKeeper(t)
	memoir.InitGenesis(ctx, k, genesisState)
	got := memoir.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
