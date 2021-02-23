package envutil

import (
	"os"
)

// SetEnv set environment variable with prefix
func SetEnv(e, v string) error {
	return os.Setenv(_EnvKey(e), _EnvValue(v))
}

// GetEnv get environment variable with prefix
func GetEnv(e ...string) string {
	for _, _e := range e {
		if _v := os.Getenv(_EnvKey(_e)); _v != "" {
			return _v
		}
	}
	return ""
}

// Exist check if the environment variable is defined
func Exist(e ...string) bool {
	for _, _e := range e {
		if _, b := os.LookupEnv(_EnvKey(_e)); b {
			return true
		}
	}
	return false
}

// SetDebug 设置debug模式
func SetDebug() (err error) {
	return SetEnv("debug", "t")
}

// SetDebugFalse 设置debug=false模式
func SetDebugFalse() (err error) {
	return SetEnv("debug", "f")
}

// SetSkipXerror 设置跳过xerror错误打印
func SetSkipXerror() (err error) {
	return SetEnv("skip_xerror", "true")
}

// IsSkipXerror skip xerror file
func IsSkipXerror() bool {
	return IsTrue(GetEnv("skip_xerror"))
}

// IsDebug 是否是debug模式
func IsDebug() bool {
	return IsTrue(GetEnv("debug", "is_debug", "isDebug"))
}

// IsVersion 是否打印项目或者类库的版本信息
func IsVersion() bool {
	return IsTrue(GetEnv("version"))
}

// SetVersion 开启打印版本信息环境变量
func SetVersion() (err error) {
	return SetEnv("version", "true")
}

// IsDev 是否是dev模式
func IsDev() bool {
	return Cfg.Env == RunMode.Dev
}

// IsStag 是否是stag模式
func IsStag() bool {
	return Cfg.Env == RunMode.Stag
}

// IsProd 是否是prod模式
func IsProd() bool {
	return Cfg.Env == RunMode.Prod
}

// IsTest 是否是test模式
func IsTest() bool {
	return Cfg.Env == RunMode.Test
}

// IsRelease 是否是release模式
func IsRelease() bool {
	return Cfg.Env == RunMode.Release
}
