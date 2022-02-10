<!-- order: 2 -->

# State

```go
type ClaimRecord struct {
 	Address                  string  // 
	InitialClaimableAmount   sdk.Int //
	RemainingClaimableAmount sdk.Int //
	DepositAction            bool    // 
	SwapAction               bool    //
	StakeAction              bool    //
}
```