package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/cosmosquad-labs/squad/x/auction/types"
)

// NewTxCmd returns the transaction commands for this module
// auction ModuleClient is slightly different from other ModuleClients in that
// it contains a slice of "auction" child commands. These commands are respective
// to auction type handlers that are implemented in other modules but are mounted
// under the auction CLI (eg. liquidfarming and boost modules specific auction).
func NewTxCmd(propCmds []*cobra.Command) *cobra.Command {
	auctionTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Auction transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmdSubmitProp := NewCmdSubmitProposal()
	for _, propCmd := range propCmds {
		flags.AddTxFlagsToCmd(propCmd)
		cmdSubmitProp.AddCommand(propCmd)
	}

	auctionTxCmd.AddCommand(
		// TODO: not implemented yet (e.g: place bid)
		cmdSubmitProp,
	)

	return auctionTxCmd
}

func NewCmdSubmitProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-auction",
		Short: "Create an auction",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create an auction.

Example:
$ %s tx auction create-auction --auction="path/to/auction.json" --from mykey

Where auction.json contains:

{
  "start_price": "1.000000000000000000",
  "selling_coins": [
	  {
		  "denom": "denom1",
		  "amount": "1000000000000"
	  }
  ],
  "paying_coin_denom": "denom2",
  "start_time": "2022-05-01T00:00:00Z",
  "end_time": "2022-06-01T00:00:00Z"
}
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			auction, err := parseCreateAuctionFlags(cmd.Flags())
			if err != nil {
				return fmt.Errorf("failed to parse auction: %w", err)
			}

			auctioneer := clientCtx.GetFromAddress()
			startPrice := auction.StartPrice
			sellingCoins, err := sdk.ParseCoinsNormalized(auction.SellingCoins.String())
			if err != nil {
				return err
			}
			payingCoinDenom := auction.PayingCoinDenom
			startTime := auction.StartTime
			endTime := auction.EndTime
			status := types.AuctionStatusStarted // test

			custom := types.CustomFromAuctionType(
				auctioneer,
				startPrice,
				sellingCoins,
				payingCoinDenom,
				startTime,
				endTime,
				status,
				types.AuctionTypeFixedPrice,
			)

			msg, err := types.NewMsgCreateAuction(custom, auctioneer.String())
			if err != nil {
				return fmt.Errorf("invalid message: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
