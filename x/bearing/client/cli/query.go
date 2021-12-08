package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/tendermint/farming/x/bearing/types"
)

// GetQueryCmd returns a root CLI command handler for all x/bearing query commands.
func GetQueryCmd() *cobra.Command {
	bearingQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the bearing module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	bearingQueryCmd.AddCommand(
		GetCmdQueryParams(),
		GetCmdQueryBearings(),
		GetCmdQueryAddress(),
	)

	return bearingQueryCmd
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the values set as bearing parameters",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as bearing parameters.

Example:
$ %s query %s params
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Params(
				context.Background(),
				&types.QueryParamsRequest{},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryBearings implements the query bearings command.
func GetCmdQueryBearings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bearings",
		Args:  cobra.NoArgs,
		Short: "Query all bearings",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Queries all bearings along with their metadata.

Example:
$ %s query %s bearings
$ %s query %s bearings --name ...
$ %s query %s bearings --source-address %s1zaavvzxez0elundtn32qnk9lkm8kmcszzsv80v
$ %s query %s bearings --destination-address %s1zaavvzxez0elundtn32qnk9lkm8kmcszzsv80v
`,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName, sdk.Bech32MainPrefix,
				version.AppName, types.ModuleName, sdk.Bech32MainPrefix,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name, _ := cmd.Flags().GetString(FlagName)
			sourceAddr, _ := cmd.Flags().GetString(FlagSourceAddress)
			destinationAddr, _ := cmd.Flags().GetString(FlagDestinationAddress)

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Bearings(
				context.Background(),
				&types.QueryBearingsRequest{
					Name:               name,
					SourceAddress:      sourceAddr,
					DestinationAddress: destinationAddr,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().AddFlagSet(flagSetBearings())
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryAddress implements the query an address that can be used as source and destination is derived according to the given type, module name, and name command.
func GetCmdQueryAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address [name]",
		Args:  cobra.ExactArgs(1),
		Short: "Query an address that can be used as source or destination address",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query an address that can be used as source or destination address. It is derived with address derivation name, module name, and address type.

Example:
$ %s query %s address testSourceAddr
$ %s query %s address fee_collector --type 1
$ %s query %s address GravityDEXFarmingBearing --module-name farming

Default flag:
$ [--type 0] - ADDRESS_TYPE_32_BYTES of ADR 028
$ [--module-name %s] - When type is 0, the default module name is %s
`,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName,
				types.ModuleName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			moduleName, _ := cmd.Flags().GetString(FlagModuleName)
			addressTypeStr, _ := cmd.Flags().GetString(FlagType)
			addressType, err := strconv.Atoi(addressTypeStr)
			if err != nil {
				addressType = 0
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Addresses(
				context.Background(),
				&types.QueryAddressesRequest{
					Type:       types.AddressType(addressType),
					ModuleName: moduleName,
					Name:       args[0],
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().AddFlagSet(flagSetAddress())
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
