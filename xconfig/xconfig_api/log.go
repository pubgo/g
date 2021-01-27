package xconfig_api

import (
	"github.com/pubgo/x/logs"
	"github.com/pubgo/x/xconfig/xconfig_log"
	"github.com/pubgo/x/xdi"
)

var logger = logs.DebugLog("pkg", "web")

func init() {
	xdi.InitInvoke(func(log xconfig_log.Log) {
		logger = log.With().Str("pkg", "web").Logger()
	})
}
