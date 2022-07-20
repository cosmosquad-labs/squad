package cli

// DONTCOVER
// client is excluded from test coverage in MVP version

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
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/cosmosquad-labs/squad/v2/x/marketmaker/types"
)

// GetTxCmd returns a root CLI command handler for all x/marketmaker transaction commands.
func GetTxCmd() *cobra.Command {
	marketmakerTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "MarketMaker transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	marketmakerTxCmd.AddCommand(
		NewApplyMarketMaker(),
		NewClaimIncentives(),
	)

	return marketmakerTxCmd
}

// TODO: fix to Eligible
// NewApplyMarketMaker implements the harvest rewards command handler.
func NewApplyMarketMaker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply [pool-ids]",
		Args:  cobra.MaximumNArgs(1),
		Short: "Apply to be a market maker",
		Long: strings.TrimSpace(
			fmt.Sprintf(`TBD.

Example:
$ %s tx %s apply 1 --from mykey
$ %s tx %s apply 1,2 --from mykey
`,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			farmer := clientCtx.GetFromAddress()
			pairIds := []uint64{}
			switch len(args) {
			case 1:
				pairIdsStr := strings.Split(args[0], ",")

				for _, i := range pairIdsStr {
					pairId, err := strconv.ParseUint(i, 10, 64)
					if err != nil {
						return fmt.Errorf("parse plan id: %w", err)
					}
					pairIds = append(pairIds, pairId)
				}
			default:
				return fmt.Errorf("either staking-coin-denoms or --all flag must be specified")
			}

			msg := types.NewMsgApplyMarketMaker(farmer, pairIds)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewClaimIncentives implements the remove plan handler.
func NewClaimIncentives() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim",
		Args:  cobra.ExactArgs(0),
		Short: "Claim all claimable incentives",
		Long: fmt.Sprintf(`TBD.
Example:
$ %s tx %s claim --from mykey`,
			version.AppName, types.ModuleName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress()

			msg := types.NewMsgClaimIncentives(creator)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// TODO: fix example json
// GetCmdSubmitMarketMakerProposal implements the inclusion/exclusion/distribution for market maker command handler.
func GetCmdSubmitMarketMakerProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market-maker-proposal [proposal-file] [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a market maker proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a market maker proposal along with an initial deposit. You can submit this governance proposal
to inclusion, exclusion and incentive distribution for market maker. The proposal details must be supplied via a JSON file. A JSON file to add plan request proposal is 
provided below.

Example:
$ %s tx gov submit-proposal market-maker-proposal <path/to/proposal.json> --from=<key_or_address> --deposit=<deposit_amount>

Where proposal.json contains:

{
  "title": "Market Maker Proposal",
  "description": "Are you ready to market making?",
  "inclusions": [
    {
      "address": "cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu",
      "pair_id": 1
    }
  ],
  "exclusions": [],
  "distributions": [
    {
      "address": "cosmos1kgdhngd6tfmfav62r7635h429ua7z84edpay0u",
      "pair_id": 2,
      "amount": [{"denom":"stake", "amount":"100000000"}]
    }
  ]
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

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			proposal, err := ParseMarketMakerProposal(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}

			content := types.NewMarketMakerProposal(
				proposal.Title,
				proposal.Description,
				proposal.Inclusions,
				proposal.Exclusions,
				proposal.Rejections,
				proposal.Distributions,
			)

			from := clientCtx.GetFromAddress()

			msg, err := gov.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")

	return cmd
}
