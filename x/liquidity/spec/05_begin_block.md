<!-- order: 5 -->

Begin block operations for the liquidity module delete batch messages that were executed or ready to be deleted.

## **Delete batch messages**

- Delete `DepositRequest` and `WithdrawRequest` messages with status `RequestStatusSucceeded`
  or `RequestStatusFailed`
- Delete `SwapRequest` messages with status `SwapRequestStatusExecuted`, 	`SwapRequestStatusCanceled` or `SwapRequestStatusExpired`