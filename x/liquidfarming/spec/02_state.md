<!-- order: 2 -->

# State

The `liquidfarming` module keeps track of the states of pool coins and LFCoins.

## QueuedFarming

```go
type QueuedFarming struct {
	PoolId      uint64  // Corresponding pool id of the target liquid farm
	Amount 		sdk.Int // amount to liquid farm
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
	PoolId                 uint64 // Corresponding pool id of the target liquid farm
	BiddingCoinDenom       string // corresponding pool coin denom
	PayingReserveAddress   string
	StartTime              time.Time
	EndTime                time.Time
	Status                 AuctionStatus
	WinnerBidId            uint64 // id of winning bid of the auction
}
```

## Bid

```go
// Bid defines a standard bid for an auction.
type Bid struct {
	PoolId        	uint64
	Bidder    		string
	BiddingCoin    	sdk.Coin
}
```

## Parameter

- ModuleName: `liquidfarming`
- RouterKey: `liquidfarming`
- StoreKey: `liquidfarming`
- QuerierRoute: `liquidfarming`

## Store
- `LastDepositRequestIdKeyPrefix: []byte{0xe1} PoolId → Uint64Value(uint64)`
- `LastBidIdKeyPrefix: []byte{0xe2} AuctionId → Uint64Value(BidId)`
- `LastRewardsAuctionIdKey: []byte{0xe3} → Uint64Value(uint64)`
- `DepositRequestKey: []byte{0xe4} | PoolId | DepositRequestId -> ProtocolBuffer(DepositRequest)`
- `DepositRequestIndexKey: []byte{0xe5} | DepositorAddressLen (1 byte) | DepositorAddress | PoolId | DepositRequestId -> nil`
- `AuctionKey: []byte{0xe8} | LiquidFarmId | AuctionId -> ProtocolBuffer(AuctionState)`