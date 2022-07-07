<!-- order: 2 -->

# State

The `liquidfarming` module keeps track of the states of pool coins and LFCoins.

## DepositRequest

```go
type DepositRequest struct {
	PoolId      uint64
	Id          uint64
	Depositor   string
	DepositCoin sdk.Coin
}
```

## RewardsAuction

```go
// AuctionStatus enumerates the valid status of an auction.
type AuctionStatus int32

const (
	AuctionStatusNil      AuctionStatus = 0
	AuctionStatusStarted  AuctionStatus = 1
	AuctionStatusFinished AuctionStatus = 2
)

type RewardsAuction struct {
	Id                     uint64
	PoolId                 uint64
	SellingRewards         sdk.Coins
	BiddingCoinDenom       string // pool coin denom
	SellingReserverAddress string
	PayingReserveAddress   string
	StartTime              time.Time
	EndTime                time.Time
	Status                 AuctionStatus
	WinnerBidId            uint64
}
```

## Bid

```go
// Bid defines a standard bid for an auction.
type Bid struct {
	Id        uint64
	AuctionId uint64
	Bidder    string
	Amount    sdk.Coin
	IsWinner  bool
}
```

## Parameter

- ModuleName: `liquidfarming`
- RouterKey: `liquidfarming`
- StoreKey: `liquidfarming`
- QuerierRoute: `liquidfarming`

## Store

...