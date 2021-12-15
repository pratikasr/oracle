package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmonaut/oracle/testutil/keeper"
	"github.com/cosmonaut/oracle/x/bandoracle/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BandoracleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
