package keeper

import (
	"context"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

type msgServer struct {
	Keeper
	types.UnimplementedMsgServer
}

// func NewMsgServerImpl(k Keeper) types.MsgServer {
// 	return &msgServer{
// 		Keeper: k,
// 	}
// }

func (k Keeper) AddStudents(ctx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.AddStudent(sdkCtx, req)
	return &types.AddStudentResponse{}, nil
}

func (k Keeper) LeaveApply(ctx context.Context, req *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.ApplyLeave(sdkCtx, req)
	return &types.ApplyLeaveResponse{}, nil
}

func (k Keeper) AdminRegister(ctx context.Context, req *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.RegisterAdmin(sdkCtx, req)
	return &types.RegisterAdminResponse{}, nil
}

func (k Keeper) LeaveAccept(ctx context.Context, req *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	err := k.AcceptLeave(sdkCtx, req)
	return &types.AcceptLeaveResponse{}, err
}
