package types

// Event types for the farming module.
const (
	EventTypeCreateLiquidFarm = "create_liquid_farm"
	EventTypeDeposit          = "deposit"
	EventTypeCancel           = "cancel"
	EventTypeWithdraw         = "withdraw"

	AttributeKeyLiquidFarmId = "liquid_farm_id"
	AttributeKeyAuctionId    = "auction_id"
)
