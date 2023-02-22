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

func (k msgServer) AddStudent(ctx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.Keeper.AddStudent(sdkCtx, req)
	return &types.AddStudentResponse{}, nil
}

func (k msgServer) ApplyLeave(ctx context.Context, req *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.Keeper.ApplyLeave(sdkCtx, req)
	return &types.ApplyLeaveResponse{}, nil
}

func (k msgServer) RegisterAdmin(ctx context.Context, req *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.Keeper.RegisterAdmin(sdkCtx, req)
	return &types.RegisterAdminResponse{}, nil
}

func (k msgServer) AcceptLeave(ctx context.Context, req *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	// sdkCtx := sdk.UnwrapSDKContext(ctx)
	// k.Keeper.AcceptLeave(sdkCtx, req)
	return &types.AcceptLeaveResponse{}, nil
}
