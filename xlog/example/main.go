package main

import (
	"github.com/pubgo/x/xlog"
	"github.com/pubgo/x/xlog/xlog_config"
)

var fields = xlog.FieldOf(
	xlog.String("key", "value"),
	xlog.Namespace("name"),
)
var log = xlog.GetDevLog().With(fields...)

func init() {
	log = xlog.GetLog().With(fields...)
}

func main() {
	log.Debug("hello",
		xlog.Skip(),
		xlog.Any("hss", "ss"),
	)

	log.Info("hello",
		xlog.Skip(),
		xlog.Any("hss", "ss"),
	)
}

func init() {
	xlog_config.InitFromJson()
}
