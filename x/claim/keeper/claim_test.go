package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	utils "github.com/cosmosquad-labs/squad/types"
	"github.com/cosmosquad-labs/squad/x/claim"
	"github.com/cosmosquad-labs/squad/x/claim/types"
	"github.com/cosmosquad-labs/squad/x/liquidity"

	_ "github.com/stretchr/testify/suite"
)

func (s *KeeperTestSuite) TestClaim_DepositCondition() {
	// Create an airdrop
	sourceAddr := s.addr(0)
	airdrop := s.createAirdrop(
		1,
		sourceAddr,
		utils.ParseCoins("1000000000denom1"),
		[]types.ConditionType{
			types.ConditionTypeDeposit,
			types.ConditionTypeSwap,
			types.ConditionTypeLiquidStake,
			types.ConditionTypeVote,
		},
		s.ctx.BlockTime(),
		s.ctx.BlockTime().AddDate(0, 1, 0),
		true,
	)

	// Create a claim record
	recipient := s.addr(1)
	record := s.createClaimRecord(
		airdrop.Id,
		recipient,
		utils.ParseCoins("666666667denom1"),
		utils.ParseCoins("666666667denom1"),
		[]types.ConditionType{},
	)

	// Create a normal pair and pool
	creator := s.addr(2)
	s.createPair(creator, "denom3", "denom4", true)
	s.createPool(creator, 1, utils.ParseCoins("1000000denom3,1000000denom4"), true)

	// The recipient makes a deposit
	s.deposit(recipient, 1, utils.ParseCoins("500000denom3,500000denom4"), true)
	liquidity.EndBlocker(s.ctx, s.app.LiquidityKeeper)

	// Claim condition
	_, err := s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeDeposit))
	s.Require().NoError(err)

	r, found := s.keeper.GetClaimRecordByRecipient(s.ctx, airdrop.Id, record.GetRecipient())
	s.Require().True(found)
	s.Require().True(coinsEq(
		record.GetClaimableCoinsForCondition(airdrop.Conditions),
		sdk.NewCoins(s.getBalance(r.GetRecipient(), "denom1"))),
	)
	s.Require().Len(r.ClaimedConditions, 1)
	s.Require().Equal(types.ConditionTypeDeposit, r.ClaimedConditions[0])
}

func (s *KeeperTestSuite) TestClaim_SwapCondition() {
	// Create an airdrop
	sourceAddr := s.addr(0)
	airdrop := s.createAirdrop(
		1,
		sourceAddr,
		utils.ParseCoins("1000000000denom1"),
		[]types.ConditionType{
			types.ConditionTypeDeposit,
			types.ConditionTypeSwap,
			types.ConditionTypeLiquidStake,
			types.ConditionTypeVote,
		},
		s.ctx.BlockTime(),
		s.ctx.BlockTime().AddDate(0, 1, 0),
		true,
	)

	// Create a claim record
	recipient := s.addr(1)
	record := s.createClaimRecord(
		airdrop.Id,
		recipient,
		utils.ParseCoins("666666667denom1"),
		utils.ParseCoins("666666667denom1"),
		[]types.ConditionType{},
	)

	// Create a normal pool
	creator := s.addr(2)
	s.createPair(creator, "denom3", "denom4", true)
	s.createPool(creator, 1, utils.ParseCoins("1000000denom3,1000000denom4"), true)

	// The recipient makes a limit order
	s.sellLimitOrder(recipient, 1, utils.ParseDec("1.0"), sdk.NewInt(1000), 10, true)
	liquidity.EndBlocker(s.ctx, s.app.LiquidityKeeper)

	// Claim condition
	_, err := s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeSwap))
	s.Require().NoError(err)

	r, found := s.keeper.GetClaimRecordByRecipient(s.ctx, airdrop.Id, record.GetRecipient())
	s.Require().True(found)
	s.Require().True(coinsEq(
		record.GetClaimableCoinsForCondition(airdrop.Conditions),
		sdk.NewCoins(s.getBalance(r.GetRecipient(), "denom1"))),
	)
	s.Require().Len(r.ClaimedConditions, 1)
	s.Require().Equal(types.ConditionTypeSwap, r.ClaimedConditions[0])
}

