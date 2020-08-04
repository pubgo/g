package log

import (
	"github.com/pubgo/x/xlog/internal"
	"github.com/pubgo/xerror"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ internal.XLog = (*xlog)(nil)

type xlog struct {
	zl *zap.Logger
}

func (log *xlog) Debug(msg string, fields ...internal.Field) {
	log.zl.Debug(msg, fields...)
}

func (log *xlog) Info(msg string, fields ...internal.Field) {
	log.zl.Info(msg, fields...)
}

func (log *xlog) Warn(msg string, fields ...internal.Field) {
	log.zl.Warn(msg, fields...)
}

func (log *xlog) Error(msg string, fields ...internal.Field) {
	log.zl.Error(msg, fields...)
}

func (log *xlog) DPanic(msg string, fields ...internal.Field) {
	log.zl.DPanic(msg, fields...)
}

func (log *xlog) Panic(msg string, fields ...internal.Field) {
	log.zl.Panic(msg, fields...)
}

func (log *xlog) Fatal(msg string, fields ...internal.Field) {
	log.zl.Fatal(msg, fields...)
}

func (log *xlog) With(fields ...zap.Field) internal.XLog {
	return &xlog{log.zl.With(fields...)}
}

func (log *xlog) Named(s string) internal.XLog {
	return &xlog{log.zl.Named(s)}
}

func GetDevLog() internal.XLog {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return &xlog{xerror.PanicErr(cfg.Build()).(*zap.Logger)}
}

var defaultLog = &xlog{}

func GetLog() internal.XLog {
	return defaultLog
}

func SetLog(lg *zap.Logger) {
	defaultLog.zl = lg
}

func Sync(ll internal.XLog) error {
	return ll.(*xlog).zl.Sync()
}
