package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/cosmosquad-labs/squad/x/liquidfarming/client/cli"
	"github.com/cosmosquad-labs/squad/x/liquidfarming/client/rest"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitLiquidFarmProposalTxCmd, rest.ProposalRESTHandler)
