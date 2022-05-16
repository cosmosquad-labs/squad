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

### Client Breaking Changes

* (x/farming) [\#305](https://github.com/cosmosquad-labs/squad/pull/305) Rename existing `Stakings` endpoint to `Position` and add three new endpoints:
  * `Stakings`: `/squad/farming/v1beta1/stakings/{farmer}`
  * `QueuedStakings`: `/squad/farming/v1beta1/queued_stakings/{farmer}`
  * `UnharvestedRewards`: `/squad/farming/v1beta1/unharvested_reward/{farmer}`

### CLI Breaking Changes

* (x/farming) [\#305](https://github.com/cosmosquad-labs/squad/pull/305) Rename existing `stakings` query to `position` and add three new queries:
  * `stakings [farmer]`
  * `queued-stakings [farmer]`
  * `unharvested-rewards [farmer]`

### State Machine Breaking

* (x/farming) [\#305](https://github.com/cosmosquad-labs/squad/pull/305) Time-based queued staking and new UnharvestedRewards struct
  * Changed/added kv-store keys:
    * QueuedStaking: `0x23 | EndTimeLen (1 byte) | sdk.FormatTimeBytes(EndTime) | StakingCoinDenomLen (1 byte) | StakingCoinDenom | FarmerAddr -> ProtocolBuffer(QueuedStaking)`
    * QueuedStakingIndex: `0x24 | FarmerAddrLen (1 byte) | FarmerAddr | StakingCoinDenomLen (1 byte) | StakingCoinDenom | sdk.FormatTimeBytes(EndTime) -> nil`
    * UnharvestedRewards: `0x34 | FarmerAddrLen (1 byte) | FarmerAddr | StakingCoinDenom -> ProtocolBuffer(UnharvestedRewards)`
