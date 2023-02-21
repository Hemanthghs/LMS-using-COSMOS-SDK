package keeper

import (
	//"cosmossdk.io/store/prefix"

	"fmt"

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
	fmt.Println("Student added")
	return nil

}

// func (k Keeper) AddStdnt(ctx sdk.Context, addstudent *types.AddStudentRequest) error {
// 	if addstudent.Name == "" {
// 		return errors.New("name cant be null")
// 	} else if addstudent.Address == "" {
// 		return errors.New("address cant be null")
// 	} else if addstudent.Id == "" {
// 		return errors.New("Id cant be null")
// 	} else {
// 		store := ctx.KVStore(k.storeKey)
// 		// key := types.StudentKey
// 		k.cdc.MustMarshal(addstudent)
// 		marshalAddStudent, err := k.cdc.Marshal(addstudent)
// 		if err != nil {
// 			panic(err)
// 		}
// 		store.Set(types.StudentStoreKey(addstudent.Id), marshalAddStudent)
// 	}
// 	return nil
// }

func (k Keeper) GetStudent(ctx sdk.Context) (a types.Student) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(a.Id))
	k.cdc.MustUnmarshal(bz, &a)
	return a
}
