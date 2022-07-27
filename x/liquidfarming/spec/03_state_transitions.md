<!-- order: 3 -->

# State Transitions

This document describes the state transaction operations in the `liquidfarming` module.

## Liquid farm creation



## Coin Escrow for Liquidfarming Module Messages

The following messages cause state transition on the `bank`, `liquidty`, and `farming` modules.

### MsgFarm

- Pool coins are sent to a reserve address of a liquid farm.
- The `liquidfarming` module stakes the pool coins to the `farming` module.

### MsgCancelQueuedFarming

- The `liquidfarming` module unstakes pool coins from the `farming` module. 
- The pool coins are sent from a reserve address of a liquid farm to a farmer.

### MsgUnfarm

- LF coins are sent to the `liquidfarm` module account, and the LF coins are burnt.
- The `liquidfarming` module unstakes pool coins from the `farming` module. 
- The pool coins are sent from a reserve address of a liquid farm to a farmer.

### MsgUnfarmWithdraw

- LF coins are sent to the `liquidfarm` module account, and the LF coins are burnt.
- The `liquidfarming` module unstakes pool coins from the `farming` module. 
- The pool coins are sent from a reserve account of a liquid farm to a farmer.
- The pool coins are sent to a reserve account in `liquidity` module, and the corresponding coins are withdrawn to the farmer.

### MsgPlaceBid

- Bidding coins are sent to the `PayingReserveAddress` of an auction.

### MsgRefundBid

- Bidding coins are sent to a bidder account from the `PayingReserveAddress` of an auction.


The following events triggered by hooks cause state transition on the `bank`, `liquidty`, and `farming` modules.