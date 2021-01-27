package request

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pubgo/x/cnst"
	"log"
)

// RequestID
// X-Request-ID middleware
func RequestID() gin.HandlerFunc {
	var reqId uuid.UUID

	return func(c *gin.Context) {
		_id := c.GetHeader(cnst.HeaderXRequestID)
		if _id != "" {
			var err error
			reqId, err = uuid.Parse(_id)
			if err != nil {
				log.Printf("[NOTICE] Parsing request ID from %s header: %v", cnst.HeaderXRequestID, err)
				reqId = uuid.New()
			}
		} else {
			reqId = uuid.New()
		}
		c.Writer.Header().Set(cnst.HeaderXRequestID, reqId.String())
		c.Next()
	}
}
