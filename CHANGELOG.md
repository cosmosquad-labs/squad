<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking Protobuf, gRPC and REST routes used by end-users.
"CLI Breaking" for breaking CLI commands.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState given same genesisState and txList.
Ref: https://keepachangelog.com/en/1.0.0/
-->
<!-- markdown-link-check-disable -->

# Changelog

## [Unreleased]

## v2.1.1

### Improvements

* (x/liquidity) [\#344](https://github.com/cosmosquad-labs/squad/pull/344), [\#341](https://github.com/cosmosquad-labs/squad/pull/341) Enable detailed configuration for order book responses

### Bug Fixes
* (x/liquidstaking) [\#342](https://github.com/cosmosquad-labs/squad/pull/342) fix: calculation bug of liquidstaking voting_power endpoint

## [v2.1.0] - 2022-07-18

### Client Breaking Changes

* (x/liquidity) [\#338](https://github.com/cosmosquad-labs/squad/pull/338) Refactor `OrderBooks` query:
  * `tick_precisions` field has been removed from `QueryOrderBooksRequest`
  * `tick_precision` field has been removed from `OrderBookResponse` and `price_unit` has been added instead
  * The order between `sells` and `buys` has been changed
* (x/liquidity) [\#335](https://github.com/cosmosquad-labs/squad/pull/335) Modify `PoolResponse`:
  * `balances` field has been modified to contain `base_coin` and `quote_coin` fields
  * `pool_coin_supply` field has been added
  * `price` field has been added
* (x/farming) [\#334](https://github.com/cosmosquad-labs/squad/pull/334) Add `HistoricalRewards` query endpoint:
  * `HistoricalRewards`: `/squad/farming/v1beta1/historical_rewards/{staking_coin_denom}`
* (x/liquidity) [\#318](https://github.com/cosmosquad-labs/squad/pull/318) Add `OrderBooks` query endpoint:
  * `OrderBooks`: `/squad/liquidity/v1beta1/order_books`
* (x/farming) [\#305](https://github.com/cosmosquad-labs/squad/pull/305) Rename existing `Stakings` endpoint to `Position` and add three new endpoints:
  * `Stakings`: `/squad/farming/v1beta1/stakings/{farmer}`
  * `QueuedStakings`: `/squad/farming/v1beta1/queued_stakings/{farmer}`
  * `UnharvestedRewards`: `/squad/farming/v1beta1/unharvested_reward/{farmer}`

### CLI Breaking Changes

* (x/liquidity) [\#338](https://github.com/cosmosquad-labs/squad/pull/338) Refactor `order-books` query cmd:
  * `[tick-precisions]` argument has been removed: `order-books [pair-ids]`
  * Response structure has been changed
* (x/farming) [\#334](https://github.com/cosmosquad-labs/squad/pull/334) Add `historical-rewards` query cmd:
  * `historical-rewards [staking-coin-denom]`
* (x/liquidity) [\#318](https://github.com/cosmosquad-labs/squad/pull/318) Add `create-ranged-pool` tx cmd and `order-books` query cmd:
  * `create-ranged-pool [pair-id] [deposit-coins] [min-price] [max-price] [initial-price]`
  * `order-books [pair-ids] [tick-precisions]`
* (x/farming) [\#305](https://github.com/cosmosquad-labs/squad/pull/305) Rename existing `stakings` query to `position` and add three new queries:
  * `stakings [farmer]`
  * `queued-stakings [farmer]`
  * `unharvested-rewards [farmer]`

### State Machine Breaking

* (x/liquidity) [\#340](https://github.com/cosmosquad-labs/squad/pull/340) Add `MaxNumActivePoolsPerPair` global constant
* (x/liquidity) [\#318](https://github.com/cosmosquad-labs/squad/pull/318) Change `Pool` struct for ranged pools and refactor matching logic
  * Add `type`, `creator`, `min_price` and `max_price` fields to `Pool` struct
  * Refactor matching logic both for fairness of matching and efficiency of pool order placement
  * Change the liquidity module's `TickPrecisions` param from 3 to 4
* (x/farming) [\#305](https://github.com/cosmosquad-labs/squad/pull/305) Time-based queued staking and new UnharvestedRewards struct
  * Changed/added kv-store keys:
    * QueuedStaking: `0x23 | EndTimeLen (1 byte) | sdk.FormatTimeBytes(EndTime) | StakingCoinDenomLen (1 byte) | StakingCoinDenom | FarmerAddr -> ProtocolBuffer(QueuedStaking)`
    * QueuedStakingIndex: `0x24 | FarmerAddrLen (1 byte) | FarmerAddr | StakingCoinDenomLen (1 byte) | StakingCoinDenom | sdk.FormatTimeBytes(EndTime) -> nil`
    * UnharvestedRewards: `0x34 | FarmerAddrLen (1 byte) | FarmerAddr | StakingCoinDenom -> ProtocolBuffer(UnharvestedRewards)`
* (x/mint, x/budget) [\#316](https://github.com/cosmosquad-labs/squad/pull/316) feat!: add mint pool address for mint module #316
  * Add `params.MintPoolAddress` on the mint module `cre1m3h30wlvsf8llruxtpukdvsy0km2kum8ve5ajd`
  * Change Mint Pool from default `cre17xpfvakm2amg962yls6f84z3kell8c5l53s97s` (fee_collector) to `cre1m3h30wlvsf8llruxtpukdvsy0km2kum8ve5ajd` (params.MintPoolAddress) to prevent mixing of inflation and tx fee
  * Change the source address of Budgets whose source address is `cre17xpfvakm2amg962yls6f84z3kell8c5l53s97s` to `cre1m3h30wlvsf8llruxtpukdvsy0km2kum8ve5ajd`
  * Add Budget to sending staking reward, and community fund to `cre17xpfvakm2amg962yls6f84z3kell8c5l53s97s` from `cre1m3h30wlvsf8llruxtpukdvsy0km2kum8ve5ajd`

* [\#303](https://github.com/cosmosquad-labs/squad/pull/303) build!: bump cosmos-sdk to v0.45.3, tendermint v0.34.19, ibc-go v2.2.0, budget v1.2.0, go 1.17

### Improvements

* (x/liquidity) [\#310](https://github.com/cosmosquad-labs/squad/pull/310) feat: add emit events for order trading volume

### Bug Fixes

* (x/liquidity) [\#296](https://github.com/cosmosquad-labs/squad/pull/296) [\#299](https://github.com/cosmosquad-labs/squad/pull/299) fix: optimize CancelAllOrders gas usage, fix offer coin checking
* (x/claim) [\#292](https://github.com/cosmosquad-labs/squad/pull/292) [\#304](https://github.com/cosmosquad-labs/squad/pull/304) fix: fix simulation for the claim module
*  [\#298](https://github.com/cosmosquad-labs/squad/pull/298) fix: fix to use query context

## [v1.1.0] - 2022-04-18

### State Machine Breaking

Running a full node will encounter wrong app hash issue if it doesn't upgrade to this version prior to `UpgradeHeight (48000)`. Instead of going through on-chain governance proposal by using `UpgradeProposal`, this upgrade mechanism is chosen as it is security hot fix that is better to be fixed as soon as it can and also it is directly related to governance proposal.

* (x/claim) [\#294](https://github.com/cosmosquad-labs/squad/pull/294) Fix gas consumption issue for `ConditionTypeVote`. `UpgradeHeight` is set as `48000`.

## [v1.0.0] - 2022-04-12

### Features

* (squad) feat: add `x/liquidity` module
* (squad) feat: add `x/liquidstaking` module
* (squad) feat: add `x/farming` module
* (squad) feat: add `x/mint`(constant inflation) module
* (squad) feat: add `x/claim` module
* (sdk) Squad Core uses a customized Cosmos SDK [v1.0.2-sdk-0.44.5](https://github.com/cosmosquad-labs/cosmos-sdk/releases/v1.0.2-sdk-0.44.5). Please check the differences on [here](https://github.com/cosmosquad-labs/cosmos-sdk/compare/v0.44.5...v1.0.2-sdk-0.44.5).
  * `x/staking` fix: allow delegate only spendable coins
  * `x/gov` feat: add additional voting powers hook on tally (liquid governance)
  * `x/vesting` feat: periodic vesting msg
  * `x/bank` feat: Add dynamic blockedAddrs

[Unreleased]: https://github.com/cosmosquad-labs/squad/compare/v2.1.0...HEAD
[v1.0.0]: https://github.com/cosmosquad-labs/squad/releases/tag/v1.0.0
[v1.1.0]: https://github.com/cosmosquad-labs/squad/releases/tag/v1.1.0
[v2.1.0]: https://github.com/cosmosquad-labs/squad/releases/tag/v2.1.0