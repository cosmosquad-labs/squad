<!-- order: 2 -->

# State

```go
type ClaimRecord struct {
 	Address                  string  // 
	InitialClaimableAmount   sdk.Int //
	UnclaimedClaimableAmount sdk.Int //
	DepositAction            bool    // 
	SwapAction               bool    //
	StakeAction              bool    //
}
```