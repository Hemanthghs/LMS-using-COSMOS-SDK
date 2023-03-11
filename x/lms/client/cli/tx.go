package cli

import (
	"time"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

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
		Use:     "registeradmin [name] [address]",
		Short:   "To register new admin",
		Long:    "To register new admin",
		Example: "./lmsd tx leave registeradmin admin1 cosmos1111 --from validator-key --chain-id testnet",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}

			name := args[0]
			address, _ := sdk.AccAddressFromBech32(args[1])

			msg := types.NewRegisterAdminReq(address, name)
			// panic("called1")
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewAddStudentRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "addstudents",
		Short:   "This is used to add new students",
		Long:    "This is used to add new students",
		Example: "./lmsd tx leave addstudents student1 1 cosmos1231212 --from validator-key --chain-id testnet",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			fromadd := clientCtx.GetFromAddress()
			if err != nil {
				return err
			}
			// adminaddress, _ := sdk.AccAddressFromBech32(args[0])
			adminaddress := fromadd
			students := []*types.Student{}
			for i := 0; i < (len(args))/3; i++ {
				student := &types.Student{
					Name:    args[3*i],
					Id:      args[3*i+1],
					Address: args[3*i+2],
				}
				students = append(students, student)
			}
			msg := types.NewAddStudentReq(adminaddress, students)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewApplyLeaveReqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "applyleave [Address] [Reason] [from] [to]",
		Short:   "To apply for a leave",
		Long:    "To apply for a leave",
		Example: "./lmsd tx leave applyleave cosmos12123 Fever 2023-Mar-01 2023-Mar-03 --from validator-key --chain-id testnet",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			fromadd := clientCtx.GetFromAddress()
			if err != nil {
				panic(err)
			}
			Signer := fromadd
			Address, _ := sdk.AccAddressFromBech32(args[0])
			Reason := args[1]
			format := "2006-Jan-02"
			from, _ := time.Parse(format, args[2])
			to, _ := time.Parse(format, args[3])
			msg := types.NewApplyLeaveReq(Signer, Address, Reason, &from, &to)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewAcceptLeaveReqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "acceptleave [LeaveId] [Status]",
		Short:   "To accept a leave request",
		Long:    "To accept a leave request",
		Example: "./lmsd tx leave acceptleave cosmos1231212 1 --from validator-key --chain-id testnet",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			fromadd := clientCtx.GetFromAddress()

			if err != nil {
				panic(err)
			}
			// Admin := args[0]
			Admin := fromadd
			LeaveId := args[0]
			Status := args[1]
			msg := types.NewAcceptLeaveReq(Admin, LeaveId, Status)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
