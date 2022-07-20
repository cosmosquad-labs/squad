package cli_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil"

	"github.com/cosmosquad-labs/squad/v2/app/params"

	"github.com/cosmosquad-labs/squad/v2/x/marketmaker/client/cli"
)

func TestParsePublicPlanProposal(t *testing.T) {
	encodingConfig := params.MakeTestEncodingConfig()

	okJSON := testutil.WriteToNewTempFile(t, `
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
`)

	proposal, err := cli.ParseMarketMakerProposal(encodingConfig.Marshaler, okJSON.Name())
	require.NoError(t, err)

	require.Equal(t, "Market Maker Proposal", proposal.Title)
	require.Equal(t, "Are you ready to market making?", proposal.Description)
}
