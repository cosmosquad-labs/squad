---
Title: Liquidstaking
Description: A high-level overview of how the command-line interfaces (CLI) works for the liquidstaking module.
---

# Liquidstaking Module

## Synopsis

This document provides a high-level overview of how the command line (CLI) interface works for the `liquidstaking` module. To set up a local testing environment, it requires the latest [Starport](https://starport.com/). If you don't have Starport set up in your local machine, see [this Starport guide](https://docs.starport.network/) to install it. Run this command under the project root directory `$ starport chain serve`.

Note that [jq](https://stedolan.github.io/jq/) is recommended to be installed as it is used to process JSON throughout the document.

- [Transaction](#Transaction)
    * [CreatePair](#CreatePair)
    * [CreatePool](#CreatePool)
- [Query](#Query)
    * [Params](#Params)
    * [Pairs](#Pairs)


# Transaction

## LiquidStake

Usage

```bash
liquid-stake [amount]
```

| **Argument** |  **Description**               |
| :----------- | :----------------------------- |
| amount       | amount of coin to liquid stake |

Example

```bash
squad tx liquidstaking liquid-stake 1000000000stake \
--chain-id localnet \
--from alice \
--keyring-backend test \
--broadcast-mode block \
--yes \
--output json | jq
```

## LiquidUnstake

Usage

```bash

```

| **Argument**      |  **Description**                     |
| :---------------- | :----------------------------------- |
| base-coin-denom   | denom of the base coin for the pair  |

Example

```bash

```

# Query

## Params

## LiquidValidators

## States
