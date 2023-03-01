package cli

import (
	"log"
	"time"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func handleError(err error) {
	log.Fatal(err)
}

func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "LMS",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 4,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		NewRegisterAdminCmd(),
		NewApplyLeaveReqCmd(),
		NewAddStudentRequestCmd(),
		NewAcceptLeaveReqCmd(),
	)
	return txCmd
}

func NewRegisterAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registeradmin [name] [address]",
		Short: "To register new admin",
		Long:  "To register new admin",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			handleError(err)
			name := args[0]
			address := sdk.AccAddress(args[1])
			msg := types.NewRegisterAdminReq(address, name)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewAddStudentRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "addstudents",
		Short: "This is used to add new students",
		Long:  "This is used to add new students",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			adminaddress := args[0]
			students := []*types.Student{}
			for i := 0; i < (len(args)-1)/3; i++ {
				student := &types.Student{
					Name:    args[3*i+1],
					Id:      args[3*i+2],
					Address: args[3*i+3],
				}
				students = append(students, student)
			}
			msg := types.NewAddStudentReq(sdk.AccAddress(adminaddress), students)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewApplyLeaveReqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "applyleave [Address] [Reason] [from] [to]",
		Short: "To apply for a leave",
		Long:  "To apply for a leave",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			handleError(err)
			Address := args[0]
			Reason := args[1]
			format := "2006-Jan-06"
			from, _ := time.Parse(format, args[2])
			to, _ := time.Parse(format, args[3])
			msg := types.NewApplyLeaveReq(sdk.AccAddress(Address), Reason, &from, &to)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewAcceptLeaveReqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "acceptleave [Admin] [LeaveId] [Status]",
		Short: "To accept a leave request",
		Long:  "To accept a leave request",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCxt, err := client.GetClientTxContext(cmd)
			handleError(err)
			Admin := args[0]
			LeaveId := args[1]
			Status := args[2]
			msg := types.NewAcceptLeaveReq(sdk.AccAddress(Admin), LeaveId, Status)
			return tx.GenerateOrBroadcastTxCLI(clientCxt, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
