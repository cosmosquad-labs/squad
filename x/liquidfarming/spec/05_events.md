<!-- order: 5 -->

# Events

The `liquidfarming` module emits the following events:

## Handlers

### MsgFarm

| Type       | Attribute Key      | Attribute Value        |
| ---------- | ------------------ | ---------------------- |
| farm       | pool_id            | {poolId}               |
| farm       | farmer             | {farmer}               |
| farm       | farming_coin       | {farmingCoin}          |
| message    | module             | liquidfarming          |
| message    | action             | farm                   |
| message    | farmer             | {farmerAddress}        |

### MsgCancelQueuedFarming

| Type                  | Attribute Key        | Attribute Value        |
| --------------------- | -------------------- | ---------------------- |
| cancel_queued_farming | pool_id              | {poolId}               |
| cancel_queued_farming | queued_farming_id    | {queuedFarmingId}      |
| cancel_queued_farming | farmer               | {farmer}               |
| message               | module               | liquidfarming          |
| message               | action               | cancel_queued_farming  |
| message               | farmer               | {farmerAddress}        |

### MsgUnfarm

| Type       | Attribute Key      | Attribute Value        |
| ---------- | ------------------ | ---------------------- |
| unfarm     | pool_id            | {poolId}               |
| unfarm     | farmer             | {farmer}               |
| unfarm     | farming_coin       | {farmingCoin}          |
| unfarm     | farmed_coin        | {farmedCoin}           |
| message    | module             | liquidfarming          |
| message    | action             | unfarm                 |
| message    | farmer             | {farmerAddress}        |

### MsgPlaceBid

| Type       | Attribute Key      | Attribute Value        |
| ---------- | ------------------ | ---------------------- |
| place_bid  | auction_id         | {auctionId}            |
| place_bid  | bidder             | {bidder}               |
| place_bid  | amount             | {amount}               |
| message    | module             | liquidfarming          |
| message    | action             | deposit                |
| message    | bidder             | {bidderAddress}        |

### MsgRefundBid 

| Type       | Attribute Key      | Attribute Value        |
| ---------- | ------------------ | ---------------------- |
