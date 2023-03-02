package cli

import (
	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func NewQueryCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "LMS",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 4,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GetStudentsCmd(),
		GetLeaveRequestsCmd(),
		GetLeaveApprovedRequestsCmd(),
		GetLeaveStatusCmd(),
	)
	return txCmd
}

func GetStudentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getstudents",
		Short: "To get the list of all students",
		Long:  "To get teh list of all students",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			getStudentsRequest := &types.GetStudentsRequest{}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetStudentsQuery(cmd.Context(), getStudentsRequest)
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetLeaveRequestsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getleaves",
		Short: "To get the list of all the leave requets",
		Long:  "To get the list of all the leave requests",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			getLeavesRequest := &types.GetLeaveRequestsRequest{}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetLeaveRequestsQuery(cmd.Context(), getLeavesRequest)
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetLeaveStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getstatus [LeaveId]",
		Short: "To get the leave status",
		Long:  "To get the leave status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			getLeaveStatusRequest := &types.GetLeaveStatusRequest{
				LeaveID: args[0],
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetLeaveStatusQuery(cmd.Context(), getLeaveStatusRequest)
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetLeaveApprovedRequestsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getapproved",
		Short: "To get the list of approved leaves",
		Long:  "To get the list of approved leaves",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			getLeavesRequest := &types.GetLeaveApprovedRequestsRequest{}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetLeaveApprovedRequestsQuery(cmd.Context(), getLeavesRequest)
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
