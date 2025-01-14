package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "memoir/testutil/keeper"
	"memoir/x/memoir/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.MemoirKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
