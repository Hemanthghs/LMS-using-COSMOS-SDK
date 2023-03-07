package keeper

import (
	"context"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) GetStudentsQuery(goCtx context.Context, req *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	res := k.GetStudents(ctx, req)
	// panic(res)
	return &types.GetStudentsResponse{
		Students: res,
	}, nil
}

func (k Keeper) GetLeaveRequestsQuery(goCtx context.Context, req *types.GetLeaveRequestsRequest) (*types.GetLeaveRequestsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	res := k.GetLeaveRequests(ctx, req)
	return &types.GetLeaveRequestsResponse{
		Leaverequests: res,
	}, nil
}

func (k Keeper) GetLeaveApprovedRequestsQuery(goCtx context.Context, req *types.GetLeaveApprovedRequestsRequest) (*types.GetLeaveApprovedRequestsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	res := k.GetLeaveApprovedRequests(ctx, req)
	return &types.GetLeaveApprovedRequestsResponse{
		Leaverequests: res,
	}, nil
}

func (k Keeper) GetLeaveStatusQuery(goCtx context.Context, req *types.GetLeaveStatusRequest) (*types.GetLeaveStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	res := k.GetLeaveStatus(ctx, req)
	return &types.GetLeaveStatusResponse{
		Leaverequest: res,
	}, nil
}

func (k Keeper) GetAdmins(goCtx context.Context, req *types.GetAdminsRequest) (*types.GetAdminsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	res := k.GetAdminsRequest(ctx, req)
	return &types.GetAdminsResponse{
		Admins: res,
	}, nil
}
