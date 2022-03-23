<!--
order: 4
-->

# Parameters

The `mint` module contains the following parameters:

| Key                  | Type                | Example |
|----------------------|---------------------|---------|
| mint_denom           | string              | "stake" |
| block_time_threshold | time.duration       | "10s"   |
| inflation_schedules  | []InflationSchedule |         |

## MintDenom

MintDenom is the denomination of the coin to be minted.

## BlockTimeThreshold

It is a parameter to prevent from inflationary manipulation attacks. Although it is highly unlikely to happen, block time manipulation can be possible with a group of malicious validators who have enough voting power. In case block time delays, `BlockTimeThreshold` will be used for `BlockDurationForInflation` to calculate inflation amount for the block. If the `BlockDurationForInflation` is greater than `BlockTimeThreshold`, the block generates with `BlockTimeThreshold` value. Therefore, it is possible that the actual minted amount is less than the defined inflation schedule amount.

## InflationSchedules

It is a list of inflation schedules to mint coins to be sent to the fee collector account. Each `InflationSchedule` defines start time, end time, and an amount of inflation. The start and end times of inflation schedules can't overlap from one another. While `StartTime` is inclusive of the current time, `EndTime` is exclusive. And, `Amount` for each schedule must be greater than the amount that is converted by subtracting the `EndTime` from the `StartTime` in seconds. Reference the [validate function](https://github.com/cosmosquad-labs/squad/blob/83e40f9204e1554cf3a0f45089fba71bb5685c70/x/mint/types/params.go#L104-L126) to fully understand them.

```go
type InflationSchedule struct {
	// start_time defines the start date time for the inflation schedule
    StartTime time.Time
	// end_time defines the end date time for the inflation schedule
    EndTime   time.Time
	// amount defines the total amount of inflation for the schedule
    Amount    sdk.Int
}
```

Example of inflation schedules

```go
ExampleInflationSchedules = []InflationSchedule{
    {
        StartTime: utils.ParseTime("2022-01-01T00:00:00Z"),
        EndTime:   utils.ParseTime("2023-01-01T00:00:00Z"),
        Amount:    sdk.NewInt(300000000000000),
    },
    {
        StartTime: utils.ParseTime("2023-01-01T00:00:00Z"),
        EndTime:   utils.ParseTime("2024-01-01T00:00:00Z"),
        Amount:    sdk.NewInt(200000000000000),
    },
}
```