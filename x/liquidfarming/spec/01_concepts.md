<!-- order: 1 -->
# Concepts

This is the concept for `Liquidfarming` module in Crescent Network.



## LiquidFarming Module

The `liquidfarming` module provides a functionality for farmers to have another option to use with their liquidity pool coins in Crescent Network. 

The module allows farmers to farm their pool coin and mint a synthetic version of the pool coin called LFCoin. 
Farmers can use the LFCoin to take a full advantage of Crescent functionality, such as Boost. 
On behalf of farmers, the module stakes their pool coin to the `farming` module and receives farming rewards for every epoch. 
The module provides auto-compounding of the rewards by going through an auction process, which results in the exchange of the farming rewards coin(s) into the pool coin.


## Farm

Once a user farms their pool coin, the user will receive LFCoin after a configurable queueing period.
The farm queueing period is configured as 1 day in Crescent Network. 
The following formula is used for an exchange rate of `LFCoinMint` when a user farms with `LPCoinFarm`.

$$LFCoinMint = \frac{LFCoinTotalSupply}{LPCoinTotalStaked} \times LPCoinFarm,$$

where `LFCoinTotalSupply` is not zero.
If `LFCoinTotalSupply` is zero, then the following formula is applied:

$$LFCoinMint = LPCoinFarm.$$

Note that the farm request can be cancelled by the user as long as minting LFCoin hasn’t occurred.

## Unfarm

When a user unfarms their LFCoin, the module burns the LFCoin and releases the corresponding amount of pool coin.
Unlike minting LFCoin that happens after a certain period from farming, burning the LFCoin occurs instantly when unfarming.
The following formula is used for an exchange rate of `LFCoinUnfarm`:

$$LPCoinUnfarm = \frac{LPCoinTotalStaked}{LFCoinTotalSupply} \times LFCoinBurn \times (1-fee).$$

## Farming Rewards and Auction

On behalf of users, the module stakes their pool coins and claims farming rewards for every epoch.
In order to exchange the rewards coin(s) into the pool coin to be additionally staked for farming, the module creates an auction.
Once the module claims the rewards, the module creates an auction to exchange the claimed rewards coin(s) into the pool coin to be additionally staked for farming.
The auction 