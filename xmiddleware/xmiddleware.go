package xmiddleware

import (
	"github.com/gin-gonic/gin"
)

func GetHeader(c *gin.Context, head ...string) string {
	for _, h := range head {
		if _h := c.GetHeader(h); _h != "" {
			return _h
		}
	}
	return ""
}

func Get(c *gin.Context, head ...string) interface{} {
	for _, h := range head {
		if _h, e := c.Get(h); e {
			return _h
		}
	}
	return nil
}
