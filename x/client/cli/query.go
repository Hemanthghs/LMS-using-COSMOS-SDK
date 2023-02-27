package main

import (
	"fmt"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

var getStudents = &cobra.Command{
	Use:   "getstudents",
	Short: "To get the details of all students",
	Long:  "To get the details of all students",

	Run: func(cmd *cobra.Command, args []string) {
		clientctx, err := client.GetClientQueryContext(cmd)
		handleError(err)
		queryClient := types.NewQueryClient(clientctx)
		res, err := queryClient.GetStudents(cmd.Context(), &types.GetStudentsRequest{})
		handleError(err)
		fmt.Println(res)
	},
}

var getLeaveRequests = &cobra.Command{
	Use:   "getleaves",
	Short: "To get all the leave requests",
	Long:  "To get all the leave requests",

	Run: func(cmd *cobra.Command, args []string) {
		clientctx, err := client.GetClientQueryContext(cmd)
		handleError(err)
		queryClient := types.NewQueryClient(clientctx)
		res, err := queryClient.GetLeaveRequests(cmd.Context(), &types.GetLeaveRequestsRequest{})
		handleError(err)
		fmt.Println(res)
	},
}

var getLeaveApprovedRequests = &cobra.Command{
	Use:   "getapproved",
	Short: "To get all the approved leaves",
	Long:  "To get all the approved leaves",

	Run: func(cmd *cobra.Command, args []string) {
		clientctx, err := client.GetClientQueryContext(cmd)
		handleError(err)
		queryClient := types.NewQueryClient(clientctx)
		res, err := queryClient.GetLeaveApprovedRequests(cmd.Context(), &types.GetLeaveApprovedRequestsRequest{})
		handleError(err)
		fmt.Println(res)
	},
}

var getLeaveStatus = &cobra.Command{
	Use:   "getstatus",
	Short: "To get the leave status",
	Long:  "To get the leave status",

	Run: func(cmd *cobra.Command, args []string) {
		leaveid, err := cmd.Flags().GetString("leaveid")
		handleError(err)
		clientctx, err := client.GetClientQueryContext(cmd)
		handleError(err)
		queryClient := types.NewQueryClient(clientctx)
		res, err := queryClient.GetLeaveStatus(cmd.Context(), &types.GetLeaveStatusRequest{
			LeaveID: leaveid,
		})
		handleError(err)
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(getStudents)
	rootCmd.AddCommand(getLeaveRequests)
	rootCmd.AddCommand(getLeaveApprovedRequests)
	rootCmd.AddCommand(getLeaveStatus)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getuserCmd.PersistentFlags().String("email", "", "User email")
	getLeaveStatus.PersistentFlags().String("leaveid", "", "Leave ID")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