func (s *KeeperTestSuite) TestClaim_LiquidStakeCondition() {
	// Create an airdrop
	sourceAddr := s.addr(0)
	airdrop := s.createAirdrop(
		1,
		sourceAddr,
		utils.ParseCoins("1000000000denom1"),
		[]types.ConditionType{
			types.ConditionTypeDeposit,
			types.ConditionTypeSwap,
			types.ConditionTypeLiquidStake,
			types.ConditionTypeVote,
		},
		s.ctx.BlockTime(),
		s.ctx.BlockTime().AddDate(0, 1, 0),
		true,
	)

	// Create a claim record
	recipient := s.addr(1)
	record := s.createClaimRecord(
		airdrop.Id,
		recipient,
		utils.ParseCoins("666666667denom1"),
		utils.ParseCoins("666666667denom1"),
		[]types.ConditionType{},
	)

	// Create whitelisted validators
	s.createWhitelistedValidators([]int64{1000000, 1000000, 1000000})

	// Make a liquid staking
	s.liquidStaking(recipient, sdk.NewInt(100_000_000), true)

	// Claim condition
	_, err := s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeLiquidStake))
	s.Require().NoError(err)

	r, found := s.keeper.GetClaimRecordByRecipient(s.ctx, airdrop.Id, record.GetRecipient())
	s.Require().True(found)
	s.Require().True(coinsEq(
		record.GetClaimableCoinsForCondition(airdrop.Conditions),
		sdk.NewCoins(s.getBalance(r.GetRecipient(), "denom1"))),
	)
	s.Require().Len(r.ClaimedConditions, 1)
	s.Require().Equal(types.ConditionTypeLiquidStake, r.ClaimedConditions[0])
}

func (s *KeeperTestSuite) TestClaim_VoteCondition() {
	// Create an airdrop
	sourceAddr := s.addr(0)
	airdrop := s.createAirdrop(
		1,
		sourceAddr,
		utils.ParseCoins("1000000000denom1"),
		[]types.ConditionType{
			types.ConditionTypeDeposit,
			types.ConditionTypeSwap,
			types.ConditionTypeLiquidStake,
			types.ConditionTypeVote,
		},
		s.ctx.BlockTime(),
		s.ctx.BlockTime().AddDate(0, 1, 0),
		true,
	)

	// Create a claim record
	recipient := s.addr(1)
	record := s.createClaimRecord(
		airdrop.Id,
		recipient,
		utils.ParseCoins("666666667denom1"),
		utils.ParseCoins("666666667denom1"),
		[]types.ConditionType{},
	)

	// Submit a governance proposal
	s.createTextProposal(sourceAddr, "Text", "Description")

	// Vote
	s.vote(recipient, 1, govtypes.OptionYes)

	// Claim condition
	_, err := s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeVote))
	s.Require().NoError(err)

	r, found := s.keeper.GetClaimRecordByRecipient(s.ctx, airdrop.Id, record.GetRecipient())
	s.Require().True(found)

	s.Require().True(coinsEq(
		record.GetClaimableCoinsForCondition(airdrop.Conditions),
		sdk.NewCoins(s.getBalance(r.GetRecipient(), "denom1"))),
	)
	s.Require().Len(r.ClaimedConditions, 1)
	s.Require().Equal(types.ConditionTypeVote, r.ClaimedConditions[0])
}

