package keeper

import (
	"context"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
	types.UnimplementedMsgServer
}

func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{
		Keeper: k,
	}
}

func (k msgServer) AddStudent(ctx context.Context, students *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.Keeper.AddStudent(sdkCtx, students)
	return &types.AddStudentResponse{}, nil
}

func (k msgServer) ApplyLeave(context.Context, *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	return &types.ApplyLeaveResponse{}, nil
}

func (k msgServer) RegisterAdmin(context.Context, *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	return &types.RegisterAdminResponse{}, nil
}

func (k msgServer) AcceptLeave(context.Context, *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	return &types.AcceptLeaveResponse{}, nil
}
