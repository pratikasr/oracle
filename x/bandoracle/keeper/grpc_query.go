package keeper

import (
	"github.com/cosmonaut/oracle/x/bandoracle/types"
)

var _ types.QueryServer = Keeper{}
