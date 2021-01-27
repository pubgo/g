package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pubgo/x/xmiddleware"
)

const ApiVersion = "API_VERSION"

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		_version := xmiddleware.GetHeader(c, "API-VERSION", "X-API-VERSION")
		if _version == "" {
			_version = "default"
		}

		c.Set(ApiVersion, _version)
		c.Next()
	}
}
