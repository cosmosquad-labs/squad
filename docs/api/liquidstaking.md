---
Title: Liquidstaking
Description: A high-level overview of what gRPC-gateway REST routes are supported in the liquidstaking module.
---

# Liquidstaking Module

## Synopsis

This document provides a high-level overview of what gRPC-gateway REST routes are supported in the `liquidstaking` module.

## gRPC-gateway REST Routes

<!-- markdown-link-check-disable -->
++https://github.com/cosmosquad-labs/squad/blob/main/proto/squad/liquidstaking/v1beta1/query.proto 

- [Params](#Params)
- [Validators](#Validators)
- [VotingPower](#VotingPower)
- [States](#States)

## Params

Example Request

```bash
http://localhost:1317/squad/liquidstaking/v1beta1/params
```

Example Response

```json
{
  "params": {
    "liquid_bond_denom": "bstake",
    "whitelisted_validators": [
      {
        "validator_address": "cosmosvaloper1zaavvzxez0elundtn32qnk9lkm8kmcsz8ycjrl",
        "target_weight": "100000000"
      }
    ],
    "unstake_fee_rate": "0.001000000000000000",
    "min_liquid_staking_amount": "1000000"
  }
}
```

## Validators

Example Request

```bash
http://localhost:1317/squad/liquidstaking/v1beta1/validators
```

Example Response

```json
{
  "liquid_validators": [
    {
      "operator_address": "cosmosvaloper1zaavvzxez0elundtn32qnk9lkm8kmcsz8ycjrl",
      "weight": "100000000",
      "status": "VALIDATOR_STATUS_ACTIVE",
      "del_shares": "0.000000000000000000",
      "liquid_tokens": "0"
    }
  ]
}
```

## VotingPower

Example Request

```bash
http://localhost:1317/squad/liquidstaking/v1beta1/voting_power/cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu
```

Example Response

```json
{
  "voting_power": {
    "voter": "cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu",
    "staking_voting_power": "0",
    "liquid_staking_voting_power": "5000000000",
    "validator_voting_power": "0"
  }
}
```

## States

Example Request

```bash
http://localhost:1317/squad/liquidstaking/v1beta1/states
```

Example Response

```json
{
  "net_amount_state": {
    "mint_rate": "0.999682079425781607",
    "btoken_total_supply": "5000000000",
    "net_amount": "5001590108.399267325000000000",
    "total_del_shares": "5000000000.000000000000000000",
    "total_liquid_tokens": "5000000000",
    "total_remaining_rewards": "1590108.399267325000000000",
    "total_unbonding_balance": "0",
    "proxy_acc_balance": "0"
  }
}
```
