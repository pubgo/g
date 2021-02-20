package xconfig_api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xdi"
	"github.com/pubgo/xerror"
	"github.com/pubgo/x/xmiddleware/admin"
	"github.com/pubgo/x/xmiddleware/cors"
	"github.com/pubgo/x/xmiddleware/csrf"
	"github.com/pubgo/x/xmiddleware/gzip"
	"github.com/pubgo/x/xmiddleware/recover"
	"github.com/pubgo/x/xservice/xredis/session_redis_store"
)

type Web struct {
	_web map[string]*gin.Engine
}

func init() {
	xdi.InitProvide(func(config *xconfig.Config) (_ *Web, err error) {
		defer xerror.RespErr(&err)

		_cfg := config.Web
		xerror.PanicT(_cfg.Default == "", "default name is empty")
		xerror.PanicT(len(_cfg.Cfg) == 0, "config count is 0")

		_web := make(map[string]*gin.Engine, len(_cfg.Cfg))

		for _, cfg := range _cfg.Cfg {

			engine := gin.New()

			if cfg.Admin.Enabled {
				engine.Use(admin.SetUp())
			}

			if cfg.Auth.Enabled {
				engine.Use(admin.SetUp())
			}

			if cfg.Cors.Enabled {
				engine.Use(cors.SetUp(cfg.Name))
			}

			if cfg.XSRF.Enabled {
				engine.Use(csrf.SetUp())
			}

			if cfg.Jwt.Enabled {

			}

			if cfg.EnableDocs {

			}

			if cfg.View.Enabled {

			}

			if cfg.Static.Enabled {

			}

			if cfg.Logger.Enabled {

			}

			if cfg.IsRecovery {
				engine.Use(recover.New())
			}

			if cfg.Session.Enabled {
				engine.Use(sessions.Sessions("", session_redis_store.New(cfg.Name)))
			}

			if cfg.EnableGzip {
				engine.Use(gzip.New(cfg.GzipLevel))
			}

			if cfg.Cookie.Enabled {
			}

			if cfg.Upload.Enabled {

			}

			_web[cfg.Name] = engine
		}

		return &Web{_web: _web}, nil
	})
}
