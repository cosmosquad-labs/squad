<!-- order: 2 -->

# State

## LiquidValidator

LiquidValidator is a validator for liquid staking. Liquid validators are set from the whitelisted validators that are defined in global parameter `params.WhitelistedValidators`. Whitelisted validators must meet the active conditions. Otherwise they become inactive status; this results to no delegation shares and being removed from the active liquid validator set. This occurs during rebalancing at every begin block. 

```go
// LiquidValidator is a validator for liquid staking
type LiquidValidator struct {
   // operator_address defines the bech32-encoded address of the validator operator
   OperatorAddress string 
}
```

LiquidValidatorState contains the validator's state of status, weight, delegation shares, and liquid tokens. Each field has derived function that syncs with the state of the `staking` module. 

```go
// LiquidValidatorState is a liquid validator state
type LiquidValidatorState struct {
	// operator_address defines the bech32-encoded address of the validator operator
	OperatorAddress string
	// weight defines the weight that corresponds to liquid staking and unstaking amount
	Weight sdk.Int
	// status defines the liquid validator status
	Status ValidatorStatus
	// del_shares defines the delegation shares of the liquid validator
	DelShares sdk.Dec
	// liquid_tokens defines the token amount worth of delegaiton shares (slashing applied amount)
	LiquidTokens sdk.Int
}
```

LiquidValidators: `0xc0 | OperatorAddrLen (1 byte) | OperatorAddr -> ProtocolBuffer(liquidValidator)`

### Status

A liquid validator has the following status:

- `Active`: active validators are the whitelisted validators who are governed and elected by governance process. Delegators' delegations are equally distributed to all liquid validators. If they commit misbehavior, they can be slashed and delisted from the active validator set. Liquid stakers who unbond their delegation must wait for the duration of the `UnStakingTime`. It is a chain-specific parameter. During the unbonding period, they are still exposed to being slashed for any liquid validator’s misbehavior.

- `Inactive`: inactive validators are the ones that do not meet active conditions (see below section), but has delegation shares in `LiquidStakingProxyAcc`. Note that inactive liquid validator's `Weight` would end up zero.

```go
const (
	// VALIDATOR_STATUS_UNSPECIFIED defines the unspecified invalid status
	ValidatorStatusUnspecified ValidatorStatus = 0
	// VALIDATOR_STATUS_ACTIVE defines the active, valid status
	ValidatorStatusActive ValidatorStatus = 1
	// VALIDATOR_STATUS_INACTIVE defines the inactive, invalid status
	ValidatorStatusInactive ValidatorStatus = 2
)
```

### Active Conditions

- Must be defined in `params.WhitelistedValidators`
- Must exist in `staking` module; a liquid validator must not have nil delegation shares and tokens and they must have valid exchange rate.
- Must not be tombstoned

### Weight

The weight of a liquid validator is derived depending on their status:

- Active LiquidValidator: `TargetWeight` value defined in `params.WhitelistedValidators` by governance

- Inactive LiquidValidator: zero (`0`)

## NetAmount

NetAmount is the sum of the following items in `LiquidStakingProxyAcc`:

- Native token balance
- Token amount worth of delegation shares of all liquid validators
- Remaining rewards 
- Unbonding balance

`MintRate` is the total supply of bTokens divided by NetAmount. 
- `bTokenTotalSupply / NetAmount` 

Depending on the equation, the value transformation between native tokens and bTokens can be calculated as follows:

- NativeTokenToBToken : `nativeTokenAmount * bTokenTotalSupply / netAmount` with truncations
- BTokenToNativeToken : `bTokenAmount * netAmount / bTokenTotalSupply * (1-params.UnstakeFeeRate)` with truncations


### NetAmountState

NetAmountState is type for NetAmount. It is a raw data and mint rate that gets from several module state when they are needed. It is only used for calculation and query. They are not stored in KVStore.

```go
// NetAmountState is type for NetAmount
type NetAmountState struct {
	// mint_rate is bTokenTotalSupply / NetAmount
	MintRate sdk.Dec
	// btoken_total_supply returns the total supply of btoken(liquid_bond_denom)
	BtokenTotalSupply sdk.Int
	// net_amount is proxy account's native token balance + total liquid tokens + total remaining rewards + total unbonding balance
	NetAmount sdk.Dec
	// total_del_shares define the delegation shares of all liquid validators
	TotalDelShares sdk.Dec
	// total_liquid_tokens define the token amount worth of delegation shares of all liquid validator (slashing applied amount)
	TotalLiquidTokens sdk.Int
	// total_remaining_rewards define the sum of remaining rewards of proxy account by all liquid validators
	TotalRemainingRewards sdk.Dec
	// total_unbonding_balance define the unbonding balance of proxy account by all liquid validator (slashing applied amount)
	TotalUnbondingBalance sdk.Int
	// proxy_acc_balance define the balance of proxy account for the native token
	ProxyAccBalance sdk.Int
}
```
