<!-- order: 8 -->

# Proposal

## MarketMakerProposal

```go
type MarketMakerProposal struct {
	// title specifies the title of the proposal
	Title string 
	// description specifies the description of the proposal
	Description string 
	Inclusions []MarketMakerHandle
	Exclusions []MarketMakerHandle
	Distributions []IncentiveDistribution
}

type MarketMakerHandle struct {
	// registered market maker address
	Address string
	PairId uint64
}

type IncentiveDistribution struct {
	// registered market maker address
	Address string
	PairId uint64
	Amount sdk.Coins
}

```