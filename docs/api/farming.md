---
Title: REST APIs
Description: A high-level overview of gRPC-gateway REST routes in farming module.
---

# Farming Module
 
## Synopsis

This document provides a high-level overview of what gRPC-gateway REST routes are supported in the `farming` module.


## Swagger Documentation

- Swagger Cosmos SDK Farming Module [REST and gRPC Gateway docs](https://app.swaggerhub.com/apis-docs/gravity-devs/farming/1.0.0)

## gRPC-gateway REST Routes

<!-- markdown-link-check-disable -->
++https://github.com/cosmosquad-labs/squad/blob/main/proto/squad/farming/v1beta1/query.proto 

- [Params](#Params)
- [Plans](#Plans)
- [Plan](#Plan)
- [Position](#Position)
- [Stakings](#Stakings)
- [QueuedStakings](#QueuedStakings)
- [TotalStakings](#TotalStakings)
- [Rewards](#Rewards)
- [UnharvestedRewards](#UnharvestedRewards)
- [CurrentEpochDays](#CurrentEpochDays)

### Params

Query the values set as farming parameters:

Example Request 

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/params
```


```json
{
  "params": {
    "private_plan_creation_fee": [
      {
        "denom": "stake",
        "amount": "100000000"
      }
    ],
    "next_epoch_days": 1,
    "farming_fee_collector": "cosmos1h292smhhttwy0rl3qr4p6xsvpvxc4v05s6rxtczwq3cs6qc462mqejwy8x"
  }
}
```

### Plans

Query all the farming plans exist in the network:


Example Request 

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/plans
```

```json
{
  "plans": [
    {
      "@type": "/squad.farming.v1beta1.MsgCreateRatioPlan",
      "base_plan": {
        "id": "1",
        "name": "Second Public Ratio Plan",
        "type": "PLAN_TYPE_PUBLIC",
        "farming_pool_address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky",
        "termination_address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky",
        "staking_coin_weights": [
          {
            "denom": "pool1",
            "amount": "0.500000000000000000"
          },
          {
            "denom": "pool2",
            "amount": "0.500000000000000000"
          }
        ],
        "start_time": "2021-09-10T00:00:00Z",
        "end_time": "2021-10-01T00:00:00Z",
        "terminated": false,
        "last_distribution_time": "2021-09-17T01:00:43.410373Z",
        "distributed_coins": [
          {
            "denom": "stake",
            "amount": "2399261190929"
          }
        ]
      },
      "epoch_ratio": "0.500000000000000000"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### Plan

Query a particular plan:


Example Request 

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/plans/1
```

```json
{
  "plan": {
    "@type": "/squad.farming.v1beta1.MsgCreateRatioPlan",
    "base_plan": {
      "id": "1",
      "name": "Second Public Ratio Plan",
      "type": "PLAN_TYPE_PUBLIC",
      "farming_pool_address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky",
      "termination_address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky",
      "staking_coin_weights": [
        {
          "denom": "pool1",
          "amount": "0.500000000000000000"
        },
        {
          "denom": "pool2",
          "amount": "0.500000000000000000"
        }
      ],
      "start_time": "2021-09-10T00:00:00Z",
      "end_time": "2021-10-01T00:00:00Z",
      "terminated": false,
      "last_distribution_time": "2021-09-17T01:00:43.410373Z",
      "distributed_coins": [
        {
          "denom": "stake",
          "amount": "2399261190929"
        }
      ]
    },
    "epoch_ratio": "0.500000000000000000"
  }
}
```

### Position

Query for farming position of a farmer:


Example Request 

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/position/cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny
```

```json
{
  "staked_coins": [
    {
      "denom": "pool1",
      "amount": "2500000"
    }
  ],
  "queued_coins": [
  ],
  "rewards": [
    {
      "denom": "stake",
      "amount": "1000000"
    }
  ]
}
```

Query for farming position of a farmer with the given staking coin denom

Example Request 

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/position/cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny?staking_coin_denom=pool1
```

```json
{
  "staked_coins": [
    {
      "denom": "pool1",
      "amount": "2500000"
    }
  ],
  "queued_coins": [
  ],
  "rewards": [
    {
      "denom": "stake",
      "amount": "1000000"
    }
  ]
}
```

### Stakings

Query for all stakings by a farmer:


Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/stakings/cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny
```

```json
{
  "stakings": [
    {
      "staking_coin_denom": "pool1",
      "amount": "1000000",
      "starting_epoch": "1"
    },
    {
      "staking_coin_denom": "pool2",
      "amount": "50000000",
      "starting_epoch": "2"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}
```

Query for all stakings by a farmer with the given staking coin denom

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/stakings/cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny?staking_coin_denom=pool2
```

```json
{
  "stakings": [
    {
      "staking_coin_denom": "pool2",
      "amount": "50000000",
      "starting_epoch": "2"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### QueuedStakings

Query for all queued stakings by a farmer:


Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/queued_stakings/cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny
```

```json
{
  "queued_stakings": [
    {
      "staking_coin_denom": "pool1",
      "amount": "2000000",
      "end_time": "2022-05-05T03:03:38.108665Z"
    },
    {
      "staking_coin_denom": "pool2",
      "amount": "10000000",
      "end_time": "2022-05-05T09:30:12.559128Z"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}
```

Query for queued stakings by a farmer with the given staking coin denom

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/queued_stakings/cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny?staking_coin_denom=pool2
```

```json
{
  "queued_stakings": [
    {
      "staking_coin_denom": "pool2",
      "amount": "10000000",
      "end_time": "2022-05-05T09:30:12.559128Z"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### TotalStakings

Query for total stakings by a staking coin denom: 


Example Request 

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/total_stakings/pool1
```

```json
{
  "amount": "2500000"
}
```

### Rewards

Query for all rewards by a farmer:

Example Request 

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/rewards/cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny
```

```json
{
  "rewards": [
    {
      "staking_coin_denom": "pool1",
      "rewards": [
        {
          "denom": "stake",
          "amount": "2346201014138"
        }
      ]
    },
    {
      "staking_coin_denom": "pool2",
      "rewards": [
        {
          "denom": "stake",
          "amount": "2346201014138"
        }
      ]
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}
```


Query for all rewards by a farmer with the staking coin denom:

Example Request 

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/rewards/cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny?staking_coin_denom=pool1
```

```json
{
  "rewards": [
    {
      "staking_coin_denom": "pool1",
      "rewards": [
        {
          "denom": "stake",
          "amount": "2346201014138"
        }
      ]
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### UnharvestedRewards

Query for unharvested rewards for a farmer:

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/unharvested_rewards/cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny
```

```json
{
  "unharvested_rewards": [
    {
      "staking_coin_denom": "pool1",
      "rewards": [
        {
          "denom": "stake",
          "amount": "2346201014138"
        }
      ]
    },
    {
      "staking_coin_denom": "pool2",
      "rewards": [
        {
          "denom": "stake",
          "amount": "2346201014138"
        }
      ]
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}
```

Query for unharvested rewards for a farmer with the given staking coin denom

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/unharvested_rewards/cosmos185fflsvwrz0cx46w6qada7mdy92m6kx4gqx0ny?staking_coin_denom=pool2
```

```json
{
  "unharvested_rewards": [
    {
      "staking_coin_denom": "pool2",
      "rewards": [
        {
          "denom": "stake",
          "amount": "2346201014138"
        }
      ]
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### CurrentEpochDays

Query for the current epoch days:

Example Request 

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/farming/v1beta1/current_epoch_days
```

```json
{
  "current_epoch_days": 1
}
```
