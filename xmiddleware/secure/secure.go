package secure

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pubgo/g/cnst"
)

type (
	SecureConfig struct {
		XSSProtection         string
		ContentTypeNosniff    string
		XFrameOptions         string
		HSTSMaxAge            int
		HSTSExcludeSubdomains bool
		ContentSecurityPolicy string
		CSPReportOnly         bool
		HSTSPreloadEnabled    bool
		ReferrerPolicy        string
	}
)

var (
	// DefaultSecureConfig is the default Secure middleware config.
	DefaultSecureConfig = SecureConfig{
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "SAMEORIGIN",
		HSTSPreloadEnabled: false,
	}
)

// Secure returns a Secure middleware.
// Secure middleware provides protection against cross-site scripting (XSS) attack,
// content type sniffing, clickjacking, insecure connection and other code injection
// attacks.
func Secure() gin.HandlerFunc {
	return SecureWithConfig(DefaultSecureConfig)
}

// SecureWithConfig returns a Secure middleware with config.
// See: `Secure()`.
func SecureWithConfig(config SecureConfig) gin.HandlerFunc {
	// Defaults

	return func(c *gin.Context) {
		req := c.Request
		res := c.Writer

		if config.XSSProtection != "" {
			res.Header().Set(cnst.HeaderXXSSProtection, config.XSSProtection)
		}
		if config.ContentTypeNosniff != "" {
			res.Header().Set(cnst.HeaderXContentTypeOptions, config.ContentTypeNosniff)
		}
		if config.XFrameOptions != "" {
			res.Header().Set(cnst.HeaderXFrameOptions, config.XFrameOptions)
		}

		if req.Header.Get(cnst.HeaderXForwardedProto) == "https" && config.HSTSMaxAge != 0 {
			subdomains := ""
			if !config.HSTSExcludeSubdomains {
				subdomains = "; includeSubdomains"
			}
			if config.HSTSPreloadEnabled {
				subdomains = fmt.Sprintf("%s; preload", subdomains)
			}
			res.Header().Set(cnst.HeaderStrictTransportSecurity, fmt.Sprintf("max-age=%d%s", config.HSTSMaxAge, subdomains))
		}
		if config.ContentSecurityPolicy != "" {
			if config.CSPReportOnly {
				res.Header().Set(cnst.HeaderContentSecurityPolicyReportOnly, config.ContentSecurityPolicy)
			} else {
				res.Header().Set(cnst.HeaderContentSecurityPolicy, config.ContentSecurityPolicy)
			}
		}
		if config.ReferrerPolicy != "" {
			res.Header().Set(cnst.HeaderReferrerPolicy, config.ReferrerPolicy)
		}

		c.Next()
	}
}


// SecureHeaders adds general security headers for basic security measures
func SecureHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Protects from MimeType Sniffing
		c.Header("X-Content-Type-Options", "nosniff")
		// Prevents browser from prefetching DNS
		c.Header("X-DNS-Prefetch-Control", "off")
		// Denies website content to be served in an iframe
		c.Header("X-Frame-Options", "DENY")
		c.Header("Strict-Transport-Security", "max-age=5184000; includeSubDomains")
		// Prevents Internet Explorer from executing downloads in site's context
		c.Header("X-Download-Options", "noopen")
		// Minimal XSS protection
		c.Header("X-XSS-Protection", "1; mode=block")
	}
}
