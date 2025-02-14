package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "gummy/testutil/keeper"
	"gummy/x/gummy/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.GummyKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
