package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/cosmosquad-labs/squad/v2/x/marketmaker/client/cli"
	"github.com/cosmosquad-labs/squad/v2/x/marketmaker/client/rest"
)

// ProposalHandler is the public plan command handler.
// Note that rest.ProposalRESTHandler will be deprecated in the future.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitMarketMakerProposal, rest.ProposalRESTHandler)
)
