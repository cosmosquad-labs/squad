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

It is a list of inflation schedules to mint coins to be sent to the fee collector account. The inflation schedules cannot overlap from one another, start time is inclusive of the current time, and end time is exclusive for each schedule so the end times and other start times could be the same. 

- `InflationSchedule` defines the start and end time of the inflation period, and the amount of inflation during that period.
- `InflationSchedule.Amount` should be over the inflation schedule duration seconds to avoid decimal loss

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