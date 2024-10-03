package rpchandlers

import (
	"github.com/arumanetwork/anumad/app/appmessage"
	"github.com/arumanetwork/anumad/app/rpc/rpccontext"
	"github.com/arumanetwork/anumad/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
