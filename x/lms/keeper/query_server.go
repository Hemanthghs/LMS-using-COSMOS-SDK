package keeper

import (
	"context"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper
	types.UnimplementedQueryServer
}

func NewQueryServerImpl(k Keeper) types.QueryServer {
	return &queryServer{
		Keeper: k,
	}
}

func (k queryServer) GetStudents(context.Context, *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	return &types.GetStudentsResponse{}, nil
}

func (k queryServer) GetLeaveRequests(context.Context, *types.GetLeaveRequestsRequest) (*types.GetLeaveRequestsResponse, error) {
	return &types.GetLeaveRequestsResponse{}, nil
}

func (k queryServer) GetLeaveApprovedRequests(context.Context, *types.GetLeaveApprovedRequestsRequest) (*types.GetLeaveApprovedRequestsResponse, error) {
	return &types.GetLeaveApprovedRequestsResponse{}, nil
}
