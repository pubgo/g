package xlog_config

import (
	"errors"
	"github.com/pubgo/xerror"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/url"
	"time"
)

var levelEncoder map[string]zapcore.LevelEncoder
var timeEncoder map[string]zapcore.TimeEncoder
var durationEncoder map[string]zapcore.DurationEncoder
var callerEncoder map[string]zapcore.CallerEncoder
var nameEncoder map[string]zapcore.NameEncoder

func RFC3339MilliTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
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
	levelEncoder = map[string]zapcore.LevelEncoder{

	}
	timeEncoder = map[string]zapcore.TimeEncoder{
		"rfc3339": RFC3339MilliTimeEncoder,
		"RFC3339": RFC3339MilliTimeEncoder,
	}
	durationEncoder = map[string]zapcore.DurationEncoder{

	}
	callerEncoder = map[string]zapcore.CallerEncoder{

	}
	nameEncoder = map[string]zapcore.NameEncoder{

	}

	xerror.Exit(zap.RegisterEncoder("test", func(config zapcore.EncoderConfig) (zapcore.Encoder, error) {
		xerror.Exit(errors.New("not implemented"))
		return nil, nil
	}))

	xerror.Exit(zap.RegisterSink("test", func(url *url.URL) (zap.Sink, error) {
		xerror.Exit(errors.New("not implemented"))
		return nil, nil
	}))
}

// encoderPatch 为zapcore的encoder做扩展
func encoderPatch(cfg []byte, config *zap.Config) {
	if fn, ok := levelEncoder[gjson.GetBytes(cfg, "encoderConfig.levelEncoder").String()]; ok {
		config.EncoderConfig.EncodeLevel = fn
	}

	if fn, ok := timeEncoder[gjson.GetBytes(cfg, "encoderConfig.timeEncoder").String()]; ok {
		config.EncoderConfig.EncodeTime = fn
	}

	if fn, ok := durationEncoder[gjson.GetBytes(cfg, "encoderConfig.durationEncoder").String()]; ok {
		config.EncoderConfig.EncodeDuration = fn
	}

	if fn, ok := callerEncoder[gjson.GetBytes(cfg, "encoderConfig.callerEncoder").String()]; ok {
		config.EncoderConfig.EncodeCaller = fn
	}

	if fn, ok := nameEncoder[gjson.GetBytes(cfg, "encoderConfig.nameEncoder").String()]; ok {
		config.EncoderConfig.EncodeName = fn
	}
}
