package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmonaut/oracle/testutil/keeper"
	"github.com/cosmonaut/oracle/x/bandoracle/keeper"
	"github.com/cosmonaut/oracle/x/bandoracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BandoracleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
