package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/version"
)

func NewSubmitLiquidFarmProposalTxCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "liquidfarm [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a liquidfarm proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a liquidfarm along with an initial deposit.
The proposal details must be supplied via a JSON file. For values that contains
objects, only non-empty fields will be updated.

Proper vetting of a parameter change proposal should prevent this from happening
(no deposits should occur during the governance process), but it should be noted
regardless.

Example:
$ %s tx gov submit-proposal liquidfarm <path/to/proposal.json> --from=<key_or_address>

Where proposal.json contains:

{
  "title": "Create LiquidFarm",
  "description": "Create LiquidFarm For bCRE/ATOM Pool",
  "liquidfarm": [
    {
      "pool_id": 1,
      "pool_coin_denom": "",
	  "lf_coin_denom": ""
    }
  ],
  "deposit": "1000000ucre"
}
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// clientCtx, err := client.GetClientTxContext(cmd)
			// if err != nil {
			// 	return err
			// }

			// proposal, err := paramscutils.ParseParamChangeProposalJSON(clientCtx.LegacyAmino, args[0])
			// if err != nil {
			// 	return err
			// }

			// from := clientCtx.GetFromAddress()
			// content := paramproposal.NewParameterChangeProposal(
			// 	proposal.Title, proposal.Description, proposal.Changes.ToParamChanges(),
			// )

			// deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			// if err != nil {
			// 	return err
			// }

			// msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			// if err != nil {
			// 	return err
			// }

			// return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			return nil
		},
	}
}
