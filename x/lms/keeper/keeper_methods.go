package keeper

import (
	//"cosmossdk.io/store/prefix"

	"fmt"
	"log"
	"strconv"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (k Keeper) RegisterAdmin(ctx sdk.Context, registerAdmin *types.RegisterAdminRequest) string {

	if registerAdmin.Name == "" {
		return "Name cannot be empty"
	} else if registerAdmin.Address == "" {
		return "Address cannot be empty"
	} else {
		store := ctx.KVStore(k.storeKey)
		// k.cdc.MustMarshal(registerAdmin)
		marshalRegisterAdmin, err := k.cdc.Marshal(registerAdmin)
		handleError(err)
		store.Set(types.AdminStoreKey(registerAdmin.Address), marshalRegisterAdmin)
		return "Admin Registered Successfully"
	}
}

func (k Keeper) GetAdmin(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	fmt.Println(types.AdminStoreKey(id))
	v := store.Get(types.AdminStoreKey(id))
	fmt.Println(v)
}

func (k Keeper) AddStudent(ctx sdk.Context, addStudent *types.AddStudentRequest) string {
	students := addStudent.Students
	store := ctx.KVStore(k.storeKey)
	for _, stud := range students {
		marshalAddStudent, err := k.cdc.Marshal(stud)
		handleError(err)
		store.Set(types.StudentStoreKey(stud.Address), marshalAddStudent)
		k.GetStudent(ctx, sdk.AccAddress("lms1").String())
	}
	return "Students Added Successfully"
}

func (k Keeper) GetStudent(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	v := store.Get(types.StudentStoreKey(id))
	fmt.Println(v)
}

func (k Keeper) ApplyLeave(ctx sdk.Context, leave *types.ApplyLeaveRequest) string {
	if _, err := sdk.AccAddressFromBech32(leave.Address); err != nil {
		log.Fatal(err)
		return fmt.Sprint(err)
	}
	store := ctx.KVStore(k.storeKey)

	if store.Get(types.StudentStoreKey(leave.Address)) == nil {
		return "Student does not exist" //
	}

	val, err := k.cdc.Marshal(leave)
	if err != nil {
		return fmt.Sprint(err)
	}
	leaveid := store.Get(types.LeaveCounterStoreKey(leave.Address))
	leaveId, _ := strconv.Atoi(string(leaveid))

	if leaveid == nil {
		leaveId = 0
	}
	leaveId++
	store.Set(types.LeaveCounterStoreKey(leave.Address), []byte(strconv.Itoa(leaveId)))
	store.Set(types.LeaveStoreKey(leave.Address, leaveId), val)
	return "Leave Applied Successfully"
}

func (k Keeper) AcceptLeave(ctx sdk.Context, acceptLeave *types.AcceptLeaveRequest) string {
	store := ctx.KVStore(k.storeKey)
	marshalAcceptLeave, err := k.cdc.Marshal(acceptLeave)
	handleError(err)
	store.Set(types.AcceptLeaveStoreKey(acceptLeave.Admin, acceptLeave.LeaveId), marshalAcceptLeave)
	return "Leave Status Updated"
}

func (k Keeper) CheckLeaveStatus(ctx sdk.Context, studentAddress string) {

}
