package client

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmosquad-labs/squad/x/auction/client/rest"
)

// function to create the rest handler
type RESTHandlerFn func(client.Context) rest.AuctionRESTHandler

// function to create the cli handler
type CLIHandlerFn func() *cobra.Command

// The combined type for a proposal handler for both cli and rest
type AuctionHandler struct {
	CLIHandler  CLIHandlerFn
	RESTHandler RESTHandlerFn
}

// NewAuctionHandler creates a new AuctionHandler object
func NewAuctionHandler(cliHandler CLIHandlerFn, restHandler RESTHandlerFn) AuctionHandler {
	return AuctionHandler{
		CLIHandler:  cliHandler,
		RESTHandler: restHandler,
	}
}
