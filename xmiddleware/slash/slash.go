package slash

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type (
	TrailingSlashConfig struct {
		RedirectCode int
	}
)

func SetUp() gin.HandlerFunc {
	return AddTrailingSlashWithConfig(TrailingSlashConfig{})
}

// AddTrailingSlashWithConfig returns a AddTrailingSlash middleware with config.
// See `AddTrailingSlash()`.
func AddTrailingSlashWithConfig(config TrailingSlashConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		url := req.URL
		path := url.Path
		qs := c.Request.URL.RawQuery
		if !strings.HasSuffix(path, "/") {
			path += "/"
			uri := path
			if qs != "" {
				uri += "?" + qs
			}

			// Redirect
			if config.RedirectCode != 0 {
				c.Redirect(config.RedirectCode, uri)
			}

			// Forward
			req.RequestURI = uri
			url.Path = path
		}
		c.Next()
	}
}

// RemoveTrailingSlash returns a root level (before router) middleware which removes a trailing slash from the request URI.
func RemoveTrailingSlash() gin.HandlerFunc {
	return RemoveTrailingSlashWithConfig(TrailingSlashConfig{})
}

// RemoveTrailingSlashWithConfig returns a RemoveTrailingSlash middleware with config.
func RemoveTrailingSlashWithConfig(config TrailingSlashConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		url := req.URL
		path := url.Path
		qs := c.Request.URL.RawQuery
		l := len(path) - 1
		if l > 0 && strings.HasSuffix(path, "/") {
			path = path[:l]
			uri := path
			if qs != "" {
				uri += "?" + qs
			}

			// Redirect
			if config.RedirectCode != 0 {
				c.Redirect(config.RedirectCode, uri)
			}

			// Forward
			req.RequestURI = uri
			url.Path = path
		}
		c.Next()
	}
}
