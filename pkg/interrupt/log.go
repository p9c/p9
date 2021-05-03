package interrupt

import (
	"github.com/p9c/p9/pkg/log"
	"github.com/p9c/p9/version"
)

var F, E, W, I, D, T = log.GetLogPrinterSet(log.AddLoggerSubsystem(version.PathBase))
