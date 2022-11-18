package ocr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/InjectiveLabs/injective-core/injective-chain/modules/ocr/keeper"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker runs on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {

	shouldProcessPayouts := ctx.BlockHeight()%int64(k.PayoutInterval(ctx)) == 0
	if !shouldProcessPayouts {
		return
	}

	feedConfigs := k.GetAllFeedConfigs(ctx)
	for _, feedConfig := range feedConfigs {
		k.ProcessRewardPayout(ctx, feedConfig)
	}
}

// EndBlocker runs on every end block
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
}
