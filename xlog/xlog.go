package xlog

import (
	"go.uber.org/zap"

	"github.com/pubgo/x/xlog/internal"
	"github.com/pubgo/x/xlog/internal/log"
	"github.com/pubgo/x/xlog/xlog_config"
	"github.com/pubgo/xerror"
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
	xerror.Exit(xlog_config.InitFromConfig(xlog_config.NewDevConfig()))
}

func FieldOf(fields ...Field) []Field {
	return fields
}

func Sync(ll internal.XLog) error {
	return log.Sync(ll)
}