func (s *KeeperTestSuite) TestClaim_All() {
	// Create an airdrop
	sourceAddr := s.addr(0)
	airdrop := s.createAirdrop(
		1,
		sourceAddr,
		utils.ParseCoins("1000000000denom1"),
		[]types.ConditionType{
			types.ConditionTypeDeposit,
			types.ConditionTypeSwap,
			types.ConditionTypeLiquidStake,
			types.ConditionTypeVote,
		},
		s.ctx.BlockTime(),
		s.ctx.BlockTime().AddDate(0, 1, 0),
		true,
	)

	// Create a claim record
	recipient := s.addr(1)
	record := s.createClaimRecord(
		airdrop.Id,
		recipient,
		utils.ParseCoins("666666667denom1"),
		utils.ParseCoins("666666667denom1"),
		[]types.ConditionType{},
	)

	// Create a normal pool
	params := s.app.LiquidityKeeper.GetParams(s.ctx)
	creator := s.addr(2)
	s.createPair(creator, "denom3", "denom4", true)
	s.createPool(creator, 1, utils.ParseCoins("1000000denom3,1000000denom4"), true)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)
	s.Require().Equal(params.MinInitialPoolCoinSupply, s.getBalance(creator, pool.PoolCoinDenom).Amount)

	// The recipient makes a deposit
	s.deposit(recipient, pool.Id, utils.ParseCoins("500000denom3,500000denom4"), true)
	liquidity.EndBlocker(s.ctx, s.app.LiquidityKeeper)

	// The recipient makes a limit order
	s.sellLimitOrder(recipient, 1, utils.ParseDec("1.0"), sdk.NewInt(1000), 10, true)
	liquidity.EndBlocker(s.ctx, s.app.LiquidityKeeper)

	// Create whitelisted validators
	s.createWhitelistedValidators([]int64{1000000, 1000000, 1000000})

	// Make a liquid staking
	s.liquidStaking(recipient, sdk.NewInt(100_000_000), true)

	// Submit a governance proposal
	s.createTextProposal(sourceAddr, "Text", "Description")

	// Vote
	s.vote(recipient, 1, govtypes.OptionYes)

	// Claim deposit condition
	_, err := s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeDeposit))
	s.Require().NoError(err)

	// Claim swap condition
	_, err = s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeSwap))
	s.Require().NoError(err)

	// Claim liquid stake condition
	_, err = s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeLiquidStake))
	s.Require().NoError(err)

	// Claim vote condition
	_, err = s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeVote))
	s.Require().NoError(err)

	r, found := s.keeper.GetClaimRecordByRecipient(s.ctx, airdrop.Id, record.GetRecipient())
	s.Require().True(found)
	s.Require().True(coinsEq(
		r.InitialClaimableCoins,
		sdk.NewCoins(s.getBalance(r.GetRecipient(), "denom1"))),
	)
	s.Require().Len(r.ClaimedConditions, 4)
}

func (s *KeeperTestSuite) TestClaim_AlreadyClaimedCondition() {
	// Create an airdrop
	sourceAddr := s.addr(0)
	airdrop := s.createAirdrop(
		1,
		sourceAddr,
		utils.ParseCoins("1000000000denom1"),
		[]types.ConditionType{
			types.ConditionTypeDeposit,
			types.ConditionTypeSwap,
			types.ConditionTypeLiquidStake,
			types.ConditionTypeVote,
		},
		s.ctx.BlockTime(),
		s.ctx.BlockTime().AddDate(0, 1, 0),
		true,
	)

	// Create a claim record
	recipient := s.addr(1)
	s.createClaimRecord(
		airdrop.Id,
		recipient,
		utils.ParseCoins("666666667denom1"),
		utils.ParseCoins("666666667denom1"),
		[]types.ConditionType{},
	)

	// Create a normal pool
	creator := s.addr(2)
	s.createPair(creator, "denom3", "denom4", true)
	s.createPool(creator, 1, utils.ParseCoins("1000000denom3,1000000denom4"), true)

	// The recipient makes a deposit
	s.deposit(recipient, 1, utils.ParseCoins("500000denom3,500000denom4"), true)
	liquidity.EndBlocker(s.ctx, s.app.LiquidityKeeper)

	// Claim deposit condition
	_, err := s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeDeposit))
	s.Require().NoError(err)

	// Claim the already completed deposit condition
	_, err = s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeDeposit))
	s.Require().ErrorIs(err, types.ErrAlreadyClaimed)
}

