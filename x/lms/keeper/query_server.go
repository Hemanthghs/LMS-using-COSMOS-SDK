package keeper

import (
	"context"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
)

var _ types.QueryServer = Keeper{}

// type queryServer struct {
// 	Keeper
// 	types.UnimplementedQueryServer
// }

// func NewQueryServerImpl(k Keeper) types.QueryServer {
// 	return &queryServer{
// 		Keeper: k,
// 	}
// }

func (k Keeper) GetStudentsQuery(context.Context, *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	return &types.GetStudentsResponse{}, nil
}

func (k Keeper) GetLeaveRequestsQuery(context.Context, *types.GetLeaveRequestsRequest) (*types.GetLeaveRequestsResponse, error) {
	return &types.GetLeaveRequestsResponse{}, nil
}

func (k Keeper) GetLeaveApprovedRequestsQuery(context.Context, *types.GetLeaveApprovedRequestsRequest) (*types.GetLeaveApprovedRequestsResponse, error) {
	return &types.GetLeaveApprovedRequestsResponse{}, nil
}

func (k Keeper) GetLeaveStatusQuery(context.Context, *types.GetLeaveStatusRequest) (*types.GetLeaveStatusResponse, error) {
	return &types.GetLeaveStatusResponse{}, nil
}
