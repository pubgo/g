package ip

import (
	"github.com/gin-gonic/gin"
	"net"
)

const RemoteAddr = "app_remote_addr"

func RemoteIp() gin.HandlerFunc {
	const (
		XForwardedFor = "X-Forwarded-For"
		XRealIP       = "X-Real-IP"
		XClientIP     = "x-client-ip"
	)

	return func(c *gin.Context) {
		remoteAddr := c.Request.RemoteAddr
		if ip := c.GetHeader(XClientIP); ip != "" {
			remoteAddr = ip
		} else if ip := c.GetHeader(XRealIP); ip != "" {
			remoteAddr = ip
		} else if ip = c.GetHeader(XForwardedFor); ip != "" {
			remoteAddr = ip
		} else {
			remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
		}

		if remoteAddr == "::1" {
			remoteAddr = "127.0.0.1"
		}

		c.Set(RemoteAddr, remoteAddr)
		c.Next()
	}
}
