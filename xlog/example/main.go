package main

import (
	"github.com/pubgo/x/xlog"
	"github.com/pubgo/x/xlog/xlog_config"
	"github.com/pubgo/xerror"
)

var fields = xlog.FieldOf(
	xlog.String("key", "value"),
)
var log = xlog.GetDevLog().With(fields...)

func init() {
	//initCfgFromJson()
	initCfgFromJsonDebug()
	log = xlog.GetLog().
		Named("service").With(fields...).
		Named("hello").With(fields...).
		Named("world").With(fields...)
}

func main() {
	log.Debug("hello",
		xlog.Any("hss", "ss"),
	)

	log.Info("hello",
		xlog.Any("hss", "ss"),
	)

	log.Error("hello",
		xlog.Any("hss", "ss"),
	)

	log.Info("hello",
		xlog.Any("hss", "ss"),
	)
}

func initCfgFromJsonDebug() {
	cfg := `{
        "level": "debug",
        "development": true,
        "disableCaller": false,
        "disableStacktrace": false,
        "sampling": null,
        "encoding": "console",
        "encoderConfig": {
                "messageKey": "M",
                "levelKey": "L",
                "timeKey": "T",
                "nameKey": "N",
                "callerKey": "C",
                "stacktraceKey": "S",
                "lineEnding": "\n",
                "levelEncoder": "capitalColor",
                "timeEncoder": "iso8601",
                "durationEncoder": "string",
                "callerEncoder": "default",
                "nameEncoder": ""
        },
        "outputPaths": [
                "stderr"
        ],
        "errorOutputPaths": [
                "stderr"
        ],
        "initialFields": null
}`
	xerror.Exit(xlog_config.InitFromJson(
		[]byte(cfg),
	))
}


