package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
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
		NewFarmCmd(),
		NewCancelQueuedFarmingCmd(),
		NewUnfarmCmd(),
		NewPlaceBidCmd(),
	)

	return cmd
}

// NewFarmCmd implements the farm command handler.
func NewFarmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "farm [pool-id] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "Farm pool coin for liquid farming",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Farm pool coin for liquid farming. 
			
Example:
$ %s tx %s farm 1 10000000pool1 --from mykey
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse pool id: %w", err)
			}

			farmingCoin, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return fmt.Errorf("invalid farming coin: %w", err)
			}

			msg := types.NewMsgFarm(
				poolId,
				clientCtx.GetFromAddress().String(),
				farmingCoin,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCancelQueuedFarmingCmd implements the cancel queued farming command handler.
func NewCancelQueuedFarmingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-queued-farming [pool-id] [deposit-request-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Cancel queued farming",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Cancel queued farming with the given pool id and deposit request id.
You can't cancel the queued farming that is already executed to mint LFCoin.
			
Example:
$ %s tx %s cancel-queued-farming 1 1 --from mykey
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse pool id: %w", err)
			}

			farmingCoin, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return fmt.Errorf("invalid farming coin: %w", err)
			}

			msg := types.NewMsgCancelQueuedFarming(
				poolId,
				clientCtx.GetFromAddress().String(),
				farmingCoin,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewUnfarmCmd implements the unfarm command handler.
func NewUnfarmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unfarm [pool-id] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "Unfarm liquid farming coin",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Unfarm liquid farming coin to receive the corresponding pool coin.
			
Example:
$ %s tx %s unfarm 1 100000lf1 --from mykey
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse pool id: %w", err)
			}

			unfarmingCoin, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return fmt.Errorf("invalid unfarming coin: %w", err)
			}

			msg := types.NewMsgUnfarm(
				poolId,
				clientCtx.GetFromAddress().String(),
				unfarmingCoin,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewPlaceBidCmd implements the place bid command handler.
func NewPlaceBidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "place-bid [pool-id] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "Place a bid for a rewards auction",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Place a bid for a rewards auction.
			
Example:
$ %s tx %s place-bid 1 100000lf1 --from mykey
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse pool id: %w", err)
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return fmt.Errorf("invalid bidding amount: %w", err)
			}

			msg := types.NewMsgPlaceBid(
				poolId,
				clientCtx.GetFromAddress().String(),
				amount,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
