package xconfig

import (
	"github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/pubgo/x/pkg"
	"github.com/pubgo/x/xdi"
	"github.com/pubgo/x/xenv"
	"github.com/pubgo/xerror"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

const (
	DefaultName = "default"
	ExtName     = "ext"
	HomeFlag    = "home"
	ConfigType  = "toml"
	ConfigName  = "config"
)

// ExtDecode
// decode extended config data
func ExtDecode(data interface{}) (err error) {
	defer xerror.RespErr(&err)

	xerror.PanicT(pkg.IsNone(viper.Get(ExtName)), "%s value is nil", ExtName)
	xerror.PanicM(viper.UnmarshalKey(ExtName, data, func(cfg *mapstructure.DecoderConfig) {
		cfg.TagName = ConfigType
	}), "toml decode error")

	return
}

// HomeDir
// 获取配置文件所在目录
func HomeDir() string {
	return filepath.Dir(xerror.PanicErr(filepath.Abs(viper.ConfigFileUsed())).(string))
}

// initViperEnv
// sets to use Env variables if set.
func initViperEnv(prefix string) {
	viper.SetEnvPrefix(prefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "/"))
	viper.AutomaticEnv()
}

func init() {
	xdi.InitProvide(func() *Config {
		defer xerror.Assert()

		// 配置viper
		if xenv.Cfg.Prefix != "" {
			initViperEnv(xenv.Cfg.Prefix)
		}

		_cfg := &Config{}

		viper.SetConfigType(ConfigType)
		viper.SetConfigName(ConfigName)

		// 指定配置文件目录
		homeDir := viper.GetString(HomeFlag)
		viper.AddConfigPath(homeDir)
		viper.AddConfigPath(filepath.Join(homeDir, ConfigName))
		viper.AddConfigPath(filepath.Join("/app", ConfigName))
		_home := xerror.PanicErr(homedir.Dir()).(string)

		if xenv.Cfg.Service != "" {
			viper.AddConfigPath(filepath.Join(_home, ".config", xenv.Cfg.Service))
			viper.AddConfigPath(filepath.Join(_home, "."+xenv.Cfg.Service))
		}

		viper.AddConfigPath(ConfigName)
		viper.AddConfigPath(filepath.Join("pdd", ConfigName))

		if err := viper.ReadInConfig(); err != nil && !strings.Contains(strings.ToLower(err.Error()), "not found") {
			xerror.Exit(err, "read config failed")
		}

		if viper.ReadInConfig() == nil {
			xerror.PanicM(viper.Unmarshal(_cfg, func(cfg *mapstructure.DecoderConfig) {
				cfg.TagName = ConfigType
			}), "decode config failed")
		}

		return _cfg
	})
}
