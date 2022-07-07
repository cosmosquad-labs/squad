package types

// Event types for the farming module.
const (
	EventTypeDeposit  = "deposit"
	EventTypeCancel   = "cancel"
	EventTypeWithdraw = "withdraw"

	AttributeKeyPoolId           = "pool_id"
	AttributeKeyDepositRequestId = "deposit_request_id"
	AttributeKeyDepositor        = "depositor"
	AttributeKeyDepositCoin      = "deposit_coin"
	AttributeKeyWithdrawer       = "withdrawer"
	AttributeKeyWithdrawingCoin  = "withdrawing_coin"
	AttributeKeyWithdrawnCoin    = "withdraw_coin"
	AttributeKeyAuctionId        = "auction_id"
)
