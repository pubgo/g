package xlog

import (
	"fmt"
	"github.com/pubgo/x/xlog/internal"
	"github.com/pubgo/x/xlog/internal/log"
	"go.uber.org/zap"
)

func defaultLog() internal.XLog {
	return log.GetLog()
}

func Debug(msg string, fields ...internal.Field) {
	defaultLog().Debug(msg, fields...)
}

func DebugF(format string, a ...interface{}) {
	defaultLog().Debug(fmt.Sprintf(format, a...))
}

func Info(msg string, fields ...internal.Field) {
	defaultLog().Info(msg, fields...)
}

func InfoF(format string, a ...interface{}) {
	defaultLog().Info(fmt.Sprintf(format, a...))
}

func Warn(msg string, fields ...internal.Field) {
	defaultLog().Warn(msg, fields...)
}
func WarnF(format string, a ...interface{}) {
	defaultLog().Warn(fmt.Sprintf(format, a...))
}

func Error(msg string, fields ...internal.Field) {
	defaultLog().Error(msg, fields...)
}

func ErrorF(format string, a ...interface{}) {
	defaultLog().Error(fmt.Sprintf(format, a...))
}

func DPanic(msg string, fields ...internal.Field) {
	defaultLog().DPanic(msg, fields...)
}

func DPanicF(format string, a ...interface{}) {
	defaultLog().DPanic(fmt.Sprintf(format, a...))
}

func Panic(msg string, fields ...internal.Field) {
	defaultLog().Panic(msg, fields...)
}

func PanicF(format string, a ...interface{}) {
	defaultLog().Panic(fmt.Sprintf(format, a...))
}

func Fatal(msg string, fields ...internal.Field) {
	defaultLog().Fatal(msg, fields...)
}

func FatalF(format string, a ...interface{}) {
	defaultLog().Fatal(fmt.Sprintf(format, a...))
}

func With(fields ...zap.Field) internal.XLog {
	return defaultLog().With(fields...)
}

func Named(s string) internal.XLog {
	return defaultLog().Named(s).With(zap.Namespace(s))
}
