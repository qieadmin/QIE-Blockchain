

package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AttoQie defines the default coin denomination used in Qie in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Qie.
	AttoQie string = "aqie"

	// BaseDenomUnit defines the base denomination unit for Qie.
	// 1 qie = 1x10^{BaseDenomUnit} aqie
	BaseDenomUnit = 18

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewQieCoin is a utility function that returns an "aqie" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewQieCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(AttoQie, amount)
}

// NewQieDecCoin is a utility function that returns an "aqie" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewQieDecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoQie, amount)
}

// NewQieCoinInt64 is a utility function that returns an "aqie" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewQieCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoQie, amount)
}
