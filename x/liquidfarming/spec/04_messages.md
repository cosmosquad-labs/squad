<!-- order: 4 -->

# Messages

Messages (Msg) are objects that trigger state transitions. Msgs are wrapped in transactions (Txs) that clients submit to the network. The Cosmos SDK wraps and unwraps `liquidfarming` module messages from transactions.

## MsgCreate

Creating new `LiquidFarm` is not possible by transaction message. It is expected to be created through governance proposal.

## MsgDeposit

```go
type MsgDeposit struct {
    PoolId       uint64
	Depositor    string
	DepositCoin  sdk.Coin
}
```

## MsgCancel

```go
type MsgCancel struct {
	Depositor        string
	PoolId           uint64 // target pool id
	DepositReqeustId uint64 // target deposit request id
}
```

## MsgWithdraw

```go
type MsgWithdraw struct {
	PoolId     uint64   // target deposit request id
	Withdrawer string   // the bech32-encoded address that withdraws liquid farm coin
	LFCoin     sdk.Coin // withdrawing amount of LF coin
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