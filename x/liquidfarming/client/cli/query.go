package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

			resp, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
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
		Use:   "liquidfarms",
		Args:  cobra.NoArgs,
		Short: "Query for all liquidfarms",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for all liquidfarms on a network.

Example:
$ %s query %s liquidfarms
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.LiquidFarms(cmd.Context(), &types.QueryLiquidFarmsRequest{
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryLiquidFarmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidfarm [pool-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the specific liquidfarm",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the specific liquidfarm on a network.

Example:
$ %s query %s liquidfarm 1
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			poolId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse pool id: %w", err)
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.LiquidFarm(cmd.Context(), &types.QueryLiquidFarmRequest{
				PoolId: poolId,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryQueuedFarmingsCmd() *cobra.Command {
	bech32PrefixAccAddr := sdk.GetConfig().GetBech32AccountAddrPrefix()

	cmd := &cobra.Command{
		Use:   "queued-farmings [pool-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query all queued farmings for the liquidfarm",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all queued farmings for the liquidfarm on a network.

Example:
$ %s query %s queued-farmings
$ %s query %s queued-farmings --farmer %s1zaavvzxez0elundtn32qnk9lkm8kmcszzsv80v
`,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName, bech32PrefixAccAddr,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			poolId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse pool id: %w", err)
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var res proto.Message
			queryClient := types.NewQueryClient(clientCtx)
			farmerAddr, _ := cmd.Flags().GetString(FlagFarmer)
			if farmerAddr == "" {
				res, err = queryClient.QueuedFarmings(cmd.Context(), &types.QueryQueuedFarmingsRequest{
					PoolId:     poolId,
					Pagination: pageReq,
				})
			} else {
				res, err = queryClient.QueuedFarmingsByFarmer(cmd.Context(), &types.QueryQueuedFarmingsByFarmerRequest{
					PoolId:        poolId,
					FarmerAddress: farmerAddr,
					Pagination:    pageReq,
				})
			}
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().AddFlagSet(flagSetQueuedFarmings())
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryRewardsAuctionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rewards-auctions",
		Args:  cobra.NoArgs,
		Short: "Query all rewards auctions for the liquidfarm",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all rewards auctions for the liquidfarm on a network.

Example:
$ %s query %s rewards-auctions
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			poolId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse pool id: %w", err)
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.RewardsAuctions(cmd.Context(), &types.QueryRewardsAuctionsRequest{
				PoolId:     poolId,
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryRewardsAuctionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reward-auction [pool-id] [auction-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Query the specific reward auction",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the specific reward auction on a network.

Example:
$ %s query %s reward-auction 1 1
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			poolId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse pool id: %w", err)
			}

			auctionId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to auction pool id: %w", err)
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.RewardsAuction(cmd.Context(), &types.QueryRewardsAuctionRequest{
				PoolId:    poolId,
				AuctionId: auctionId,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryBidsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bids [pool-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query all bids for the rewards auction",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all bids for the rewards auction on a network.

Example:
$ %s query %s bids
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			poolId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse pool id: %w", err)
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Bids(cmd.Context(), &types.QueryBidsRequest{
				PoolId:     poolId,
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
