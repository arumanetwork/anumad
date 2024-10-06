package rpcclient

import (
	"github.com/arumanetwork/anumad/infrastructure/logger"
	"github.com/arumanetwork/anumad/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
