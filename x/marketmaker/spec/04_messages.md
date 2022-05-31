<!-- order: 4 -->

# Messages

## MsgIntentMarketMaker

```go
type MsgIntentMarketMaker struct {
    // orderer specifies the bech32-encoded address of the market maker that will makes an order
    Address string
    // pair_id specifies the pair ids
    PairIds []uint64
}
```

## MsgClaimIncentives

```go
type MsgClaimIncentives struct {
    // address specifies the bech32-encoded address of the market maker that claim incentives
    Address string
}
```