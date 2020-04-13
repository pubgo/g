package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/pubgo/g/xconfig"
	"github.com/pubgo/g/xdi"
	"time"

	"github.com/gin-gonic/gin"
)

func SetUp() gin.HandlerFunc {

	var cfg *xconfig.Config
	xdi.InitInvoke(func(config *xconfig.Config) {
		cfg = config
	})
	_cors := cfg.Web.Cors
	return cors.New(cors.Config{
		AllowAllOrigins:  _cors.AllowAllOrigins,
		AllowOrigins:     _cors.AllowOrigins,
		AllowMethods:     _cors.AllowMethods,
		AllowHeaders:     _cors.AllowHeaders,
		AllowCredentials: _cors.AllowCredentials,
		MaxAge:           time.Second * time.Duration(_cors.MaxAge),
	})
}
