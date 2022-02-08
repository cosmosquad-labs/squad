<!-- order: 3 -->

# State Transitions

These messages (Msg) in the liquidity module trigger state transitions.

## **Coin Escrow for Liquidity Module Messages**

Transaction confirmation causes state transition on the Bank module. Some messages on the liquidity module require coin escrow before confirmation.

### **MsgDepositBatch**

To deposit coins into an existing `Pool`, the depositor must escrow `DepositCoins` into `GlobalEscrowAddr`.

### **MsgWithdrawBatch**

To withdraw coins from a `Pool`, the withdrawer must escrow `PoolCoin` into `GlobalEscrowAddr`.

### **MsgSwapBatch**

To request a coin swap, the swap requestor must escrow `OfferCoin` into each pair’s `EscrowAddress`.

## **LiquidityPoolBatch Execution**

Batch execution causes state transitions on the `Bank` module. The following categories describe state transition executed by each process in the `PoolBatch` execution.

### **Coin Swap**

After a successful coin swap, coins accumulated in `LiquidityModuleEscrowAccount` for coin swaps are sent to other swap requestors or to the `Pool`.

### **Deposit**

After a successful deposit transaction, escrowed coins are sent to the `ReserveAccount` of the targeted `Pool` and new pool coins are minted and sent to the depositor.

### **Withdrawal**

After a successful withdraw transaction, escrowed pool coins are burned and a corresponding amount of reserve coins are sent to the withdrawer from the liquidity `Pool`.

## **Change states of swap requests with expired lifespan**

After execution of `Batch`, status of all remaining swap requests with `ExpiredAt` higher than current block time are changed to `SwapRequestStatusExpired`

## **Refund escrowed coins**

Refunds are issued for escrowed coins for cancelled swap order and failed create pool, deposit, and withdraw messages.