<!-- order: 2 -->

# State

The `liquidfarming` module keeps track of the states of pool coins and LFCoins.

## QueuedFarming

```go
type QueuedFarming struct {
	PoolId uint64  // Corresponding pool id of the target liquid farm
	Amount sdk.Int // amount to liquid farm
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
	Id                   uint64
	PoolId               uint64 // Corresponding pool id of the target liquid farm
	BiddingCoinDenom     string // corresponding pool coin denom
	PayingReserveAddress string
	StartTime            time.Time
	EndTime              time.Time
	Status               AuctionStatus
	Winner               string // winner's account address
	Rewards              sdk.Coins
}
```

## Bid

```go
// Bid defines a standard bid for an auction.
type Bid struct {
	PoolId      uint64
	Bidder      string
	BiddingCoin sdk.Coin
}
```

## Parameter

- ModuleName: `liquidfarming`
- RouterKey: `liquidfarming`
- StoreKey: `liquidfarming`
- QuerierRoute: `liquidfarming`

## Store

- LastRewardsAuctionIdKey: `[]byte{0xe1} | PoolId -> Uint64Value(uint64)`
- QueuedFarmingKey: `[]byte{0xe4} | EndTimeLen (1 byte) | EndTime | FarmingCoinDenomLen (1 byte) | FarmingCoinDenom |  FarmerAddressLen (1 byte) | FarmerAddress -> ProtocolBuffer(QueuedFarming)`
- QueuedFarmingIndexKey: `[]byte{0xe5} | FarmerAddressLen (1 byte) | FarmerAddress | FarmingCoinDenomLen (1 byte) | FarmingCoinDenom | EndTimeLen (1 byte) | EndTime | FarmingCoinDenomLen (1 byte) -> nil`
- RewardsAuctionKey: `[]byte{0xe7} | PoolId | AuctionId -> ProtocolBuffer(RewardsAuction)`
- BidKey: `[]byte{0xea} | PoolId | BidderAddressLen (1 byte) | BidderAddress -> ProtocolBuffer(Bid)`
- WinningBidKey: `[]byte{0xeb} | PoolId | AuctionId -> ProtocolBuffer(WinningBid)`