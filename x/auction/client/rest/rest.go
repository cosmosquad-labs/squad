package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
)

// REST Variable names
const (
	RestParamsType = "type"
	RestAuctionID  = "auction-id"
	RestAuctioneer = "auctioneer"
)

// AuctionRESTHandler defines a REST handler implemented in another module. The
// sub-route is mounted on the governance REST handler.
type AuctionRESTHandler struct {
	SubRoute string
	Handler  func(http.ResponseWriter, *http.Request)
}

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router, phs []AuctionRESTHandler) {
	fmt.Println(">>> RegisterHandlers...")
	// r := clientrest.WithHTTPDeprecationHeaders(rtr)
	// registerQueryRoutes(clientCtx, r)
	// registerTxHandlers(clientCtx, r, phs)
}
