package envutil

import (
	"github.com/pubgo/x/typeutil"
	"os"
	"strings"
)

const (
	// DefaultPrefixSeparator 项目环境变量默认分隔符
	DefaultPrefixSeparator = '_'
	// DefaultPrefix 项目环境变量默认前缀
	DefaultPrefix = "ggg"
	// DefaultSecret 项目应用默认密码
	DefaultSecret = "zpCjWPsbqK@@^hR01qLDmZcXhKRIZgjHfxSG2KA%J#bFp!7YQVSmzXGc!sE!^qSM7@&d%oXHQtpR7K*8eRTdhRKjaxF#t@bd#A!"
)

var (
	trim  = strings.TrimSpace
	upper = strings.ToUpper
	// DefaultSecretKey 项目应用默认密码名字
	DefaultSecretKey = typeutil.StrOf(
		"secret", "token", "app_secret", "app.secret", "app_token", "app.token")
)

// RunMode 项目运行模式
var RunMode = struct {
	Dev     string
	Test    string
	Stag    string
	Prod    string
	Release string
}{
	Dev:     "dev",
	Test:    "test",
	Stag:    "stag",
	Prod:    "prod",
	Release: "release",
}

// Cfg 默认配置
var Cfg = struct {
	Prefix   string
	Env      string
	Debug    bool
	Home     string
	LogLevel string
}{
	Prefix: func() string {
		// DefaultPrefixKey 项目环境变量默认名字
		DefaultPrefixKey := typeutil.StrOf(
			"prefix",
			"app_prefix", "app.prefix", "app-prefix",
			"env_prefix", "env.prefix", "env-prefix",
			"prefix_name", "prefix.name", "prefix-name",
			"server_prefix", "server.prefix", "server-prefix",
		)

		for _, _e := range DefaultPrefixKey {
			if _v := os.Getenv(_e); _v != "" {
				return _v
			}
		}
		return ""
	}(),
	Env:      RunMode.Dev,
	Debug:    false,
	Home:     os.ExpandEnv("$PWD"),
	LogLevel: "debug",
}

func init() {

	// 环境变量前缀处理
	{
		var prefix = Cfg.Prefix
		if prefix == "" {
			prefix = DefaultPrefix
		}
		prefix = upper(trim(prefix))

		// 前缀分隔符处理
		if len(prefix) > 0 {
			_pl := len(prefix) - 1
			switch prefix[_pl] {
			case '_', '-', '.', '/', ' ':
				prefix = prefix[:_pl] + string(DefaultPrefixSeparator)
			default:
				prefix = prefix + string(DefaultPrefixSeparator)
			}
		}
		Cfg.Prefix = prefix
	}

	// 环境变量前缀处理 加载系统中带有本项目环境变量前缀的环境变量
	copyFromSystem(Cfg.Prefix)

	_envK := typeutil.StrOf(
		"env",
		"dev_env", "dev.env", "dev-env",
		"app_env", "app.env", "app-env",
		"run_mode", "run.mode", "run-mode",
		"env_mode", "env.mode", "env-mode",
	)
	if _env := GetEnv(_envK...); _env != "" {
		// 匹配开发或者线上环境
		switch _env {
		case RunMode.Dev, RunMode.Stag, RunMode.Prod, RunMode.Test, RunMode.Release:
			Cfg.Env = _env
		default:
			logger.Fatalf("run mode not match error, env %s", _env)
		}
	}

	// 是否是debug模式
	switch Cfg.Env {
	case RunMode.Dev, RunMode.Test, "":
		Cfg.Debug = true
		fatal(SetDebug())
	}

	// 是否跳过xerror里面的打印文件
	fatal(SetSkipXerror())

	return
}
