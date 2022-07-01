package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	auctiontypes "github.com/cosmosquad-labs/squad/x/auction/types"
	"github.com/cosmosquad-labs/squad/x/liquidfarming/client/utils"
	"github.com/cosmosquad-labs/squad/x/liquidfarming/types"
)

func NewCreateFixedPriceAuctionTxCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create-fixed-price-auction [auction-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Create a fixed price auction",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a fixed price auction

Example:
$ %s tx liquidfarming create-fixed-price-auction <path/to/auction.json> --from=<key_or_address>

Where auction.json contains:

{
  "start_price": "1.000000000000000000",
  "selling_coin": {
	  "denom": "denom1",
	  "amount": "1000000000000"
   },
  "paying_coin_denom": "denom2",
  "start_time": "2022-06-01T00:00:00Z",
  "end_time": "2022-12-01T00:00:00Z"
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

			auction, err := utils.ParseFixedPriceAuctionJSON(clientCtx.LegacyAmino, args[0])
			if err != nil {
				return fmt.Errorf("failed to parse auction: %w", err)
			}

			auctioneer := clientCtx.GetFromAddress().String()
			startPrice := auction.StartPrice
			sellingCoin, err := sdk.ParseCoinNormalized(auction.SellingCoin.String())
			if err != nil {
				return err
			}
			payingCoinDenom := auction.PayingCoinDenom
			startTime := auction.StartTime
			endTime := auction.EndTime
			status := types.AuctionStatusStarted // test

			custom := types.NewFixedPriceAuction(
				auctioneer,
				startPrice,
				sellingCoin,
				sellingCoin,
				payingCoinDenom,
				startTime,
				endTime,
				status,
			)

			msg, err := auctiontypes.NewMsgCreateAuction(custom, auctioneer)
			if err != nil {
				return fmt.Errorf("invalid message: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}
