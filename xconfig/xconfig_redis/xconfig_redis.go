package xconfig_redis

import (
	"github.com/pubgo/g/xconfig"
	"github.com/pubgo/g/xdi"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/pubgo/g/pkg"
	"github.com/pubgo/g/xerror"
)

type Redis struct {
	ss map[string]*redis.Client
}

// CloseRedis
// close redis connection
func (t *Redis) CloseRedis(name ...string) error {
	return t.GetRedis(name...).Close()
}

// GetRedis get redis instance with name
func (t *Redis) GetRedis(name ...string) (c *redis.Client) {
	_name := xconfig.DefaultName
	if len(name) > 0 {
		_name = name[0]
	}

	xerror.PanicT(_name == "", "name is empty")

	c = t.ss[_name]
	xerror.PanicT(pkg.IsNone(c), "redis instance %s is nil", _name)
	return
}

func init() {
	xdi.InitProvide(func(cfg *xconfig.Config) *Redis {
		defer xerror.Assert()

		// 加载配置
		_cfg := cfg.Redis
		xerror.PanicT(_cfg.Default == "", "default name is empty")
		xerror.PanicT(len(_cfg.Cfg) == 0, "redis config count is 0")

		ss := make(map[string]*redis.Client, len(_cfg.Cfg))

		for _, cfg := range _cfg.Cfg {
			var opt = new(redis.Options)
			if cfg.URL != "" {
				opt = xerror.PanicErr(_ParseRedisUrl(cfg.URL)).(*redis.Options)
			} else {
				opt.Addr = cfg.Addr
				opt.Network = cfg.Network
				opt.Password = cfg.Password
			}

			if opt.DB == 0 {
				opt.DB = cfg.Db
			}

			opt.MaxRetries = cfg.MaxRetries
			if cfg.MaxConnAge > 0 {
				opt.MaxConnAge = time.Second * time.Duration(cfg.MaxConnAge)
			}

			if cfg.DialTimeout > 0 {
				opt.DialTimeout = time.Second * time.Duration(cfg.DialTimeout)
			}

			if cfg.ReadTimeout > 0 {
				opt.ReadTimeout = time.Second * time.Duration(cfg.ReadTimeout)
			}

			if cfg.WriteTimeout > 0 {
				opt.WriteTimeout = time.Second * time.Duration(cfg.WriteTimeout)
			}

			if cfg.PoolSize > 0 {
				opt.PoolSize = cfg.PoolSize
			}

			if cfg.PoolTimeout > 0 {
				opt.PoolTimeout = time.Second * time.Duration(cfg.PoolTimeout)
			}

			if cfg.IdleTimeout > 0 {
				opt.IdleTimeout = time.Second * time.Duration(cfg.IdleTimeout)
			}

			if cfg.IdleCheckFrequency > 0 {
				opt.IdleCheckFrequency = time.Second * time.Duration(cfg.IdleCheckFrequency)
			}

			_rdb := redis.NewClient(opt)
			xerror.PanicM(_rdb.Ping().Err(), "redis %s ping failed", cfg.Name)
			ss[cfg.Name] = _rdb
		}

		ss[xconfig.DefaultName] = ss[_cfg.Default]
		return &Redis{ss: ss}
	})
}
