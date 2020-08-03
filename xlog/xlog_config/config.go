package xlog_config

// 配置的加载方式
// 配置文件
// flags

import (
	"encoding/json"
	"github.com/pubgo/x/xlog/internal/log"
	"github.com/pubgo/xerror"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// rolling
// https://github.com/natefinch/lumberjack

type Config config

type encoderConfig struct {
	MessageKey     string `json:"messageKey" yaml:"messageKey" toml:"messageKey"`
	LevelKey       string `json:"levelKey" yaml:"levelKey" toml:"levelKey"`
	TimeKey        string `json:"timeKey" yaml:"timeKey" toml:"timeKey"`
	NameKey        string `json:"nameKey" yaml:"nameKey" toml:"nameKey"`
	CallerKey      string `json:"callerKey" yaml:"callerKey" toml:"callerKey"`
	StacktraceKey  string `json:"stacktraceKey" yaml:"stacktraceKey" toml:"stacktraceKey"`
	LineEnding     string `json:"lineEnding" yaml:"lineEnding" toml:"lineEnding"`
	EncodeLevel    string `json:"levelEncoder" yaml:"levelEncoder" toml:"levelEncoder"`
	EncodeTime     string `json:"timeEncoder" yaml:"timeEncoder" toml:"timeEncoder"`
	EncodeDuration string `json:"durationEncoder" yaml:"durationEncoder" toml:"durationEncoder"`
	EncodeCaller   string `json:"callerEncoder" yaml:"callerEncoder" toml:"callerEncoder"`
	EncodeName     string `json:"nameEncoder" yaml:"nameEncoder" toml:"nameEncoder"`
}

type samplingConfig struct {
	Initial    int `json:"initial" yaml:"initial" toml:"initial"`
	Thereafter int `json:"thereafter" yaml:"thereafter" toml:"thereafter"`
}

type config struct {
	zapOpts           []zap.Option
	Level             zap.AtomicLevel        `json:"level" yaml:"level" toml:"level"`
	Development       bool                   `json:"development" yaml:"development" toml:"development"`
	DisableCaller     bool                   `json:"disableCaller" yaml:"disableCaller" toml:"disableCaller"`
	DisableStacktrace bool                   `json:"disableStacktrace" yaml:"disableStacktrace" toml:"disableStacktrace"`
	Sampling          *samplingConfig        `json:"sampling" yaml:"sampling" toml:"sampling"`
	Encoding          string                 `json:"encoding" yaml:"encoding" toml:"encoding"`
	EncoderConfig     encoderConfig          `json:"encoderConfig" yaml:"encoderConfig" toml:"encoderConfig"`
	OutputPaths       []string               `json:"outputPaths" yaml:"outputPaths" toml:"outputPaths"`
	ErrorOutputPaths  []string               `json:"errorOutputPaths" yaml:"errorOutputPaths" toml:"errorOutputPaths"`
	InitialFields     map[string]interface{} `json:"initialFields" yaml:"initialFields" toml:"initialFields"`
}

func (t config) toZapLogger() (_ *zap.Logger, err error) {
	defer xerror.RespErr(&err)
	zapCfg := zap.Config{}
	xerror.Panic(json.Unmarshal(xerror.PanicBytes(json.Marshal(&t)), &zapCfg))

	var ok bool

	if zapCfg.EncoderConfig.EncodeLevel, ok = levelEncoder[t.EncoderConfig.EncodeLevel]; !ok {
		zapCfg.EncoderConfig.EncodeLevel = levelEncoder[defaultKey]
	}

	if zapCfg.EncoderConfig.EncodeTime, ok = timeEncoder[t.EncoderConfig.EncodeTime]; !ok {
		zapCfg.EncoderConfig.EncodeTime = timeEncoder[defaultKey]
	}

	if zapCfg.EncoderConfig.EncodeDuration, ok = durationEncoder[t.EncoderConfig.EncodeDuration]; !ok {
		zapCfg.EncoderConfig.EncodeDuration = durationEncoder[defaultKey]
	}

	if zapCfg.EncoderConfig.EncodeCaller, ok = callerEncoder[t.EncoderConfig.EncodeCaller]; !ok {
		zapCfg.EncoderConfig.EncodeCaller = callerEncoder[defaultKey]
	}

	if zapCfg.EncoderConfig.EncodeName, ok = nameEncoder[t.EncoderConfig.EncodeName]; !ok {
		zapCfg.EncoderConfig.EncodeName = nameEncoder[defaultKey]
	}
	return xerror.PanicErr(zapCfg.Build(t.zapOpts...)).(*zap.Logger), nil
}

type Option func(opts *config)

func InitFromOption(opts ...Option) (err error) {
	defer xerror.RespErr(&err)

	cfg := config(NewProdConfig())
	for _, opt := range opts {
		opt(&cfg)
	}
	log.SetLog(xerror.PanicErr(cfg.toZapLogger()).(*zap.Logger))
	return
}

func InitFromConfig(conf Config, opts ...Option) (err error) {
	defer xerror.RespErr(&err)

	cfg := config(conf)
	for _, opt := range opts {
		opt(&cfg)
	}
	log.SetLog(xerror.PanicErr(cfg.toZapLogger()).(*zap.Logger))
	return
}

func InitFromJson(conf []byte, opts ...Option) (err error) {
	defer xerror.RespErr(&err)

	var cfg config
	xerror.Panic(json.Unmarshal(conf, &cfg))

	for _, opt := range opts {
		opt(&cfg)
	}

	log.SetLog(xerror.PanicErr(cfg.toZapLogger()).(*zap.Logger))
	return
}

func NewDevConfig() Config {
	return Config{
		Level:       zap.NewAtomicLevelAt(DebugLevel),
		Development: true,
		Encoding:    "console",
		EncoderConfig: encoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			EncodeLevel:    "capitalColor",
			EncodeTime:     "iso8601",
			EncodeDuration: "string",
			EncodeCaller:   "default",
			LineEnding:     zapcore.DefaultLineEnding,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func NewProdConfig() Config {
	return Config{
		Level:       zap.NewAtomicLevelAt(InfoLevel),
		Development: false,
		Sampling: &samplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: encoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    "default",
			EncodeTime:     "default",
			EncodeDuration: "default",
			EncodeCaller:   "default",
			LineEnding:     zapcore.DefaultLineEnding,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}
