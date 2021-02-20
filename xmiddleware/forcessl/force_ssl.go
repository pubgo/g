package forcessl

import (
	"github.com/gin-gonic/gin"
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xdi"
	"github.com/pubgo/xerror"
	"net/http"
	"strings"
	"testing"
)

func isNotSecure(c *gin.Context, xfpHeader string, trustXfpHeader bool) bool {
	if trustXfpHeader {
		return xfpHeader != "https"
	}
	return c.Request.TLS == nil && c.Request.URL.Scheme != "https"
}

func SetUp() gin.HandlerFunc {
	_message := "SSL Required."
	_trustXFPHeader := false
	_enable301Redirects := false
	xerror.Panic(xdi.Invoke(func(config *xconfig.Config) {
		_ssl := config.Web.Forcessl
		_message = _ssl.Message
		_trustXFPHeader = _ssl.TrustXFPHeader
		_enable301Redirects = _ssl.Enable301Redirects
	}))

	return func(c *gin.Context) {
		xfpHeader := strings.ToLower(c.GetHeader("X-Forwarded-Proto"))

		if isNotSecure(c, xfpHeader, _trustXFPHeader) {
			if _enable301Redirects {
				redirectURL := c.Request.URL
				redirectURL.Scheme = "https"
				c.Redirect(http.StatusMovedPermanently, redirectURL.String())
			} else {
				c.AbortWithStatus(http.StatusForbidden)
				c.String(http.StatusForbidden, _message)
			}
			return
		}
		c.Next()
	}
}

func TestName(t *testing.T) {
	//api.Use(forceSSL.Middleware{
	//	TrustXFPHeader: true,
	//	Enable301Redirects: true,
	//	Message: "We are unable to process your request over HTTP."
	//})

	//forceSSLMiddleware := &forceSSL.Middleware{
	//	TrustXFPHeader:     true,
	//	Enable301Redirects: false,
	//	Message:            "Login required for Admin portal.",
	//}

	//api.Use(&rest.IfMiddleware{
	//	Condition: func(request *rest.Request) bool {
	//		return request.URL.Path == "/admin"
	//	},
	//	IfTrue: forceSSLMiddleware,
	//})
}
