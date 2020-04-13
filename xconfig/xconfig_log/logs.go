package xconfig_log

import (
	"github.com/pubgo/g/cnst"
	"github.com/pubgo/g/pkg/fileutil"
	"github.com/pubgo/g/xconfig"
	"github.com/pubgo/g/xconfig/xconfig_log/internal/hooks"
	"github.com/pubgo/g/xconfig/xconfig_log/internal/zwriter"
	"github.com/pubgo/g/xdi"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"runtime"
)

func appendExtra(ctx zerolog.Context) zerolog.Context {
	h, err := os.Hostname()
	xerror.PanicM(err, "get Hostname error")

	ctx = ctx.Int("cpus", runtime.NumCPU())

	//if t.Service != "" {
	//	ctx = ctx.Str("service", cfg.Service)
	//}

	if h != "" {
		ctx = ctx.Str("hostname", h)
	}

	//if t.Version != "" {
	//	ctx = ctx.Str("version", cfg.Version)
	//}

	return ctx
}

func _Match(e string) bool {
	return zerolog.DebugLevel.String() == e ||
		zerolog.ErrorLevel.String() == e ||
		zerolog.WarnLevel.String() == e ||
		zerolog.FatalLevel.String() == e ||
		zerolog.InfoLevel.String() == e ||
		zerolog.PanicLevel.String() == e
}

type Log = zerolog.Logger

func init() {
	xdi.InitProvide(func(cfg *xconfig.Config) (_log Log) {
		defer xerror.Assert()

		_cfg := cfg.Log
		if !_cfg.Enabled {
			return
		}

		if _cfg.TimestampFieldName == "" {
			_cfg.TimestampFieldName = "time"
		}
		if _cfg.LevelFieldName == "" {
			_cfg.LevelFieldName = "level"
		}
		if _cfg.TimeFormat == "" {
			_cfg.TimeFormat = cnst.DateTimeFormat
		}
		if _cfg.ErrorFieldName == "" {
			_cfg.ErrorFieldName = "error"
		}
		if _cfg.CallerFieldName == "" {
			_cfg.CallerFieldName = "caller"
		}
		if _cfg.ErrorStackFieldName == "" {
			_cfg.ErrorStackFieldName = "stack"
		}
		if _cfg.MessageFieldName == "" {
			_cfg.MessageFieldName = "msg"
		}

		zerolog.LevelFieldName = _cfg.LevelFieldName
		zerolog.TimestampFieldName = _cfg.TimestampFieldName
		zerolog.MessageFieldName = _cfg.MessageFieldName
		zerolog.ErrorFieldName = _cfg.ErrorFieldName
		zerolog.CallerFieldName = _cfg.CallerFieldName
		zerolog.ErrorStackFieldName = _cfg.ErrorStackFieldName

		if _l := xenv.GetEnv("log_level", "log.level"); _l != "" {
			_cfg.LogLevel = _l
		}
		if _cfg.LogLevel == "" {
			_cfg.LogLevel = zerolog.DebugLevel.String()
		}

		xerror.PanicT(!_Match(_cfg.LogLevel), "log level is not match")
		zerolog.SetGlobalLevel(xerror.PanicErr(zerolog.ParseLevel(_cfg.LogLevel)).(zerolog.Level))

		w := zwriter.NewZWriter()
		if _cfg.OutputType == "console" {
			w.Append(zerolog.NewConsoleWriter())
		} else {
			_home := xconfig.HomeDir()
			xerror.PanicM(fileutil.IsNotExistMkDir(filepath.Join(_home, "logs")), "create logs failed")
			w.Append(zwriter.NewFileWriter(filepath.Join(_home, "logs", _cfg.Filename)))
		}

		_log = appendExtra(zerolog.New(w).With().Timestamp().Caller()).Logger()
		_log = _log.Hook(hooks.HH)

		// set global logger
		log.Logger = _log
		return
	})
}
