package xlog_config

import (
	"fmt"
	"github.com/pubgo/x/xlog/internal/writer/rotate"
	"github.com/pubgo/xerror"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/url"
	"strconv"
	"time"
)

const defaultKey = "default"

var levelEncoder = map[string]zapcore.LevelEncoder{
	"capital":      zapcore.CapitalLevelEncoder,
	"capitalColor": zapcore.CapitalColorLevelEncoder,
	"color":        zapcore.LowercaseColorLevelEncoder,
	defaultKey:     zapcore.LowercaseLevelEncoder,
}
var timeEncoder = map[string]zapcore.TimeEncoder{
	"rfc3339": _RFC3339MilliTimeEncoder,
	"RFC3339": _RFC3339MilliTimeEncoder,
	//"rfc3339":     zapcore.RFC3339TimeEncoder,
	//"RFC3339":     zapcore.RFC3339TimeEncoder,
	"rfc3339nano": zapcore.RFC3339NanoTimeEncoder,
	"RFC3339Nano": zapcore.RFC3339NanoTimeEncoder,
	"iso8601":     zapcore.ISO8601TimeEncoder,
	"ISO8601":     zapcore.ISO8601TimeEncoder,
	"millis":      zapcore.EpochMillisTimeEncoder,
	"nanos":       zapcore.EpochNanosTimeEncoder,
	defaultKey:    zapcore.EpochTimeEncoder,
}
var durationEncoder = map[string]zapcore.DurationEncoder{
	"string":   zapcore.StringDurationEncoder,
	"nanos":    zapcore.NanosDurationEncoder,
	defaultKey: zapcore.SecondsDurationEncoder,
}
var callerEncoder = map[string]zapcore.CallerEncoder{
	"full":     zapcore.FullCallerEncoder,
	defaultKey: zapcore.ShortCallerEncoder,
}
var nameEncoder = map[string]zapcore.NameEncoder{
	"full":     zapcore.FullNameEncoder,
	defaultKey: zapcore.FullNameEncoder,
}

func _RFC3339MilliTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, "2006-01-02T15:04:05.000Z07:00")
		return
	}

	enc.AppendString(t.Format("2006-01-02T15:04:05.000Z07:00"))
}

func init() {
	//zap.RegisterEncoder()
	// rotate:///hello.go
	xerror.Exit(zap.RegisterSink("rotate", func(u *url.URL) (zap.Sink, error) {
		if u.User != nil {
			return nil, fmt.Errorf("user and password not allowed with file URLs: got %v", u)
		}
		if u.Fragment != "" {
			return nil, fmt.Errorf("fragments not allowed with file URLs: got %v", u)
		}
		if u.RawQuery != "" {
			return nil, fmt.Errorf("query parameters not allowed with file URLs: got %v", u)
		}
		// Error messages are better if we check hostname and port separately.
		if u.Port() != "" {
			return nil, fmt.Errorf("ports not allowed with file URLs: got %v", u)
		}
		if hn := u.Hostname(); hn != "" && hn != "localhost" {
			return nil, fmt.Errorf("file URLs must leave host empty or use localhost: got %v", u)
		}

		query := u.Query()
		var cfg = rotate.NewWriterConfig()
		for k := range query {
			v := query.Get(k)
			switch k {
			case "dir":
				cfg.Dir = v
			case "sub":
				cfg.Sub = v
			case "name":
				cfg.Filename = v
			case "age":
				cfg.Age = xerror.PanicErr(time.ParseDuration(v)).(time.Duration)
			case "dur":
				cfg.Duration = xerror.PanicErr(time.ParseDuration(v)).(time.Duration)
			case "pattern":
				cfg.Pattern = v
			case "count":
				cfg.Count = uint(xerror.PanicErr(strconv.Atoi(v)).(int))
			}
		}

		w, err := rotate.NewRotateLogger(cfg)
		return &nopCloserSink{zapcore.AddSync(w)}, err
	}))
}

type nopCloserSink struct{ zapcore.WriteSyncer }

func (t nopCloserSink) Close() error { return nil }