func (s *KeeperTestSuite) TestClaim_All_TerminateAidrop() {
	// Create an airdrop
	sourceAddr := s.addr(0)
	airdrop := s.createAirdrop(
		1,
		sourceAddr,
		utils.ParseCoins("1000000000denom1"),
		[]types.ConditionType{
			types.ConditionTypeDeposit,
			types.ConditionTypeSwap,
			types.ConditionTypeLiquidStake,
			types.ConditionTypeVote,
		},
		s.ctx.BlockTime(),
		s.ctx.BlockTime().AddDate(0, 1, 0),
		true,
	)

	// Create a claim record
	recipient := s.addr(1)
	s.createClaimRecord(
		airdrop.Id,
		recipient,
		utils.ParseCoins("1000000000denom1"),
		utils.ParseCoins("1000000000denom1"),
		[]types.ConditionType{},
	)

	// Create a normal pool
	creator := s.addr(2)
	s.createPair(creator, "denom3", "denom4", true)
	s.createPool(creator, 1, utils.ParseCoins("1000000denom3,1000000denom4"), true)

	// The recipient makes a deposit
	s.deposit(recipient, 1, utils.ParseCoins("500000denom3,500000denom4"), true)
	liquidity.EndBlocker(s.ctx, s.app.LiquidityKeeper)

	// The recipient makes a limit order
	s.sellLimitOrder(recipient, 1, utils.ParseDec("1.0"), sdk.NewInt(1000), 10, true)
	liquidity.EndBlocker(s.ctx, s.app.LiquidityKeeper)

	// Create whitelisted validators
	s.createWhitelistedValidators([]int64{1000000, 1000000, 1000000})

	// Make a liquid staking
	s.liquidStaking(recipient, sdk.NewInt(100_000_000), true)

	// Submit a governance proposal
	s.createTextProposal(sourceAddr, "Text", "Description")

	// Vote
	s.vote(recipient, 1, govtypes.OptionYes)

	// Claim deposit condition
	_, err := s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeDeposit))
	s.Require().NoError(err)

	// Claim swap condition
	_, err = s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeSwap))
	s.Require().NoError(err)

	// Claim liquid stake condition
	_, err = s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeLiquidStake))
	s.Require().NoError(err)

	// Claim vote condition
	_, err = s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeVote))
	s.Require().NoError(err)

	// Terminate the airdrop
	s.ctx = s.ctx.WithBlockTime(airdrop.EndTime.AddDate(0, 0, 1))
	claim.EndBlocker(s.ctx, s.keeper)

	// Source account balances must be zero
	s.Require().True(s.getAllBalances(airdrop.GetSourceAddress()).IsZero())

	// Community pool must be zero
	feePool := s.app.DistrKeeper.GetFeePool(s.ctx)
	s.Require().True(feePool.CommunityPool.IsZero())
}

