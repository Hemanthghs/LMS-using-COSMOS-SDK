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
		marshalRegisterAdmin, err := k.cdc.Marshal(registerAdmin)
		handleError(err)
		store.Set(types.AdminStoreKey(registerAdmin.Address), marshalRegisterAdmin)
		return nil
	}
}

func (k Keeper) GetAdmin(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	v := store.Get(types.AdminStoreKey(id))
	fmt.Println(v)
}

func (k Keeper) AddStudent(ctx sdk.Context, addStudent *types.AddStudentRequest) error {
	students := addStudent.Students
	store := ctx.KVStore(k.storeKey)
	admin := store.Get(types.AdminStoreKey(addStudent.Admin))
	if admin == nil {
		return types.ErrAdminNotRegistered
	}
	for _, stud := range students {
		marshalAddStudent, err := k.cdc.Marshal(stud)
		handleError(err)
		store.Set(types.StudentStoreKey(stud.Address), marshalAddStudent)
		k.GetStudent(ctx, sdk.AccAddress("lms1").String())
	}
	return nil
}

func (k Keeper) GetStudent(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	v := store.Get(types.StudentStoreKey(id))
	fmt.Println(v)
}

func (k Keeper) ApplyLeave(ctx sdk.Context, applyLeave *types.ApplyLeaveRequest) error {
	store := ctx.KVStore(k.storeKey)
	if store.Get(types.StudentStoreKey(applyLeave.Address)) == nil {
		return types.ErrStudentNotFound
	}
	applyLeave.Status = types.LeaveStatus_STATUS_UNDEFINED
	marshalApplyLeave, err := k.cdc.Marshal(applyLeave)
	handleError(err)
	addr := types.StudentLeavesCounterKey(applyLeave.Address)
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
	return nil
}

func (k Keeper) AcceptLeave(ctx sdk.Context, acceptLeave *types.AcceptLeaveRequest) error {
	store := ctx.KVStore(k.storeKey)
	stud := store.Get(types.StudentLeavesCounterKey(acceptLeave.LeaveId))
	if stud == nil {
		return types.ErrLeaveNotFound
	}
	marshalAcceptLeave, err := k.cdc.Marshal(acceptLeave)
	handleError(err)
	leaveCount := store.Get(types.StudentLeavesCounterKey(acceptLeave.LeaveId))
	leaveid := store.Get(types.LeaveStoreKey(acceptLeave.LeaveId, string(leaveCount)))
	store.Set(types.AcceptLeaveStoreKey(string(leaveid)), marshalAcceptLeave)
	c := store.Get(types.StudentLeavesCounterKey(acceptLeave.LeaveId))
	count := string(c)
	l := store.Get(types.LeaveStoreKey(acceptLeave.LeaveId, count))
	var leave types.ApplyLeaveRequest
	k.cdc.Unmarshal(l, &leave)
	leave.Status = acceptLeave.Status
	leaveMarhsal, _ := k.cdc.Marshal(&leave)
	store.Set(types.LeaveStoreKey(acceptLeave.LeaveId, count), leaveMarhsal)
	return nil
}

func (k Keeper) GetLeaveStatus(ctx sdk.Context, getLeaveStatus *types.GetLeaveStatusRequest) (*types.ApplyLeaveRequest, error) {
	store := ctx.KVStore(k.storeKey)
	leaveCount := store.Get(types.StudentLeavesCounterKey(getLeaveStatus.LeaveID))

	res := store.Get(types.LeaveStoreKey(getLeaveStatus.LeaveID, string(leaveCount)))
	var leave types.ApplyLeaveRequest
	k.cdc.Unmarshal(res, &leave)
	return &leave, nil
}

func (k Keeper) GetStudents(ctx sdk.Context, getStudents *types.GetStudentsRequest) []*types.Student {
	store := ctx.KVStore(k.storeKey)

	var students []*types.Student
	itr := sdk.KVStorePrefixIterator(store, types.StudentKey)
	for ; itr.Valid(); itr.Next() {
		var t types.Student
		k.cdc.Unmarshal(itr.Value(), &t)
		students = append(students, &t)
	}
	return students
}

func (k Keeper) GetLeaveRequests(ctx sdk.Context, getLeaveRequests *types.GetLeaveRequestsRequest) []*types.ApplyLeaveRequest {
	store := ctx.KVStore(k.storeKey)
	var leaves []*types.ApplyLeaveRequest
	itr := sdk.KVStorePrefixIterator(store, types.LeaveKey)
	for ; itr.Valid(); itr.Next() {
		var t types.ApplyLeaveRequest
		k.cdc.Unmarshal(itr.Value(), &t)
		leaves = append(leaves, &t)
	}
	// panic(leaves)
	return leaves
}

func (k Keeper) GetLeaveApprovedRequests(ctx sdk.Context, getLeaveApprovedRequests *types.GetLeaveApprovedRequestsRequest) []*types.AcceptLeaveRequest {
	store := ctx.KVStore(k.storeKey)
	var approvedleaves []*types.AcceptLeaveRequest
	itr := sdk.KVStorePrefixIterator(store, types.AcceptLeaveKey)
	for ; itr.Valid(); itr.Next() {
		var t types.AcceptLeaveRequest
		k.cdc.Unmarshal(itr.Value(), &t)
		approvedleaves = append(approvedleaves, &t)
	}
	return approvedleaves
}

func (k Keeper) GetAdminsRequest(ctx sdk.Context, getLeaveRequests *types.GetAdminsRequest) []*types.RegisterAdminRequest {
	store := ctx.KVStore(k.storeKey)
	var admins []*types.RegisterAdminRequest
	itr := sdk.KVStorePrefixIterator(store, types.AdminKey)
	for ; itr.Valid(); itr.Next() {
		var t types.RegisterAdminRequest
		k.cdc.Unmarshal(itr.Value(), &t)
		admins = append(admins, &t)
	}
	return admins
}
