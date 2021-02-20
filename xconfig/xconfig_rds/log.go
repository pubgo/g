package xconfig_rds

import (
	"github.com/pubgo/x/logs"
	"github.com/pubgo/x/xconfig/xconfig_log"
	"github.com/pubgo/x/xdi"
	"github.com/rs/zerolog"
	"xorm.io/core"
)

var logger = logs.DebugLog("pkg", "rds")

func init() {
	xdi.InitInvoke(func(log xconfig_log.Log) {
		logger = log.With().Str("pkg", "rds").Logger()
	})
}

//newSQLLogger create a new logger for xorm
func newSQLLogger() core.ILogger {
	return &SQLLogger{logger: &logger}
}

type SQLLogger struct {
	logger  *zerolog.Logger
	showSQL bool
}

func (l *SQLLogger) Debug(v ...interface{}) {
	l.logger.Print(v...) //TODO debug level
}

func (l *SQLLogger) Debugf(format string, v ...interface{}) {
	l.logger.Printf(format, v...) //TODO debug level
}

func (l *SQLLogger) Error(v ...interface{}) {
	l.logger.Print(v...) //TODO Error level
}

func (l *SQLLogger) Errorf(format string, v ...interface{}) {
	l.logger.Printf(format, v...) //TODO Error level
}

func (l *SQLLogger) Info(v ...interface{}) {
	l.logger.Print(v...) //TODO Info level
}

func (l *SQLLogger) Infof(format string, v ...interface{}) {
	l.logger.Printf(format, v...) //TODO Info level
}

func (l *SQLLogger) Warn(v ...interface{}) {
	l.logger.Print(v...) //TODO Warn level
}

func (l *SQLLogger) Warnf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *SQLLogger) Level() core.LogLevel {
	return core.LOG_DEBUG //TODO
}

func (l *SQLLogger) SetLevel(lvl core.LogLevel) {
	l.logger.Debug().Msgf("xorm.log.SetLevel %d", lvl)
}
func (l *SQLLogger) ShowSQL(show ...bool) {
	if len(show) > 0 {
		l.showSQL = show[0]
	}

}

func (l *SQLLogger) IsShowSQL() bool {
	return l.showSQL
}
