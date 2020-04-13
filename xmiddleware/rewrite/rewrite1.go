package rewrite

import (
	"github.com/gin-gonic/gin"
	"github.com/pubgo/g/xmiddleware"
	"regexp"
	"strings"
)

// Rewrite returns a Rewrite middleware.
// Rewrite middleware rewrites the URL path based on the provided rules.
func Rewrite1() gin.HandlerFunc {
	var rules map[string]string
	var rulesRegex map[*regexp.Regexp]string

	// Initialize
	for k, v := range rules {
		k = strings.Replace(k, "*", "(.*)", -1)
		k = k + "$"
		rulesRegex[regexp.MustCompile(k)] = v
	}

	return func(c *gin.Context) {
		req := c.Request

		// Rewrite
		for k, v := range rulesRegex {
			replacer := xmiddleware.CaptureToken(k, req.URL.Path)
			if replacer != nil {
				req.URL.Path = replacer.Replace(v)
				break
			}
		}
		c.Next()
	}
}
