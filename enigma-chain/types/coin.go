package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// FURY defines the default coin denomination used in Ethermint in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Ethermint.
	EnigmaCoin string = "fury"

	// BaseDenomUnit defines the base denomination unit for Photons.
	// 1 photon = 1x10^{BaseDenomUnit} fury
	BaseDenomUnit = 18
)

// NewEnigmaCoin is a utility function that returns an "fury" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewEnigmaCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(EnigmaCoin, amount)
}

// NewEnigmaDecCoin is a utility function that returns an "fury" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewEnigmaDecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(EnigmaCoin, amount)
}

// NewEnigmaCoinInt64 is a utility function that returns an "fury" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewEnigmaCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(EnigmaCoin, amount)
}
