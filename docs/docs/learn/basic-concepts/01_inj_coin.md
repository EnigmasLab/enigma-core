---
sidebar_position: 1
description: FURY is Enigma’s native staking token.
title: FURY Coin
hide_title: true
---

# FURY

FURY is Enigma’s native staking token. Staked holders can govern and decide the future of the protocol.

## Base Denomination

FURY uses [Atto](https://en.wikipedia.org/wiki/Atto-) as the base denomination to maintain parity with Ethereum.

```
1 fury = 1×10⁻¹⁸ FURY
```

This matches Ethereum denomination of:

```
1 wei = 1x10⁻¹⁸ ETH
```

## Enigma Token Economics (Tokenomics)


## 1. Proof of Stake Security

The Enigma PoS blockchain is governed by the native FURY token.

Use cases for the token include but are not limited to the following: governance, staking, and dApp value capture.

The initial supply of FURY is 100,000,000 tokens and shall increase over time through block rewards.

The target FURY inflation is 7% at genesis and will decrease over time to 2%. Gradually, the total supply of FURY may be lower than the initial supply due to the deflationary mechanism detailed in the Exchange Fee Value Accrual section below.

## 2. Governance
The FURY token also serves as the native governance token for the Enigma Chain. 

FURY is used to govern all aspects of the chain including:
- Auction Module [Parameters](../../develop/modules/Enigma/auction/05_params.md)
- Exchange Module [Custom proposals](../../develop/modules/Enigma/exchange/06_proposals.md) and [Parameters](../../develop/modules/Enigma/exchange/10_params.md)
- Insurance Module [Parameters](../../develop/modules/Enigma/insurance/06_params.md)
- Oracle Module [Custom proposals](../../develop/modules/Enigma/oracle/04_proposals.md)
- Peggy Module [Parameters](../../develop/modules/Enigma/peggy/08_params.md)
- Wasmx Module [Parameters](../../develop/modules/Enigma/wasmx/05_params.md)
- Software upgrades
- Cosmos-SDK module parameters for the [auth](https://docs.cosmos.network/main/modules/auth#parameters), [bank](https://docs.cosmos.network/main/modules/bank), [crisis](https://docs.cosmos.network/main/modules/crisis), [distribution](https://docs.cosmos.network/main/modules/distribution), [gov](https://docs.cosmos.network/main/modules/gov), [mint](https://docs.cosmos.network/main/modules/mint), [slashing](https://docs.cosmos.network/main/modules/slashing), and [staking](https://docs.cosmos.network/main/modules/staking) modules.

Full details on the governance process can be found [here](https://blog.enigmaprotocol.com/enigma-governance-proposal-procedure).

## 3. Exchange dApps Incentives
The exchange protocol implements a global minimum trading fee of $r_m=0.1\%$ for makers and $r_t=0.2\%$ for takers.
As an incentive mechanism to encourage exchange dApps to source trading activity on the exchange protocol, exchange dApps who originate orders into the shared orderbook are rewarded with $\beta = 40\%$ of the trading fee arising from orders that they source.

## 4. Exchange Fee Value Accrual
The remaining $60\%$ of the exchange fee will undergo an on-chain buy-back-and-burn event where the aggregate exchange fee basket is auctioned off to the highest bidder in exchange for FURY. 
The FURY proceeds of this auction are then burned, thus deflating the total FURY supply. 

More details on the auction mechanism can be found [here](../../develop/modules/Enigma/auction/README.md). 

## 5. Collateral Backing for Derivatives
FURY will be utilized as an alternative to stablecoins as margin and collateral for Enigma's derivatives markets. 
In some derivative markets, FURY can also be used as collateral backing or insurance pool staking where stakers can earn interest on their locked tokens.