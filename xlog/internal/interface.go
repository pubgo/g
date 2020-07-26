package internal

import "go.uber.org/zap"

type Field = zap.Field

type XLog interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	DPanic(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Named(s string) XLog
	With(fields ...Field) XLog
}
