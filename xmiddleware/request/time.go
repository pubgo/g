package request

import (
	"github.com/gin-gonic/gin"
	"time"
)

// Time
// Time middleware
func Time() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Header("START_TIME", start.String())
		c.Next()
		c.Header("ELAPSED_TIME", time.Now().Sub(start).String())
	}
}
