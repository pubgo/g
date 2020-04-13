package redirect

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RedirectConfig defines the config for Redirect middleware.
type RedirectConfig struct {
	Code int
}

// redirectLogic represents a function that given a scheme, host and uri
type redirectLogic func(scheme, host, uri string) (ok bool, url string)

const www = "www."

// DefaultRedirectConfig is the default Redirect middleware config.
var DefaultRedirectConfig = RedirectConfig{
	Code: http.StatusMovedPermanently,
}

// HTTPSRedirect redirects http requests to https.
func HTTPSRedirect() gin.HandlerFunc {
	return HTTPSRedirectWithConfig(DefaultRedirectConfig)
}

// HTTPSRedirectWithConfig returns an HTTPSRedirect middleware with config.
func HTTPSRedirectWithConfig(config RedirectConfig) gin.HandlerFunc {
	return redirect(config, func(scheme, host, uri string) (ok bool, url string) {
		if ok = scheme != "https"; ok {
			url = "https://" + host + uri
		}
		return
	})
}

// HTTPSWWWRedirect redirects http requests to https www.
func HTTPSWWWRedirect() gin.HandlerFunc {
	return HTTPSWWWRedirectWithConfig(DefaultRedirectConfig)
}

// HTTPSWWWRedirectWithConfig returns an HTTPSRedirect middleware with config.
func HTTPSWWWRedirectWithConfig(config RedirectConfig) gin.HandlerFunc {
	return redirect(config, func(scheme, host, uri string) (ok bool, url string) {
		if ok = scheme != "https" && host[:4] != www; ok {
			url = "https://www." + host + uri
		}
		return
	})
}

// HTTPSNonWWWRedirect redirects http requests to https non www.
func HTTPSNonWWWRedirect() gin.HandlerFunc {
	return HTTPSNonWWWRedirectWithConfig(DefaultRedirectConfig)
}

// HTTPSNonWWWRedirectWithConfig returns an HTTPSRedirect middleware with config.
func HTTPSNonWWWRedirectWithConfig(config RedirectConfig) gin.HandlerFunc {
	return redirect(config, func(scheme, host, uri string) (ok bool, url string) {
		if ok = scheme != "https"; ok {
			if host[:4] == www {
				host = host[4:]
			}
			url = "https://" + host + uri
		}
		return
	})
}

// WWWRedirect redirects non www requests to www.
func WWWRedirect() gin.HandlerFunc {
	return WWWRedirectWithConfig(DefaultRedirectConfig)
}

// WWWRedirectWithConfig returns an HTTPSRedirect middleware with config.
func WWWRedirectWithConfig(config RedirectConfig) gin.HandlerFunc {
	return redirect(config, func(scheme, host, uri string) (ok bool, url string) {
		if ok = host[:4] != www; ok {
			url = scheme + "://www." + host + uri
		}
		return
	})
}

// NonWWWRedirect redirects www requests to non www.
func NonWWWRedirect() gin.HandlerFunc {
	return NonWWWRedirectWithConfig(DefaultRedirectConfig)
}

// NonWWWRedirectWithConfig returns an HTTPSRedirect middleware with config.
func NonWWWRedirectWithConfig(config RedirectConfig) gin.HandlerFunc {
	return redirect(config, func(scheme, host, uri string) (ok bool, url string) {
		if ok = host[:4] == www; ok {
			url = scheme + "://" + host[4:] + uri
		}
		return
	})
}

func redirect(config RedirectConfig, cb redirectLogic) gin.HandlerFunc {
	if config.Code == 0 {
		config.Code = DefaultRedirectConfig.Code
	}

	return func(c *gin.Context) {
		req := c.Request
		host := req.Host
		if ok, url := cb("http", host, req.RequestURI); ok {
			c.Redirect(config.Code, url)
		}
		c.Next()
	}
}
