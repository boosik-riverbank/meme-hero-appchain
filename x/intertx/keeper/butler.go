package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-security/v6/x/intertx/types"
)

func (k Keeper) CreateButler(ctx context.Context, msg *types.MsgCreateButler, portId string) (*types.MsgCreateButlerResponse, error) {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}

	exist, err := k.butlers.Has(ctx, msg.TargetChainId)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("target chain id is exist")
	}

	err = k.butlers.Set(ctx, msg.TargetChainId, types.Butler{
		TargetChainId: msg.TargetChainId,
		ChainTag:      msg.ChainTag,
		IcaConnectionInfo: &types.IcaConnectionInfo{
			ConnectionId: msg.ConnectionId,
			PortId:       portId,
		},
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateButlerResponse{}, nil
}

func (k Keeper) UpdateIcaAccountInfoToButler(ctx context.Context, targetChainId uint64, hostAddress string, channelId string) error {
	butler, err := k.butlers.Get(ctx, targetChainId)
	if err != nil {
		return err
	}

	butler.HostAddress = hostAddress

	err = k.butlers.Set(ctx, targetChainId, butler)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetButler(ctx context.Context, id uint64) (*types.Butler, error) {
	butler, err := k.butlers.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &butler, nil
}

func (k Keeper) IsExist(ctx context.Context, id uint64) (bool, error) {
	return k.butlers.Has(ctx, id)
}
