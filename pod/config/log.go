package config

import (
	"github.com/p9c/p9/pkg/log"
	"github.com/p9c/p9/version"
)

var subsystem = log.AddLoggerSubsystem(version.PathBase)
var F, E, W, I, D, T log.LevelPrinter = log.GetLogPrinterSet(subsystem)
