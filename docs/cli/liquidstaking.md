---
Title: Liquidstaking
Description: A high-level overview of how the command-line interfaces (CLI) works for the liquidstaking module.
---

# Liquidstaking Module

## Synopsis

This document provides a high-level overview of how the command line (CLI) interface works for the `liquidstaking` module. To set up a local testing environment, it requires the latest [Starport](https://starport.com/). If you don't have Starport set up in your local machine, see [this Starport guide](https://docs.starport.network/) to install it. Run this command under the project root directory `$ starport chain serve`.

Note that [jq](https://stedolan.github.io/jq/) is recommended to be installed as it is used to process JSON throughout the document.

- [Transaction](#Transaction)
    * [LiquidStake](#LiquidStake)
    * [LiquidUnstake](#LiquidUnstake)
- [Query](#Query)
    * [Params](#Params)
    * [LiquidValidators](#LiquidValidators)
    * [States](#States)
    * [VotingPower](#VotingPower)

# Transaction

## LiquidStake

Liquid stake coin.

It requires `whitelisted_validators` to be registered. The [config.yml](https://github.com/cosmosquad-labs/squad/blob/main/config.yml) file registers a single whitelist validator for testing purpose. 

Usage

```bash
liquid-stake [amount]
```

| **Argument** |  **Description**                                          |
| :----------- | :-------------------------------------------------------- |
| amount       | amount of coin to liquid stake; it must be the bond denom |

Example

```bash
squad tx liquidstaking liquid-stake 5000000000stake \
--chain-id localnet \
--from bob \
--keyring-backend test \
--gas 1000000 \
--broadcast-mode block \
--yes \
--output json | jq

#
# Tips
#
# Query account balances
# Notice the newly minted bToken
squad q bank balances cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu -o json | jq

# Query the voter's liquid staking voting power
squad q liquidstaking voting-power cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu -o json | jq
```

## LiquidUnstake

Unstake coin.

Usage

```bash
liquid-unstake [amount]
```

| **Argument**  |  **Description**                                      |
| :------------ | :---------------------------------------------------- |
| amount        | amount of coin to unstake; it must be the bToken denom|

Example

```bash
squad tx liquidstaking liquid-unstake 1000000000bstake \
--chain-id localnet \
--from bob \
--keyring-backend test \
--gas 1000000 \
--broadcast-mode block \
--yes \
--output json | jq

#
# Tips
#
# Query account balances
# Notice the newly minted bToken
squad q bank balances cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu -o json | jq

# Query the voter's liquid staking voting power
squad q liquidstaking voting-power cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu -o json | jq
```

# Query

## Params

Query the current liquidstaking parameters information

Usage

```bash
params
```

Example

```bash
squad query liquidstaking params -o json | jq
```

## LiquidValidators

Query all liquid validators

Usage

```bash
liquid-validators
```

Example

```bash
squad query liquidstaking liquid-validators -o json | jq
```
## States

Query net amount state.

Usage

```bash
states
```

Example

```bash
squad query liquidstaking states -o json | jq
```

## VotingPower

Query the voter’s staking and liquid staking voting power 

Usage

```bash
voting-power [voter]
```

| **Argument** |  **Description**      |
| :----------- | :-------------------- |
| voter        | voter account address |

Example

```bash
squad query liquidstaking voting-power cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu -o json | jq
```