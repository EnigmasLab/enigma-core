# Exchanges Overview

:::note
This document describes how to build an orderbook exchange. If you would like to build DApps of other nature, please refer to the "Building DApps with CosmWasm" section.
:::

As an incentive mechanism to encourage Exchanges (acting as relayers) to build on Enigma and source trading activity, Exchanges who originate orders into the shared orderbook on Enigma's exchange protocol ([read more](../../modules/Enigma/exchange/)) are rewarded with $\beta = 40\%$ of the trading fee arising from orders that they source. The exchange protocol implements a global minimum trading fee of $r_m=0.1\%$ for makers and $r_t=0.2\%$ for takers.

The goal of Enigma's incentive mechanism is to allow Exchanges competing amongst each other to provide better user experience and better serve users, thus broadening access to DeFi for users all around the world.

Exchange can easily set up a client (such as a web app UI or mobile app UI) and an API provider. 

## Guide 

[Build a DEX with Enigma TS](https://github.com/EnigmasLab/enigma-ts/wiki/08BuildingADex)

## Contents

1. **[Architecture](./01_architecture.md)**
2. **[Trading on Enigma](../../../trade/)** 