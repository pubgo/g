package xconfig_oss

import (
	"github.com/pubgo/g/logs"
	"github.com/pubgo/g/xconfig/xconfig_log"
	"github.com/pubgo/g/xdi"
)

var logger = logs.DebugLog("pkg", "oss")

func init() {
	xdi.InitInvoke(func(log xconfig_log.Log) {
		logger = log.With().Str("pkg", "oss").Logger()
	})
}
