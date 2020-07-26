package xlog

import (
	"github.com/pubgo/x/xlog/internal"
	"github.com/pubgo/x/xlog/internal/log"
	"github.com/pubgo/x/xlog/xlog_config"
	"github.com/pubgo/xerror"
	"go.uber.org/zap"
)

type XLog = internal.XLog
type Field = zap.Field

func GetDevLog() XLog {
	return log.GetDevLog()
}

func GetLog() XLog {
	return log.GetLog()
}

func init() {
	// 初始化加载
	xerror.Exit(xlog_config.InitDevLog())
}