func (s *KeeperTestSuite) TestClaim_Partial_TerminatAirdrop() {
	// Create an airdrop
	sourceAddr := s.addr(0)
	airdrop := s.createAirdrop(
		1,
		sourceAddr,
		utils.ParseCoins("1000000000denom1"),
		[]types.ConditionType{
			types.ConditionTypeDeposit,
			types.ConditionTypeSwap,
			types.ConditionTypeLiquidStake,
			types.ConditionTypeVote,
		},
		s.ctx.BlockTime(),
		s.ctx.BlockTime().AddDate(0, 1, 0),
		true,
	)

	// Create a claim record
	recipient := s.addr(1)
	s.createClaimRecord(
		airdrop.Id,
		recipient,
		utils.ParseCoins("1000000000denom1"),
		utils.ParseCoins("1000000000denom1"),
		[]types.ConditionType{},
	)

	// Create a normal pool
	creator := s.addr(2)
	s.createPair(creator, "denom3", "denom4", true)
	s.createPool(creator, 1, utils.ParseCoins("1000000denom3,1000000denom4"), true)

	// The recipient makes a deposit
	s.deposit(recipient, 1, utils.ParseCoins("500000denom3,500000denom4"), true)
	liquidity.EndBlocker(s.ctx, s.app.LiquidityKeeper)

	// The recipient makes a limit order
	s.sellLimitOrder(recipient, 1, utils.ParseDec("1.0"), sdk.NewInt(1000), 10, true)
	liquidity.EndBlocker(s.ctx, s.app.LiquidityKeeper)

	// Create whitelisted validators
	s.createWhitelistedValidators([]int64{1000000, 1000000, 1000000})

	// Make a liquid staking
	s.liquidStaking(recipient, sdk.NewInt(100_000_000), true)

	// Submit a governance proposal
	s.createTextProposal(sourceAddr, "Text", "Description")

	// Vote
	s.vote(recipient, 1, govtypes.OptionYes)

	// Claim deposit condition
	_, err := s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeDeposit))
	s.Require().NoError(err)

	// Terminate the airdrop
	s.ctx = s.ctx.WithBlockTime(airdrop.EndTime.AddDate(0, 0, 1))
	claim.EndBlocker(s.ctx, s.keeper)

	// Claim swap condition
	// Must return ErrTerminatedAirdrop
	_, err = s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeSwap))
	s.Require().ErrorIs(err, types.ErrTerminatedAirdrop)

	// Source account balances must be zero
	s.Require().True(s.getAllBalances(airdrop.GetSourceAddress()).IsZero())

	// Community pool must have the remaining coins
	feePool := s.app.DistrKeeper.GetFeePool(s.ctx)
	s.Require().False(feePool.CommunityPool.IsZero())
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *KeeperTestSuite) TestClaim_Simulation() {
	airdropSourceAddress := s.addr(0)
	airdropTotalSupply := utils.ParseCoins("40000000000000stake")

	// Create an airdrop
	airdrop := s.createAirdrop(
		1,
		airdropSourceAddress,
		airdropTotalSupply, // 50 mil - 20% initial allocation
		[]types.ConditionType{
			types.ConditionTypeDeposit,
			types.ConditionTypeSwap,
			types.ConditionTypeLiquidStake,
			types.ConditionTypeVote,
		},
		s.ctx.BlockTime(),
		s.ctx.BlockTime().AddDate(0, 6, 0),
		true,
	)

	records := []types.ClaimRecord{
		{
			AirdropId:             airdrop.Id,
			Recipient:             s.addr(1).String(),
			InitialClaimableCoins: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 404783390)),
			ClaimableCoins:        sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 404783390)),
			ClaimedConditions:     []types.ConditionType{},
		},
		{
			AirdropId:             airdrop.Id,
			Recipient:             s.addr(2).String(),
			InitialClaimableCoins: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 28538190)),
			ClaimableCoins:        sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 28538190)),
			ClaimedConditions:     []types.ConditionType{},
		},
		{
			AirdropId:             airdrop.Id,
			Recipient:             s.addr(3).String(),
			InitialClaimableCoins: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 239263740)),
			ClaimableCoins:        sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 239263740)),
			ClaimedConditions:     []types.ConditionType{},
		},
		{
			AirdropId:             airdrop.Id,
			Recipient:             s.addr(4).String(),
			InitialClaimableCoins: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 104088579)),
			ClaimableCoins:        sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 104088579)),
			ClaimedConditions:     []types.ConditionType{},
		},
		{
			AirdropId:             airdrop.Id,
			Recipient:             s.addr(5).String(),
			InitialClaimableCoins: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 478195171)),
			ClaimableCoins:        sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 478195171)),
			ClaimedConditions:     []types.ConditionType{},
		},
		{
			AirdropId:             airdrop.Id,
			Recipient:             s.addr(6).String(),
			InitialClaimableCoins: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 230378299)),
			ClaimableCoins:        sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 230378299)),
			ClaimedConditions:     []types.ConditionType{},
		},
		{
			AirdropId:             airdrop.Id,
			Recipient:             s.addr(7).String(),
			InitialClaimableCoins: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 88602486)),
			ClaimableCoins:        sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 88602486)),
			ClaimedConditions:     []types.ConditionType{},
		},
		{
			AirdropId:             airdrop.Id,
			Recipient:             s.addr(8).String(),
			InitialClaimableCoins: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 26939)),
			ClaimableCoins:        sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 26939)),
			ClaimedConditions:     []types.ConditionType{},
		},
		{
			AirdropId:             airdrop.Id,
			Recipient:             s.addr(9).String(),
			InitialClaimableCoins: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 142181104)),
			ClaimableCoins:        sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 142181104)),
			ClaimedConditions:     []types.ConditionType{},
		},
		{
			AirdropId:             airdrop.Id,
			Recipient:             s.addr(10).String(),
			InitialClaimableCoins: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 161962396)),
			ClaimableCoins:        sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 161962396)),
			ClaimedConditions:     []types.ConditionType{},
		},
	}

	// Set all claim records
	var totalInitialClaimableCoins sdk.Coins
	for _, r := range records {
		s.createClaimRecord(r.AirdropId, r.GetRecipient(), r.InitialClaimableCoins, r.ClaimableCoins, r.ClaimedConditions)

		totalInitialClaimableCoins = totalInitialClaimableCoins.Add(r.InitialClaimableCoins...)
	}

	// Create a normal pool
	masterAddr := s.addr(20)
	params := s.app.LiquidityKeeper.GetParams(s.ctx)
	s.createPair(masterAddr, "denom3", "denom4", true)
	s.createPool(masterAddr, 1, utils.ParseCoins("1000000denom3,1000000denom4"), true)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)
	s.Require().Equal(params.MinInitialPoolCoinSupply, s.getBalance(masterAddr, pool.PoolCoinDenom).Amount)

	// Create whitelisted validators
	s.createWhitelistedValidators([]int64{1000000, 1000000, 1000000})

	// Submit a governance proposal
	s.createTextProposal(masterAddr, "Text", "Description")

	// All recipient execute their conditions
	for i := 1; i <= len(records); i++ {
		recipient := s.addr(i)

		// The recipient makes a deposit
		s.deposit(recipient, pool.Id, utils.ParseCoins("500000denom3,500000denom4"), true)
		liquidity.EndBlocker(s.ctx, s.app.LiquidityKeeper)

		// The recipient makes a limit order
		s.sellLimitOrder(recipient, 1, utils.ParseDec("1.0"), sdk.NewInt(1000), 10, true)
		liquidity.EndBlocker(s.ctx, s.app.LiquidityKeeper)

		// Make a liquid staking
		s.liquidStaking(recipient, sdk.NewInt(100_000_000), true)

		// Vote
		s.vote(recipient, 1, govtypes.OptionYes)
	}

	fmt.Println("+--------Must be empty--------+")
	for i := 1; i <= len(records); i++ {
		recipient := s.addr(i)
		fmt.Printf("Recipient Balance #%d: %s\n", i, s.getBalance(recipient, sdk.DefaultBondDenom).String())
	}
	fmt.Println("")

	// Claim all conditions
	for i := 1; i <= len(records); i++ {
		recipient := s.addr(i)

		// Claim deposit condition
		_, err := s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeDeposit))
		s.Require().NoError(err)

		// Claim swap condition
		_, err = s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeSwap))
		s.Require().NoError(err)

		// Claim liquid stake condition
		_, err = s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeLiquidStake))
		s.Require().NoError(err)

		// Claim vote condition
		_, err = s.keeper.Claim(s.ctx, types.NewMsgClaim(airdrop.Id, recipient, types.ConditionTypeVote))
		s.Require().NoError(err)
	}

	fmt.Println("+--------Must be all claimed--------+")
	for i := 1; i <= len(records); i++ {
		recipient := s.addr(i)
		fmt.Printf("Recipient Balance #%d: %s\n", i, s.getBalance(recipient, sdk.DefaultBondDenom).String())
	}
	fmt.Println("")

	// Forcefully end the airdrop event
	s.ctx = s.ctx.WithBlockTime(airdrop.EndTime)
	claim.EndBlocker(s.ctx, s.app.ClaimKeeper)

	// Source account balances must be zero
	s.Require().True(s.getAllBalances(airdrop.GetSourceAddress()).IsZero())

	// Community pool must have received the remaining coins
	feePool := s.app.DistrKeeper.GetFeePool(s.ctx)
	feePoolAmt := feePool.CommunityPool.AmountOf(sdk.DefaultBondDenom)
	s.Require().False(feePool.CommunityPool.IsZero())

	// The amount must be the same
	remainingCoins := airdropTotalSupply.Sub(totalInitialClaimableCoins)
	remainingAmt := remainingCoins.AmountOf(sdk.DefaultBondDenom)
	s.Require().True(remainingAmt.ToDec().Equal(feePoolAmt))
	fmt.Println("feePoolAmt: ", feePoolAmt)
	fmt.Println("remainingAmt: ", remainingAmt)
}
