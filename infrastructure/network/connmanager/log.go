package connmanager

import (
	"github.com/arumanetwork/anumad/infrastructure/logger"
	"github.com/arumanetwork/anumad/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
