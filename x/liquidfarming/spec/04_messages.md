<!-- order: 4 -->

# Messages

Messages (Msg) are objects that trigger state transitions. Msgs are wrapped in transactions (Txs) that clients submit to the network. The Cosmos SDK wraps and unwraps `liquidfarming` module messages from transactions.

## MsgCreate

Creating new `LiquidFarm` is not possible by transaction message. It is expected to be created through governance proposal.

## MsgFarm

```go
type MsgFarm struct {
	PoolId      uint64
	Farmer      string
	FarmingCoin sdk.Coin
}
```

## MsgCancelQueuedFarming

```go
type MsgCancelQueuedFarming struct {
	Depositor       string
	PoolId          uint64
	QueuedFarmingId uint64
}
```

## MsgUnfarm

```go
type MsgUnfarm struct {
	PoolId uint64
	Farmer string
	LFCoin sdk.Coin
}
```

## MsgPlaceBid

```go
type MsgPlaceBid struct {
	AuctionId uint64
	Bidder    string
	Amount    sdk.Coin
}
```

## MsgRefundBid

```go
type MsgRefundBid struct {
	AuctionId uint64
	BidId     string
	Bidder    sdk.Coin
}
```
