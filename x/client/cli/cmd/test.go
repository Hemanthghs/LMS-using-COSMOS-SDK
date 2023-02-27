package cmd

import (
	"fmt"

	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func GetStudentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "qgetstudents",
		Short: "To get the details of all students",
		Long:  "To get the details of all students",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientctx, err := client.GetClientQueryContext(cmd)
			handleError(err)
			queryClient := types.NewQueryClient(clientctx)
			res, err := queryClient.GetStudents(cmd.Context(), &types.GetStudentsRequest{})
			handleError(err)
			fmt.Println(res)
			return nil
		},
	}

	return cmd
}
