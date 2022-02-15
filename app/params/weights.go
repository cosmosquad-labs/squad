package params

// Default simulation operation weights for messages and gov proposals.
const (
	DefaultWeightMsgCreateFixedAmountPlan int = 10
	DefaultWeightMsgCreateRatioPlan       int = 10
	DefaultWeightMsgStake                 int = 85
	DefaultWeightMsgUnstake               int = 30
	DefaultWeightMsgHarvest               int = 30

	DefaultWeightMsgCreatePair      int = 10
	DefaultWeightMsgCreatePool      int = 15
	DefaultWeightMsgDeposit         int = 20
	DefaultWeightMsgWithdraw        int = 20
	DefaultWeightMsgLimitOrder      int = 80
	DefaultWeightMsgMarketOrder     int = 60
	DefaultWeightMsgCancelOrder     int = 20
	DefaultWeightMsgCancelAllOrders int = 20

	DefaultWeightAddPublicPlanProposal    int = 5
	DefaultWeightUpdatePublicPlanProposal int = 5
	DefaultWeightDeletePublicPlanProposal int = 5

	DefaultWeightMsgLiquidStake   int = 80
	DefaultWeightMsgLiquidUnstake int = 30

	DefaultWeightAddWhitelistValidatorsProposal    int = 50
	DefaultWeightUpdateWhitelistValidatorsProposal int = 5
	DefaultWeightDeleteWhitelistValidatorsProposal int = 5
)
