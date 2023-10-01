package insurance

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/EnigmasLab/enigma-core/enigma-chain/modules/insurance/keeper"
	"github.com/InjectiveLabs/metrics"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	// call automatic withdraw keeper function
	err := k.WithdrawAllMaturedRedemptions(ctx)
	if err != nil {
		metrics.ReportFuncError(metrics.Tags{
			"svc": "insurance_abci",
		})
	}
}
