package log

import (
	"github.com/pubgo/x/xlog/internal"
	"github.com/pubgo/xerror"
	"go.uber.org/zap"
)

var _ internal.XLog = (*xlog)(nil)

type xlog struct {
	*zap.Logger
}

func (log *xlog) With(fields ...zap.Field) internal.XLog {
	return &xlog{log.Logger.With(fields...)}
}

func (log *xlog) Named(s string) internal.XLog {
	return &xlog{log.Logger.Named(s)}
}

func GetDevLog() internal.XLog {
	return &xlog{xerror.PanicErr(zap.NewDevelopment()).(*zap.Logger)}
}

var defaultLog = &xlog{}

func GetLog() internal.XLog {
	return defaultLog
}

func SetLog(lg *zap.Logger) {
	defaultLog.Logger = lg
}

func Sync(ll internal.XLog) error {
	return ll.(*xlog).Sync()
}
