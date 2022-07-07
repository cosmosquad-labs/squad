package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/cosmosquad-labs/squad/x/liquidfarming/types"
)

// GetTxCmd returns the cli transaction commands for the module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Transaction commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewDepositCmd(),
		NewCancelCmd(),
		NewWithdrawCmd(),
		// TODO: add NewPlaceBidCmd()
	)

	return cmd
}

// NewDepositCmd implements the deposit command handler.
func NewDepositCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [amount]",
		Args:  cobra.ExactArgs(1),
		Short: "deposit coins",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Deposit ... 
			
Example:
$ %s tx %s deposit 1000000poolxxxxxx --from mykey
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// clientCtx, err := client.GetClientTxContext(cmd)
			// if err != nil {
			// 	return err
			// }

			// return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCancelCmd implements the cancel command handler.
func NewCancelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel [id]",
		Args:  cobra.ExactArgs(1),
		Short: "cancel deposit request",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Deposit ... 
			
Example:
$ %s tx %s cancel 1 --from mykey
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// clientCtx, err := client.GetClientTxContext(cmd)
			// if err != nil {
			// 	return err
			// }

			// return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewWithdrawCmd implements the withdraw command handler.
func NewWithdrawCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [amount]",
		Args:  cobra.ExactArgs(1),
		Short: "withdraw liquid farming coin",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Withdraw ... 
			
Example:
$ %s tx %s deposit --from mykey
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// clientCtx, err := client.GetClientTxContext(cmd)
			// if err != nil {
			// 	return err
			// }

			// return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
