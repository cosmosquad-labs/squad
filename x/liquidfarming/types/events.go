package types

// Event types for the farming module.
const (
	EventTypeFarm                = "farm"
	EventTypeUnfarm              = "unfarm"
	EventTypeUnfarmAndWithdraw   = "unfarm_and_withdraw"
	EventTypeCancelQueuedFarming = "cancel_queued_farming"
	EventTypePlaceBid            = "place_bid"
	EventTypeRefundBid           = "refund_bid"

	AttributeKeyPoolId        = "pool_id"
	AttributeKeyAuctionId     = "auction_id"
	AttributeKeyBidId         = "bid_id"
	AttributeKeyFarmer        = "farmer"
	AttributeKeyBidder        = "bidder"
	AttributeKeyFarmingCoin   = "farming_coin"
	AttributeKeyBiddingCoin   = "bidding_coin"
	AttributeKeyCanceledCoin  = "canceled_coin"
	AttributeKeyUnfarmingCoin = "unfarming_coin"
	AttributeKeyUnfarmedCoin  = "unfarmed_coin"
)
