<!-- order: 6 -->

# Parameters

The `liquidfarming` module contains the following parameters:

| Key                        | Type         | Example                                        |
| -------------------------- | ------------ | ---------------------------------------------- |
| LiquidFarmCreationFee      | sdk.Coins    | [{"denom":"stake","amount":"100000000"}]       |
| LiquidFarms                | []LiquidFarm | TBD                                            |


## LiquidFarmCreationFee

`LiquidFarmCreationFee` ...

## LiquidFarms

`LiquidFarms` is a list of `LiquidFarm` ...

```go
type LiquidFarm struct {
	PoolId               uint64
	MinimumDepositAmount sdk.Int
	MinimumBidAmount     sdk.Int
}
```
