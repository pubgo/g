package xconfig_web

import (
	"github.com/Masterminds/sprig"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xdi"
	"github.com/pubgo/x/xerror"
	"github.com/pubgo/x/xmiddleware/admin"
	"github.com/pubgo/x/xmiddleware/cors"
	"github.com/pubgo/x/xmiddleware/csrf"
	"github.com/pubgo/x/xmiddleware/gzip"
	"github.com/pubgo/x/xmiddleware/logger"
	"github.com/pubgo/x/xmiddleware/recover"
	"github.com/pubgo/x/xmiddleware/slash"
	"github.com/pubgo/x/xmiddleware/static"
)

type XWeb struct {
	web *gin.Engine
}

func (t *XWeb) Run() error {
	return t.web.Run()
	//t.web.RunFd()
	//t.web.RunTLS()
	//t.web.RunUnix()
}

func (t *XWeb) Router(route func(web *gin.Engine)) {
	route(t.web)
}

func init() {
	xdi.InitProvide(func(config *xconfig.Config) (_ *XWeb, err error) {
		defer xerror.RespErr(&err)

		cfg := config.Web
		xerror.PanicT(!cfg.Enabled, "web service not enabled")

		_web := gin.New()

		if cfg.Logger.Enabled {
			_web.Use(logger.SetUp())
		}

		if cfg.Static.Enabled {
			_web.Use(static.SetUp())
		}

		if cfg.IsSlash {
			_web.Use(slash.SetUp())
		}

		if cfg.Admin.Enabled {
			_web.Use(admin.SetUp())
		}

		if cfg.Auth.Enabled {
			_web.Use(admin.SetUp())
		}

		if cfg.Cors.Enabled {
			_web.Use(cors.SetUp())
		}

		if cfg.XSRF.Enabled {
			_web.Use(csrf.SetUp())
		}

		if cfg.Jwt.Enabled {

		}

		if cfg.EnableDocs {

		}

		if cfg.View.Enabled {
			_web.HTMLRender = ginview.New(goview.Config{
				Root:         cfg.View.ViewPath,
				Extension:    ".html",
				Master:       "layouts/master",
				Partials:     []string{},
				Funcs:        sprig.FuncMap(),
				DisableCache: false,
				Delims:       goview.Delims{Left: "{{", Right: "}}"},
			})
		}

		if cfg.Static.Enabled {
			_web.Use(static.SetUp())
		}

		if cfg.Logger.Enabled {
			_web.Use(logger.SetUp())
		}

		if cfg.IsRecovery {
			_web.Use(recover.SetUp())
		}

		//if cfg.Session.Enabled {
		//	_web.Use(sessions.Sessions("", session_redis_store.New(cfg.Name)))
		//}

		if cfg.Gzip.Enable {
			_web.Use(gzip.SetUp())
		}

		if cfg.Cookie.Enabled {
		}

		if cfg.Upload.Enabled {

		}

		return &XWeb{web: _web}, nil
	})
}
