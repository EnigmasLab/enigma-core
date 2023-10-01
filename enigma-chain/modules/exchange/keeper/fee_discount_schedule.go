package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/EnigmasLab/enigma-core/enigma-chain/modules/exchange/types"
	"github.com/EnigmasLab/metrics"
)

// GetFeeDiscountSchedule fetches the FeeDiscountSchedule.
func (k *Keeper) GetFeeDiscountSchedule(ctx sdk.Context) *types.FeeDiscountSchedule {
	defer metrics.ReportFuncCallAndTiming(k.svcTags)()

	store := k.getStore(ctx)
	bz := store.Get(types.FeeDiscountScheduleKey)
	if bz == nil {
		return nil
	}

	var campaignInfo types.FeeDiscountSchedule
	k.cdc.MustUnmarshal(bz, &campaignInfo)
	return &campaignInfo
}

// DeleteFeeDiscountSchedule deletes the FeeDiscountSchedule.
func (k *Keeper) DeleteFeeDiscountSchedule(ctx sdk.Context) {
	defer metrics.ReportFuncCallAndTiming(k.svcTags)()

	store := k.getStore(ctx)
	store.Delete(types.FeeDiscountScheduleKey)
}

// SetFeeDiscountSchedule sets the FeeDiscountSchedule.
func (k *Keeper) SetFeeDiscountSchedule(ctx sdk.Context, schedule *types.FeeDiscountSchedule) {
	defer metrics.ReportFuncCallAndTiming(k.svcTags)()

	store := k.getStore(ctx)
	bz := k.cdc.MustMarshal(schedule)
	store.Set(types.FeeDiscountScheduleKey, bz)

	k.SetFeeDiscountBucketCount(ctx, schedule.BucketCount)
	k.SetFeeDiscountBucketDuration(ctx, schedule.BucketDuration)

	// nolint:errcheck //ignored on purpose
	ctx.EventManager().EmitTypedEvent(&types.EventFeeDiscountSchedule{Schedule: schedule})
}