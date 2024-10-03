package rpchandlers

import (
	"github.com/arumanetwork/anumad/infrastructure/logger"
	"github.com/arumanetwork/anumad/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
