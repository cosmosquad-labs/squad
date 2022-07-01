package client

import (
	auctionclient "github.com/cosmosquad-labs/squad/x/auction/client"
	"github.com/cosmosquad-labs/squad/x/liquidfarming/client/cli"
	"github.com/cosmosquad-labs/squad/x/liquidfarming/client/rest"
)

// AuctionHandler is the fixed price auction creation handler.
var AuctionHandler = auctionclient.NewAuctionHandler(cli.NewCreateFixedPriceAuctionTxCmd, rest.AuctionRESTHandler)
