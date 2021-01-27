package xlog

import (
	"github.com/pubgo/x/cnst"
	"github.com/pubgo/x/pkg/fileutil"
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xenv"
	"github.com/pubgo/x/xerror"
	"github.com/pubgo/x/xlog1/internal/hooks"
	"github.com/pubgo/x/xlog1/internal/zwriter"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"runtime"
)

func WithFormat(timeFormat string) func(*options) {
	return func(opt *options) {
		opt.TimeFormat = timeFormat
	}
}

func WithOutput(outputType string, filename string) func(*options) {
	return func(opt *options) {
		opt.OutputType = outputType
		opt.Filename = filename
	}
}

func WithRotate(maxSize int64, rotate bool) func(*options) {
	return func(opt *options) {
		opt.MaxSize = maxSize
		opt.Rotate = rotate
	}
}

type options struct {
	TimeFormat          string `toml:"time_format"`
	LogLevel            string `toml:"log_level"`
	TimestampFieldName  string `toml:"timestamp_field_name"`
	LevelFieldName      string `toml:"level_field_name"`
	MessageFieldName    string `toml:"message_field_name"`
	ErrorFieldName      string `toml:"error_field_name"`
	CallerFieldName     string `toml:"caller_field_name"`
	ErrorStackFieldName string `toml:"error_stack_field_name"`
	OutputType          string `toml:"output_type"`
	Filename            string `toml:"filename"`
	MaxSize             int64  `toml:"max_size"`
	Rotate              bool   `toml:"rotate"`
}

func appendExtra(ctx zerolog.Context) zerolog.Context {
	h, err := os.Hostname()
	xerror.PanicF(err, "get Hostname error")

	ctx = ctx.Int("cpus", runtime.NumCPU())

	if h != "" {
		ctx = ctx.Str("hostname", h)
	}

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

func GetLog(fn ...func(*options)) Log {
	opt := &options{}
	for _, f := range fn {
		f(opt)
	}

	if opt.TimestampFieldName == "" {
		opt.TimestampFieldName = "time"
	}
	if opt.LevelFieldName == "" {
		opt.LevelFieldName = "level"
	}
	if opt.TimeFormat == "" {
		opt.TimeFormat = cnst.DateTimeFormat
	}
	if opt.ErrorFieldName == "" {
		opt.ErrorFieldName = "error"
	}
	if opt.CallerFieldName == "" {
		opt.CallerFieldName = "caller"
	}
	if opt.ErrorStackFieldName == "" {
		opt.ErrorStackFieldName = "stack"
	}
	if opt.MessageFieldName == "" {
		opt.MessageFieldName = "msg"
	}

	zerolog.LevelFieldName = opt.LevelFieldName
	zerolog.TimestampFieldName = opt.TimestampFieldName
	zerolog.MessageFieldName = opt.MessageFieldName
	zerolog.ErrorFieldName = opt.ErrorFieldName
	zerolog.CallerFieldName = opt.CallerFieldName
	zerolog.ErrorStackFieldName = opt.ErrorStackFieldName

	if _l := xenv.GetEnv("log_level", "log.level"); _l != "" {
		opt.LogLevel = _l
	}
	if opt.LogLevel == "" {
		opt.LogLevel = zerolog.DebugLevel.String()
	}

	if !_Match(opt.LogLevel) {
		panic("log level is not match")
	}

	zerolog.SetGlobalLevel(xerror.PanicErr(zerolog.ParseLevel(opt.LogLevel)).(zerolog.Level))

	w := zwriter.NewZWriter()
	if opt.OutputType == "console" {
		w.Append(zerolog.NewConsoleWriter())
	} else {
		_home := xconfig.HomeDir()
		xerror.PanicF(fileutil.IsNotExistMkDir(filepath.Join(_home, "logs")), "create logs failed")
		w.Append(zwriter.NewFileWriter(filepath.Join(_home, "logs", opt.Filename)))
	}

	_log := appendExtra(zerolog.New(w).With().Timestamp().Caller()).Logger()
	_log = _log.Hook(hooks.HH)
	return _log
}
