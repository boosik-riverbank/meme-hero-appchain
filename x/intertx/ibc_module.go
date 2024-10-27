package intertx

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	"github.com/cosmos/interchain-security/v6/x/intertx/keeper"
	"github.com/cosmos/interchain-security/v6/x/intertx/types"
	osmosistypes "github.com/osmosis-labs/osmosis/v26/x/gamm/pool-models/balancer"
	"strconv"
	"strings"
)

var _ porttypes.IBCModule = IBCModule{}

// IBCModule implements the ICS26 interface for interchain accounts controller chains
type IBCModule struct {
	keeper keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k keeper.Keeper) IBCModule {
	return IBCModule{
		keeper: k,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	im.keeper.Logger.Info("[OnChanOpenInit]")

	return version, nil
}

// OnChanOpenTry implements the IBCModule interface
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	return "", nil
}

// OnChanOpenAck implements the IBCModule interface
func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID string,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	im.keeper.Logger.Info(fmt.Sprintf("[OnChanOpenAck] portId=%v channelId=%v", portID, channelID))
	connId, err := im.keeper.GetConnectionId(ctx, portID)
	if err != nil {
		return err
	}

	hostAddr, ok := im.keeper.IcaControllerKeeper.GetInterchainAccountAddress(ctx, connId, portID)
	if !ok {
		return fmt.Errorf("cannot find host address")
	}

	butler, err := im.keeper.GetButlerPortId(ctx, portID)
	if err != nil {
		return fmt.Errorf("cannot find butler for given port")
	}

	butler.HostAddress = hostAddr
	butler.IcaConnectionInfo.ChannelId = channelID

	err = im.keeper.UpdateIcaAccountInfoToButler(ctx, butler.TargetChainId, hostAddr, channelID)
	return err
}

// OnChanOpenConfirm implements the IBCModule interface
func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnChanCloseInit implements the IBCModule interface
func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	panic("UNIMPLEMENTED")
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnRecvPacket implements the IBCModule interface. A successful acknowledgement
// is returned if the packet data is successfully decoded and the receive application
// logic returns without error.
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	return channeltypes.NewErrorAcknowledgement(fmt.Errorf("cannot receive packet via interchain accounts authentication module"))
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	im.keeper.Logger.Info("[OnAcknowledgementPacket]", packet.String(), string(acknowledgement))
	var ack channeltypes.Acknowledgement
	if err := channeltypes.SubModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return fmt.Errorf("cannot unmarshal ICS-27 packet acknowledgement: %v", err)
	}

	var transferPacketData transfertypes.FungibleTokenPacketData
	err := channeltypes.SubModuleCdc.UnmarshalJSON(packet.Data, &transferPacketData)
	if err == nil {
		splited := strings.Split(transferPacketData.Memo, ":")
		id, err := strconv.ParseUint(splited[0], 10, 64)
		if err != nil {
			return err
		}
		memo := splited[1]
		if memo == "DEX_LISTING_TOKEN" {
			im.keeper.Logger.Info("Token is transferred")
			senderAcc, err := sdk.AccAddressFromBech32(transferPacketData.Sender)
			if err != nil {
				return err
			}

			queue, err := im.keeper.GetQueue(ctx, id)
			if err != nil {
				return err
			}

			err = im.keeper.TransferAssets(context.Background(), id, senderAcc, transferPacketData.Receiver, *queue.Pair, types.MEMO_TRANSFER_PAIR)
			if err != nil {
				return err
			}
		} else if memo == "DEX_LISTING_PAIR" {
			im.keeper.Logger.Info("Pair is transferred")
		}

		return fmt.Errorf("unknown transfer packet")
	}

	var balancerPoolCreatePacketData osmosistypes.MsgCreateBalancerPool
	err = channeltypes.SubModuleCdc.UnmarshalJSON(packet.Data, &balancerPoolCreatePacketData)
	if err == nil {
		return nil
	}

	return fmt.Errorf("unknown ibc packet")
	//txMsgData := &sdk.TxMsgData{}
	//if err := proto.Unmarshal(ack.GetResult(), txMsgData); err != nil {
	//	return fmt.Errorf("cannot unmarshal ICS-27 tx message data: %v", err)
	//}
	//
	//if !ack.Success() {
	//	if err := im.keeper.HandleAckFail(ctx, packet); err != nil {
	//		return err
	//	}
	//	im.keeper.Logger.Error("ICA ack result is fail", "receive packet", packet)
	//}
	//
	//for _, data := range txMsgData.MsgResponses {
	//	switch data.TypeUrl {
	//	case "/osmosis/gamm/create-balancer-pool":
	//		// TODO : handle osmosis pool created msg
	//		break
	//
	//	default:
	//		break
	//	}
	//}
	//
	//return nil

	//switch len(txMsgData.Data) {
	//case 1: // Delegate, Undelegate, IcaWithdraw
	//	response, err := im.keeper.HandleAckMsgData(ctx, packet, txMsgData.Data[0])
	//	if err != nil {
	//		return err
	//	}
	//	im.keeper.Logger(ctx).Info("message response in ICS-27 packet response", "response", response)
	//
	//	return nil
	//case 2: // AutoStaking
	//	if txMsgData.Data[0].MsgType == sdk.MsgTypeURL(&distributiontype.MsgWithdrawDelegatorReward{}) &&
	//		txMsgData.Data[1].MsgType == sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}) {
	//		response, err := im.keeper.HandleAckMsgData(ctx, packet, txMsgData.Data[0])
	//		if err != nil {
	//			return err
	//		}
	//		im.keeper.Logger(ctx).Info("message response in ICS-27 packet response", "response", response)
	//	}
	//
	//	return nil
	//default:
	//	ctx.Logger().Debug("Unknown ICA msg", "receive packet", packet)
	//	return nil
	//}
}

// OnTimeoutPacket implements the IBCModule interface.
func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	return im.keeper.HandleAckFail(ctx, packet)
}

// NegotiateAppVersion implements the IBCModule interface
func (im IBCModule) NegotiateAppVersion(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionID string,
	portID string,
	counterparty channeltypes.Counterparty,
	proposedVersion string,
) (string, error) {
	return "", nil
}
