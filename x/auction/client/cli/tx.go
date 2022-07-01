package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

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

	for _, propCmd := range propCmds {
		flags.AddTxFlagsToCmd(propCmd)
		auctionTxCmd.AddCommand(propCmd)
	}

	return auctionTxCmd
}
