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

// func (k BaseKeeper) SetStudent(ctx sdk.Context, name string, id string, email string, passwd string) {
// 	store :=
// }
