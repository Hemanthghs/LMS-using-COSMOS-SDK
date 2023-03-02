package keeper

import (
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

func (k Keeper) RegisterAdmin(ctx sdk.Context, registerAdmin *types.RegisterAdminRequest) error {

	if registerAdmin.Name == "" {
		return types.ErrAdminNameNil
	} else if registerAdmin.Address == "" {
		return types.ErrAdminAddressNil
	} else {
		store := ctx.KVStore(k.storeKey)
		// panic("hello")
		marshalRegisterAdmin, err := k.cdc.Marshal(registerAdmin)
		handleError(err)
		store.Set(types.AdminStoreKey(registerAdmin.Address), marshalRegisterAdmin)
		return nil
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

func (k Keeper) ApplyLeave(ctx sdk.Context, applyLeave *types.ApplyLeaveRequest) string {
	store := ctx.KVStore(k.storeKey)
	marshalApplyLeave, err := k.cdc.Marshal(applyLeave)
	handleError(err)
	addr := types.StudentLeavesCounterKey(applyLeave.Address)
	// addr := types.StudentLeavesCounterKey(sdk.AccAddress(string(applyLeave.Address)).String())
	counter := store.Get(addr)
	if counter == nil {
		store.Set(addr, []byte("1"))
	} else {
		c, err := strconv.Atoi(string(counter))
		handleError(err)
		c = c + 1
		store.Set(addr, []byte(fmt.Sprint(c)))
	}
	counter = store.Get(addr)
	handleError(err)
	store.Set(types.LeaveStoreKey(applyLeave.Address, string(counter)), marshalApplyLeave)
	// id := types.LeaveStoreKey(applyLeave.Address, string(counter))
	return "Leave Applied Successfully"
}

func (k Keeper) AcceptLeave(ctx sdk.Context, acceptLeave *types.AcceptLeaveRequest) string {
	store := ctx.KVStore(k.storeKey)
	marshalAcceptLeave, err := k.cdc.Marshal(acceptLeave)
	handleError(err)
	leaveid := store.Get(types.StudentLeavesCounterKey(acceptLeave.LeaveId))
	store.Set(types.AcceptLeaveStoreKey(string(leaveid)), marshalAcceptLeave)
	return "Leave Status Updated"
}

func (k Keeper) GetLeaveStatus(ctx sdk.Context, getLeaveStatus *types.GetLeaveStatusRequest) string {
	store := ctx.KVStore(k.storeKey)
	leaveid := store.Get(types.StudentLeavesCounterKey(getLeaveStatus.LeaveID))
	res := store.Get(types.AcceptLeaveStoreKey(string(leaveid)))
	var leave types.AcceptLeaveRequest
	k.cdc.Unmarshal(res, &leave)
	fmt.Println(leave)
	status := leave.Status.String()
	return status
}

func (k Keeper) GetStudents(ctx sdk.Context, getStudents *types.GetStudentsRequest) []*types.Student {
	store := ctx.KVStore(k.storeKey)

	var students []*types.Student
	itr := store.Iterator(types.StudentKey, nil)
	for ; itr.Valid(); itr.Next() {
		var t types.Student
		k.cdc.Unmarshal(itr.Value(), &t)
		students = append(students, &t)
	}
	return students
}

func (k Keeper) GetLeaveRequests(ctx sdk.Context, getLeaveRequests *types.GetLeaveRequestsRequest) []types.ApplyLeaveRequest {
	store := ctx.KVStore(k.storeKey)
	var t types.ApplyLeaveRequest
	var leaves []types.ApplyLeaveRequest
	itr := store.Iterator(types.LeaveKey, nil)
	for ; itr.Valid(); itr.Next() {
		k.cdc.Unmarshal(itr.Value(), &t)
		leaves = append(leaves, t)
	}
	return leaves
}

func (k Keeper) GetLeaveApprovedRequests(ctx sdk.Context, getLeaveApprovedRequests *types.GetLeaveApprovedRequestsRequest) []types.ApplyLeaveRequest {
	store := ctx.KVStore(k.storeKey)
	var t types.ApplyLeaveRequest
	var approvedleaves []types.ApplyLeaveRequest
	itr := store.Iterator(types.AcceptLeaveKey, nil)
	for ; itr.Valid(); itr.Next() {
		k.cdc.Unmarshal(itr.Value(), &t)
		approvedleaves = append(approvedleaves, t)
	}
	return approvedleaves
}
