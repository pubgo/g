package xlog_config

// 配置的加载方式
// 配置文件
// flags

import (
	"encoding/json"
	"errors"
	"net/url"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/pubgo/x/xlog/internal/log"
	"github.com/pubgo/xerror"
)

type Config = zap.Config

type config struct {
	zap.Config
	EncodeLevel    string
	EncodeTime     string
	EncodeDuration string
	EncodeCaller   string
	EncodeName     string
}

func init() {
	xerror.Exit(zap.RegisterEncoder("", func(config zapcore.EncoderConfig) (zapcore.Encoder, error) {
		xerror.Exit(errors.New("not implemented"))
		return nil, nil
	}))

	xerror.Exit(zap.RegisterSink("", func(url *url.URL) (zap.Sink, error) {
		xerror.Exit(errors.New("not implemented"))
		return nil, nil
	}))
}

func InitDevLog(opts ...Option) (err error) {
	defer xerror.RespErr(&err)
	logger := xerror.PanicErr(zap.NewDevelopment(opts...)).(*zap.Logger)
	xerror.Panic(logger.Sync())
	log.SetLog(logger)
	return
}

func InitProdLog(opts ...Option) (err error) {
	defer xerror.RespErr(&err)
	logger := xerror.PanicErr(zap.NewProduction(opts...)).(*zap.Logger)
	xerror.Panic(logger.Sync())
	log.SetLog(logger)
	return
}

func InitFromConfig(config zap.Config) (err error) {
	defer xerror.RespErr(&err)
	logger := xerror.PanicErr(config.Build()).(*zap.Logger)
	xerror.Panic(logger.Sync())
	log.SetLog(logger)
	return
}

func InitFromJson(config []byte) (err error) {
	defer xerror.RespErr(&err)
	var cfg zap.Config
	xerror.Panic(json.Unmarshal(config, &cfg))
	// 替换zap的encoder
	encoderPatch(config, &cfg)
	logger := xerror.PanicErr(cfg.Build()).(*zap.Logger)
	xerror.Panic(logger.Sync())
	log.SetLog(logger)
	return
}

func InitFromOptions(opt ...Option) (err error) {
	defer xerror.RespErr(&err)
	logger := xerror.PanicErr(zap.NewProductionConfig().Build(opt...)).(*zap.Logger)
	xerror.Panic(logger.Sync())
	log.SetLog(logger)
	return
}

func NewDevConfig() zap.Config {
	return zap.NewDevelopmentConfig()
}

func NewProdConfig() zap.Config {
	return zap.NewProductionConfig()
}
