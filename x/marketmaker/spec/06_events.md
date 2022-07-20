<!-- order: 6 -->

# Events

The `marketmaker` module emits the following events:

## Handlers

### MsgIntentMarketMaker

| Type               | Attribute Key | Attribute Value    |
|--------------------|---------------|--------------------|
| intent_marketmaker | address       | {mmAddress}        |
| intent_marketmaker | pair_ids      | []{pairId}         |
| message            | module        | marketmaker        |
| message            | action        | intent_marketmaker |
| message            | sender        | {senderAddress}    |

### MsgClaimIncentives

| Type             | Attribute Key | Attribute Value  |
|------------------|---------------|------------------|
| claim_incentives | address       | {mmAddress}      |
| message          | module        | marketmaker      |
| message          | action        | claim_incentives |
| message          | sender        | {senderAddress}  |