package bandoracle

import (
	"github.com/bandprotocol/bandchain-packet/obi"
	"github.com/bandprotocol/bandchain-packet/packet"
	"github.com/cosmonaut/oracle/x/bandoracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
)

// handleOraclePacket handles the result of the received BandChain oracles
// packet and saves the data into the KV database
func (am AppModule) handleOraclePacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
) (channeltypes.Acknowledgement, error) {
	var ack channeltypes.Acknowledgement
	var modulePacketData packet.OracleResponsePacketData
	if err := types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &modulePacketData); err != nil {
		return ack, nil
	}

	switch modulePacketData.GetClientID() {

	case types.FetchPriceClientIDKey:
		var fetchPriceResult types.FetchPriceResult
		if err := obi.Decode(modulePacketData.Result, &fetchPriceResult); err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err.Error())
			return ack, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
				"cannot decode the fetchPrice received packet")
		}
		am.keeper.SetFetchPriceResult(ctx, types.OracleRequestID(modulePacketData.RequestID), fetchPriceResult)

		// TODO: FetchPrice oracle data reception logic
		// this line is used by starport scaffolding # oracle/module/recv

	default:
		err := sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal,
			"oracle received packet not found: %s", modulePacketData.GetClientID())
		ack = channeltypes.NewErrorAcknowledgement(err.Error())
		return ack, err

	}
	ack = channeltypes.NewResultAcknowledgement(
		types.ModuleCdc.MustMarshalJSON(
			packet.NewOracleRequestPacketAcknowledgement(modulePacketData.RequestID),
		),
	)
	return ack, nil
}

// handleOracleAcknowledgment handles the acknowledgment result from the BandChain
// request and saves the request-id into the KV database
func (am AppModule) handleOracleAcknowledgment(
	ctx sdk.Context,
	ack channeltypes.Acknowledgement,
	modulePacket channeltypes.Packet,
) (*sdk.Result, error) {
	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		var oracleAck packet.OracleRequestPacketAcknowledgement
		err := types.ModuleCdc.UnmarshalJSON(resp.Result, &oracleAck)
		if err != nil {
			return nil, nil
		}

		var data packet.OracleRequestPacketData
		if err = types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &data); err != nil {
			return nil, nil
		}
		requestID := types.OracleRequestID(oracleAck.RequestID)

		switch data.GetClientID() {

		case types.FetchPriceClientIDKey:
			var fetchPriceData types.FetchPriceCallData
			if err = obi.Decode(data.GetCalldata(), &fetchPriceData); err != nil {
				return nil, sdkerrors.Wrap(err,
					"cannot decode the fetchPrice oracle acknowledgment packet")
			}
			am.keeper.SetLastFetchPriceID(ctx, requestID)
			return &sdk.Result{}, nil
			// this line is used by starport scaffolding # oracle/module/ack

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal,
				"oracle acknowledgment packet not found: %s", data.GetClientID())
		}
	}
	return nil, nil
}
