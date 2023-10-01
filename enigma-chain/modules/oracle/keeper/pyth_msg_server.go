package keeper

import (
	"context"

	"github.com/EnigmasLab/metrics"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/EnigmasLab/enigma-core/enigma-chain/modules/oracle/types"
)

type PythMsgServer struct {
	Keeper
	svcTags metrics.Tags
}

// NewPythMsgServerImpl returns an implementation of the Pyth MsgServer interface for the provided Keeper for Pyth oracle functions.
func NewPythMsgServerImpl(keeper Keeper) PythMsgServer {
	return PythMsgServer{
		Keeper: keeper,
		svcTags: metrics.Tags{
			"svc": "pyth_msg_h",
		},
	}
}

func (k Keeper) RelayPythPrices(goCtx context.Context, msg *types.MsgRelayPythPrices) (*types.MsgRelayPythPricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pythContract, err := sdk.AccAddressFromBech32(k.GetParams(ctx).PythContract)
	if err != nil {
		return nil, types.ErrPythContractNotFound
	}

	if !pythContract.Equals(sdk.MustAccAddressFromBech32(msg.Sender)) {
		return nil, types.ErrUnauthorizedPythPriceRelay
	}

	k.ProcessPythPriceAttestations(ctx, msg.PriceAttestations)
	return &types.MsgRelayPythPricesResponse{}, nil
}
