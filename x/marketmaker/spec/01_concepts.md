<!-- order: 1 -->

 # Concepts

## MarketMaker Module

The MarketMaker module implement a decentralized

## Market Maker

Market makers provide liquidity to the exchange by placing orders which are not executed immediately. These orders make it easy for traders to instantly buy or sell when the condition is met.

## Market Maker Incentives

Market makers can earn CRE as incentives for providing liquidity in a given month. These incentives are calculated based on a formula rewarding a combination of spread, depth and uptime for participating markets. The amount of incentives is determined by the relative share of score.

## Inclusion of Market Maker

Any wallets who satisfy minimum requirement of uptime and submit market maker intent for certain pairs, can receive incentives coming month. Market maker intent transaction is followed by 1,000 CRE deposit for spam prevention. New market maker inclusion is decided through governance proposal. Existing market makers who fail to maintain uptime requirement can be excluded by governance.

## Spread

The spread measures the transaction cost how much an investor loses resulting from opening and closing a position. Spread(or Bid-ask spread) is price difference between lowest ask and highest bid in a market and divide by mid-price(average price of bid and ask). Spread larger than MaxSpread for given market will not be recognized.

## Quantity(Depth)

Simple explanation for market depth is how much of the asset can be bought without causing slippage in price. If market depth is not enough to handle large amount of trade, market liquidity is insufficient despite its narrow spread. Quantity is measured from total amount of base coin denominated with quote coin. To appreciate 2-sided liquidity, quantity is calculated both bid and ask side and take the smaller of them. Quantity smaller than MinDepth for given market will not be recognized.

## Uptime

Uptime measure the availability of the liquidity. Especially in market turmoil, availability of liquidity is critical for market participants so that uptime is the most important factors in assessing the performance of market makers. Uptime is calculated as percentage of time a market maker is alive and quoting both bid and ask orders. Those orders should meet certain criteria to be recognized as valid orders.

## Market Maker Order Type

Market makers are provided a special order type so that they can place and cancel market making orders easily at once. Market maker order type supports to place up to 10 orders in comparable amount. This specific order type can relieve market makers’ risk in providing liquidity under various of conditions. Moreover, traders can enjoy a better trading environment through quality liquidity.