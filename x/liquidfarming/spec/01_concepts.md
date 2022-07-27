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
In order to exchange the rewards coin(s) into the pool coin to be additionally staked for farming, the module creates an auction to sell the rewards that will be received at the end of the epoch. Note that the exact amount of the rewards being auctioned is not determined when the auction is created, but will be determined when the auction ends.
The amount of the rewards depends on the total amount of staked pool coins and the `liquidfarming` module’s staked pool coin, which can be varied during the epoch.
Therefore, a bidder to place a bid for the auction should be aware of this uncertainty of the rewards amount.

## Bidding for Auction and Winning Bid

A bidder can place a bid with the pool coin, which is the paying coin of the auction. A bidder only can place a single bid per auction of a liquid farm.
The bid amount of the pool coin must be higher than the current winning bid amount that is the highest bid amount of the auction. The bidder placing the bid with the highest amount of the pool coin becomes the winner of the auction and will takes all the accumulated rewards amount at the end of the auction.
