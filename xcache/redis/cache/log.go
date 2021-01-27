package cache

import (
	"github.com/pubgo/x/logs"
	"github.com/rs/zerolog/log"
)

var logger = logs.DebugLog("pkg", "cache")

func InitLog() {
	logger = log.With().Str("pkg", "cache").Logger()
}
