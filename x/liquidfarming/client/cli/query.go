package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// GetQueryCmd returns the cli query commands for the module
func GetQueryCmd(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewQueryParamsCmd(),
		NewQueryLiquidFarmsCmd(),
		NewQueryLiquidFarmCmd(),
		NewQueryQueuedFarmingsCmd(),
		NewQueryQueuedFarmingCmd(),
		NewQueryRewardsAuctionsCmd(),
		NewQueryRewardsAuctionCmd(),
		NewQueryBidsCmd(),
	)

	return cmd
}

func NewQueryParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current liquidfarming parameters information",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as liquidfarming parameters.
Example:
$ %s query %s params
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			resp, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&resp.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryLiquidFarmsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current liquidfarming parameters information",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as liquidfarming parameters.
Example:
$ %s query %s params
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// clientCtx, err := client.GetClientQueryContext(cmd)
			// if err != nil {
			// 	return err
			// }

			// queryClient := types.NewQueryClient(clientCtx)

			// resp, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			// if err != nil {
			// 	return err
			// }

			// return clientCtx.PrintProto(&resp.Params)
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryLiquidFarmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current liquidfarming parameters information",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as liquidfarming parameters.
Example:
$ %s query %s params
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// clientCtx, err := client.GetClientQueryContext(cmd)
			// if err != nil {
			// 	return err
			// }

			// queryClient := types.NewQueryClient(clientCtx)

			// resp, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			// if err != nil {
			// 	return err
			// }

			// return clientCtx.PrintProto(&resp.Params)
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryQueuedFarmingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current liquidfarming parameters information",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as liquidfarming parameters.
Example:
$ %s query %s params
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// clientCtx, err := client.GetClientQueryContext(cmd)
			// if err != nil {
			// 	return err
			// }

			// queryClient := types.NewQueryClient(clientCtx)

			// resp, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			// if err != nil {
			// 	return err
			// }

			// return clientCtx.PrintProto(&resp.Params)
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryQueuedFarmingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current liquidfarming parameters information",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as liquidfarming parameters.
Example:
$ %s query %s params
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// clientCtx, err := client.GetClientQueryContext(cmd)
			// if err != nil {
			// 	return err
			// }

			// queryClient := types.NewQueryClient(clientCtx)

			// resp, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			// if err != nil {
			// 	return err
			// }

			// return clientCtx.PrintProto(&resp.Params)
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryRewardsAuctionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current liquidfarming parameters information",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as liquidfarming parameters.
Example:
$ %s query %s params
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// clientCtx, err := client.GetClientQueryContext(cmd)
			// if err != nil {
			// 	return err
			// }

			// queryClient := types.NewQueryClient(clientCtx)

			// resp, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			// if err != nil {
			// 	return err
			// }

			// return clientCtx.PrintProto(&resp.Params)
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryRewardsAuctionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current liquidfarming parameters information",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as liquidfarming parameters.
Example:
$ %s query %s params
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// clientCtx, err := client.GetClientQueryContext(cmd)
			// if err != nil {
			// 	return err
			// }

			// queryClient := types.NewQueryClient(clientCtx)

			// resp, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			// if err != nil {
			// 	return err
			// }

			// return clientCtx.PrintProto(&resp.Params)
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryBidsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current liquidfarming parameters information",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as liquidfarming parameters.
Example:
$ %s query %s params
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// clientCtx, err := client.GetClientQueryContext(cmd)
			// if err != nil {
			// 	return err
			// }

			// queryClient := types.NewQueryClient(clientCtx)

			// resp, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			// if err != nil {
			// 	return err
			// }

			// return clientCtx.PrintProto(&resp.Params)
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
