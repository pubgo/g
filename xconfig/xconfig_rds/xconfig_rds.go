package xconfig_rds

import (
	"encoding/json"
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xdi"
	"time"
	"xorm.io/core"
	"xorm.io/xorm"

	"github.com/pubgo/x/pkg"
	"github.com/pubgo/xerror"
)

type Rds struct {
	name string
	_rds map[string]*xorm.Engine
}

// GetRDS
// get rds instance with name
func (t *Rds) GetRDS(name ...string) (db *xorm.Engine) {
	_name := xconfig.DefaultName
	if len(name) > 0 {
		_name = name[0]
	}

	db = t._rds[_name]
	xerror.PanicT(pkg.IsNone(db), "rds instance %s is nil", _name)
	return t._rds[_name]
}

func init() {
	xdi.InitProvide(func(cfg *xconfig.Config) *Rds {
		defer xerror.Assert()

		// 加载配置
		_cfg := cfg.Rds
		xerror.PanicT(len(_cfg.Cfg) == 0, "rds config count is 0")
		_rds := make(map[string]*xorm.Engine, len(_cfg.Cfg))

		for _, cfg := range _cfg.Cfg {
			xerror.PanicT(cfg.Driver == "" || cfg.URL == "", "db driver or url is null")

			db := xerror.PanicErr(xorm.NewEngine(cfg.Driver, cfg.URL)).(*xorm.Engine)
			_log := newSQLLogger()
			_log.ShowSQL(true)
			db.SetLogger(_log)

			xerror.PanicT(_cfg.Prefix == "", "db prefix is null")
			db.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, _cfg.Prefix+"_"))

			if cfg.MaxIdle > 0 {
				db.SetMaxIdleConns(cfg.MaxIdle)
			}

			if cfg.MaxOpen > 0 {
				db.SetMaxOpenConns(cfg.MaxOpen)
			}

			if cfg.MaxLfetime > 0 {
				db.SetConnMaxLifetime(time.Second * time.Duration(cfg.MaxLfetime))
			}

			xerror.PanicM(db.Ping(), "rds %s ping failed", cfg.Name)
			if logger.Debug().Enabled() {
				logger.Debug().Msgf("DataSourceName %s", db.DataSourceName())
				logger.Debug().Msgf("DBMetas %s", xerror.PanicBytes(json.MarshalIndent(xerror.PanicErr(db.DBMetas()).([]*core.Table), "", "\t")))
			}
			_rds[cfg.Name] = db
		}
		_rds[xconfig.DefaultName] = _rds[_cfg.Default]
		return &Rds{_rds: _rds}
	})
}
