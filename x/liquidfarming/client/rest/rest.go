package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"

	auctionrest "github.com/cosmosquad-labs/squad/x/auction/client/rest"
)

func AuctionRESTHandler(clientCtx client.Context) auctionrest.AuctionRESTHandler {
	return auctionrest.AuctionRESTHandler{
		SubRoute: "fixed_price_auction",
		Handler:  postAuctionHandlerFn(clientCtx),
	}
}

func postAuctionHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(">>> postAuctionHandlerFn...")

		// var req paramscutils.ParamChangeProposalReq
		// if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
		// 	return
		// }

		// req.BaseReq = req.BaseReq.Sanitize()
		// if !req.BaseReq.ValidateBasic(w) {
		// 	return
		// }

		// content := proposal.NewParameterChangeProposal(req.Title, req.Description, req.Changes.ToParamChanges())

		// msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, req.Proposer)
		// if rest.CheckBadRequestError(w, err) {
		// 	return
		// }
		// if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
		// 	return
		// }

		// tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)

	}
}
