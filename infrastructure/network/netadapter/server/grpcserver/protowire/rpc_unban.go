package protowire

import (
	"github.com/arumanetwork/anumad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *AnumadMessage_UnbanRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "AnumadMessage_UnbanRequest is nil")
	}
	return x.UnbanRequest.toAppMessage()
}

func (x *AnumadMessage_UnbanRequest) fromAppMessage(message *appmessage.UnbanRequestMessage) error {
	x.UnbanRequest = &UnbanRequestMessage{Ip: message.IP}
	return nil
}

func (x *UnbanRequestMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "UnbanRequestMessage is nil")
	}
	return &appmessage.UnbanRequestMessage{
		IP: x.Ip,
	}, nil
}

func (x *AnumadMessage_UnbanResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "AnumadMessage_UnbanResponse is nil")
	}
	return x.UnbanResponse.toAppMessage()
}

func (x *AnumadMessage_UnbanResponse) fromAppMessage(message *appmessage.UnbanResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.UnbanResponse = &UnbanResponseMessage{
		Error: err,
	}
	return nil
}

func (x *UnbanResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "UnbanResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.UnbanResponseMessage{
		Error: rpcErr,
	}, nil
}
