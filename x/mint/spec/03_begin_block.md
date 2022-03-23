<!--
order: 3
-->

# Begin-Block

Begin block operation for the `mint` module sets `LastBlockTime` value and calculate `BlockInflation` to mint coins to be sent to the fee collector.
As genesis block doesn't have last block time, there is no inflation.

## Inflation Calculation

At the beginning of each block, block inflration is calculated with the following calculation.

```
BlockInflation = InflationScheduleAmount * min(BlockDurationForInflation, BlockTimeThreshold) / (InflationScheduleEndTime - InflationScheduleStartTime)
```
