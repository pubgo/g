package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatus(401)
		c.Next()
	}
}
