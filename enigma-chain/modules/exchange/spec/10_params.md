---
sidebar_position: 11
title: Parameters
---

# Parameters

The exchange module contains the following parameters:

| Key                               | Type     | Example      |
| --------------------------------- | -------- | ------------ |
| SpotMarketInstantListingFee       | sdk.Coin | 100fury      |
| DerivativeMarketInstantListingFee | sdk.Coin | 1000fury      |
| DefaultSpotMakerFeeRate           | sdk.Dec  | 0.1%         |
| DefaultSpotTakerFeeRate           | sdk.Dec  | 0.2%         |
| DefaultDerivativeMakerFeeRate     | sdk.Dec  | 0.1%         |
| DefaultDerivativeTakerFeeRate     | sdk.Dec  | 0.2%         |
| DefaultInitialMarginRatio         | sdk.Dec  | 5%           |
| DefaultMaintenanceMarginRatio     | sdk.Dec  | 2%           |
| DefaultFundingInterval            | int64    | 3600         |
| FundingMultiple                   | int64    | 3600         |
| RelayerFeeShareRate               | sdk.Dec  | 40%          |
| DefaultHourlyFundingRateCap       | sdk.Dec  | 0.0625%      |
| DefaultHourlyInterestRate         | sdk.Dec  | 0.000416666% |
| MaxDerivativeOrderSideCount       | int64    | 20           |