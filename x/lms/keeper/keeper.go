package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	// "github.com/cosmos/gogoproto/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey

	cdc codec.Codec
}

func NewKeeper(key storetypes.StoreKey, cdc codec.Codec) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: key,
	}
}
