<!-- order: 6 -->

# Events

The `liquidfarming` module emits the following events:

## Handlers

### MsgDeposit

| Type       | Attribute Key      | Attribute Value        |
| ---------- | ------------------ | ---------------------- |
| deposit    | pool_id            | {poolId}               |
| deposit    | depositor          | {depositor}            |
| deposit    | deposit_coin       | {depositCoin}          |
| message    | module             | liquidfarming          |
| message    | action             | deposit                |
| message    | depositor          | {depositorAddress}     |

### MsgCancel

| Type       | Attribute Key        | Attribute Value        |
| ---------- | -------------------- | ---------------------- |
| cancel     | pool_id              | {poolId}               |
| cancel     | deposit_request_id   | {depositRequestId}     |
| cancel     | depositor            | {depositor}            |
| message    | module               | liquidfarming          |
| message    | action               | deposit                |
| message    | depositor            | {depositorAddress}     |

### MsgWithdraw

| Type       | Attribute Key      | Attribute Value        |
| ---------- | ------------------ | ---------------------- |
| withdraw   | pool_id            | {poolId}               |
| withdraw   | withdrawer         | {withdrawer}           |
| withdraw   | withdrawing_coin   | {withdrawingCoin}      |
| withdraw   | withdrawn_coin     | {withdrawnCoin}        |
| message    | module             | liquidfarming          |
| message    | action             | deposit                |
| message    | withdrawer         | {withdrawerAddress}    |

### MsgPlaceBid

| Type       | Attribute Key      | Attribute Value        |
| ---------- | ------------------ | ---------------------- |
| place_bid  | auction_id         | {auctionId}            |
| place_bid  | bidder             | {bidder}               |
| place_bid  | amount             | {amount}               |
| message    | module             | liquidfarming          |
| message    | action             | deposit                |
| message    | bidder             | {bidderAddress}        |
