package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "voter/testutil/keeper"
	"voter/x/voter/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.VoterKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
