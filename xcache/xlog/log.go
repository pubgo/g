package xlog

import (
	"github.com/pubgo/x/logs"
	"github.com/rs/zerolog/log"
)

var Logger = logs.DebugLog("pkg", "xcache")

func InitLog() {
	Logger = log.With().Str("pkg", "xcache").Logger()
}
