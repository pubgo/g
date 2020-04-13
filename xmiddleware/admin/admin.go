package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/pubgo/g/xmiddleware"
	"strings"
)

func SetUp() gin.HandlerFunc {
	_adminPrefix := "/admin"
	_adminToken := ""

	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, _adminPrefix) {
			c.Next()
			return
		}

		reqToken := xmiddleware.GetHeader(c, "Authorization")
		if reqToken == "" {
			c.AbortWithStatus(401)
			return
		}

		if _adminToken == "" || reqToken != _adminToken {
			c.AbortWithStatus(401)
		} else {
			c.Next()
		}
	}
}
