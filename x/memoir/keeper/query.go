package keeper

import (
	"memoir/x/memoir/types"
)

var _ types.QueryServer = Keeper{}
