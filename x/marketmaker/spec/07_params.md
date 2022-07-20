<!-- order: 7 -->

# Parameters

# Parameters

The `marketmaker` module contains the following parameters:

| Key                    | Type               | Example                                                        |
|------------------------|--------------------|----------------------------------------------------------------|
| IncentiveBudgetAddress | string             | cre1ddn66jv0sjpmck0ptegmhmqtn35qsg2vxyk2hn9sqf4qxtzqz3sq3qhhde |
| IntentDepositAmount    | string (sdk.Coins) | [{"denom":"ucre","amount":"1000000000"}]                       |
| IncentivePairs         | []IncentivePair    | see below                                                      |
|                        |                    |                                                                |
|                        |                    |                                                                |

## IncentiveBudgetAddress

## IntentDepositAmount

## IncentivePair

```go
type IncentivePairs struct {
    PairName            string
    PairId              uint64
    UpdateTime          time.Time
    MinDepth            sdk.Int 
    MaxSpread           sdk.Dec 
    MinWidth            sdk.Dec 
    RemnantThreshold    sdk.Dec 
    DailyUptimeIntent   sdk.Dec 
    DailyUptime         sdk.Dec 
    MonthlyUptimeIntent sdk.Dec 
    MonthlyUptime       sdk.Dec 
    IncentiveWeight     sdk.Dec 
}
```

ex)

```go
[{
	PairName: "atom-usdc",
	PairID: 1,
	UpdateTime: 2022-09-01T00:00:00Z
	MinDepth: 5000,($1000 value)
	MaxSpread: 0.60,
	MinWidth: 0.15,
	RemnantTreshold: 0.7,
	DailyUptimeIntent: 0.875,
	DailyUptime: 0.835,
	MonthlyUptimeIntent: 0.90,
	MonthlyUptime: 0.85,
	IncentiveWeight: 0.1
}, 
{
	PairName: "eth-usdc",
	PairID: 2,
	UpdateTime: 2022-09-01T00:00:00Z
	MinDepth: 5000,($1000 value)
	MaxSpread: 0.60,
	MinWidth: 0.15,
	RemnantTreshold: 0.7,
	DailyUptimeIntent: 0.875,
	DailyUptime: 0.835,
	MonthlyUptimeIntent: 0.90,
	MonthlyUptime: 0.85,
	IncentiveWeight: 0.1
}
]
```