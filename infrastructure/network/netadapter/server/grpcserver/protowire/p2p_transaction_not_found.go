package protowire

import (
	"github.com/arumanetwork/anumad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *AnumadMessage_TransactionNotFound) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "AnumadMessage_TransactionNotFound is nil")
	}
	return x.TransactionNotFound.toAppMessage()
}

func (x *TransactionNotFoundMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "TransactionNotFoundMessage is nil")
	}
	id, err := x.Id.toDomain()
	if err != nil {
		return nil, err
	}
	return appmessage.NewMsgTransactionNotFound(id), nil
}

func (x *AnumadMessage_TransactionNotFound) fromAppMessage(msgTransactionsNotFound *appmessage.MsgTransactionNotFound) error {
	x.TransactionNotFound = &TransactionNotFoundMessage{
		Id: domainTransactionIDToProto(msgTransactionsNotFound.ID),
	}
	return nil
}
