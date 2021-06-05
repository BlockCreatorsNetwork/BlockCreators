package keeper

import (
	"github.com/BlockCreatorsNetwork/BlockCreators/x/blockcreators/types"
)

var _ types.QueryServer = Keeper{}
