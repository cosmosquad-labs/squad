package types

// Event types for the farming module.
const (
	EventTypeFarm                = "farm"
	EventTypeCancelQueuedFarming = "cancel_queued_farming"
	EventTypeUnfarm              = "unfarm"

	AttributeKeyPoolId      = "pool_id"
	AttributeKeyFarmer      = "farmer"
	AttributeKeyFarmingCoin = "farming_coin"

	AttributeKeyCanceledCoin  = "canceled_coin"
	AttributeKeyUnfarmingCoin = "unfarming_coin"
	AttributeKeyUnfarmedCoin  = "unfarmed_coin"
	AttributeKeyAuctionId     = "auction_id"
)