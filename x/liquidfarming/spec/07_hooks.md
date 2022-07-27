<!-- order: 7 -->

# Hooks

The `Liquidfarming` module uses the following hooks registered in `farming` module.

## AfterAllocateRewards Hook

```go
AfterAllocateRewards(ctx sdk.Context)
```

When `AfterAllocateRewards` hook is delivered, the following operations are performed.
- If the auction currently going on exists, the current auction becomes closed. And, 
  - the winner is chosen,
  - the rewards is harvested and sent to the winner,
  - the pool coins from the winner in the paying reserve address is sent to the module account,
  - the pool coins from the others not winner in the paying reserve address is refunded to each bidder’s account.
- A new auction is created.

## AfterStaked Hook

```go
AfterStaked(
    ctx sdk.Context, 
    farmer sdk.AccAddress, 
    stakingCoinDenom string, 
    stakingAmt sdk.Int,
)
```

When `AfterStaked` hook is delivered, the following operation is performed.
- LF coins are minted according to the mint rate.