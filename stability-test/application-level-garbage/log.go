package main

import (
	"github.com/arumanetwork/anumad/infrastructure/logger"
	"github.com/arumanetwork/anumad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("APLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
