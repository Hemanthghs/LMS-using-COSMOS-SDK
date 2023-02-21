package keeper

import (
	//"cosmossdk.io/store/prefix"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetStudent(ctx sdk.Context, a types.Student) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&a)
	if err != nil {
		return err
	}
	classID := a.Id
	key := make([]byte, len(types.PrefstudentId)+len(classID))
	copy(key, types.PrefstudentId)
	copy(key[len(types.PrefstudentId):], []byte(classID))
	store.Set(key, bz)
	return nil

}
func (k Keeper) GetStudent(ctx sdk.Context) (a types.Student) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(a.Id))
	k.cdc.MustUnmarshal(bz, &a)
	return a
}
