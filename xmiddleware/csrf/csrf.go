package csrf

import (
	"crypto/subtle"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pubgo/x/cnst"
	"github.com/pubgo/x/pkg/randutil"
	"net/http"
	"strings"
	"time"
)

type (
	// CSRFConfig defines the config for CSRF middleware.
	CSRFConfig struct {
		// Skipper defines a function to skip middleware.

		// TokenLength is the length of the generated token.
		TokenLength uint8 `yaml:"token_length"`
		// Optional. Default value 32.

		// TokenLookup is a string in the form of "<source>:<key>" that is used
		// to extract token from the request.
		// Optional. Default value "header:X-CSRF-Token".
		// Possible values:
		// - "header:<name>"
		// - "form:<name>"
		// - "query:<name>"
		TokenLookup string `yaml:"token_lookup"`

		// Context key to store generated CSRF token into context.
		// Optional. Default value "csrf".
		ContextKey string `yaml:"context_key"`

		// Name of the CSRF cookie. This cookie will store CSRF token.
		// Optional. Default value "csrf".
		CookieName string `yaml:"cookie_name"`

		// Domain of the CSRF cookie.
		// Optional. Default value none.
		CookieDomain string `yaml:"cookie_domain"`

		// Path of the CSRF cookie.
		// Optional. Default value none.
		CookiePath string `yaml:"cookie_path"`

		// Max age (in seconds) of the CSRF cookie.
		// Optional. Default value 86400 (24hr).
		CookieMaxAge int `yaml:"cookie_max_age"`

		// Indicates if CSRF cookie is secure.
		// Optional. Default value false.
		CookieSecure bool `yaml:"cookie_secure"`

		// Indicates if CSRF cookie is HTTP only.
		// Optional. Default value false.
		CookieHTTPOnly bool `yaml:"cookie_http_only"`
	}

	// csrfTokenExtractor defines a function that takes `echo.Context` and returns
	// either a token or an error.
	csrfTokenExtractor func(*gin.Context) (string, error)
)

var (
	// DefaultCSRFConfig is the default CSRF middleware config.
	DefaultCSRFConfig = CSRFConfig{
		TokenLength:  32,
		TokenLookup:  "header:" + cnst.HeaderXCSRFToken,
		ContextKey:   "csrf",
		CookieName:   "_csrf",
		CookieMaxAge: 86400,
	}
)

// CSRF returns a Cross-Site Request Forgery (CSRF) middleware.
func SetUp() gin.HandlerFunc {
	c := DefaultCSRFConfig
	return CSRFWithConfig(c)
}

// CSRFWithConfig returns a CSRF middleware with config.
func CSRFWithConfig(config CSRFConfig) gin.HandlerFunc {
	// Defaults
	if config.TokenLength == 0 {
		config.TokenLength = DefaultCSRFConfig.TokenLength
	}
	if config.TokenLookup == "" {
		config.TokenLookup = DefaultCSRFConfig.TokenLookup
	}
	if config.ContextKey == "" {
		config.ContextKey = DefaultCSRFConfig.ContextKey
	}
	if config.CookieName == "" {
		config.CookieName = DefaultCSRFConfig.CookieName
	}
	if config.CookieMaxAge == 0 {
		config.CookieMaxAge = DefaultCSRFConfig.CookieMaxAge
	}

	// Initialize
	parts := strings.Split(config.TokenLookup, ":")
	extractor := csrfTokenFromHeader(parts[1])
	switch parts[0] {
	case "form":
		extractor = csrfTokenFromForm(parts[1])
	case "query":
		extractor = csrfTokenFromQuery(parts[1])
	}

	return func(c *gin.Context) {
		req := c.Request
		k, err := c.Cookie(config.CookieName)
		token := ""

		// Generate token
		if err != nil {
			token = randutil.String(int(config.TokenLength))
		} else {
			// Reuse token
			token = k
		}

		switch req.Method {
		case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodTrace:
		default:
			// Validate token only for requests which are not defined as 'safe' by RFC7231
			clientToken, err := extractor(c)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
			}
			if !validateCSRFToken(token, clientToken) {
				c.String(http.StatusForbidden, "", "invalid csrf token")
			}
		}

		// Set CSRF cookie
		c.SetCookie(
			config.CookieName,
			token,
			time.Now().Add(time.Duration(config.CookieMaxAge) * time.Second).Second(),
			config.CookiePath,
			config.CookieDomain,
			config.CookieSecure,
			config.CookieHTTPOnly,
		)

		// Store token in the context
		c.Set(config.ContextKey, token)

		// Protect clients from caching the response
		c.Writer.Header().Add(cnst.HeaderVary, cnst.HeaderCookie)

		c.Next()
	}
}

// csrfTokenFromForm returns a `csrfTokenExtractor` that extracts token from the
// provided request header.
func csrfTokenFromHeader(header string) csrfTokenExtractor {
	return func(c *gin.Context) (string, error) {
		return c.Request.Header.Get(header), nil
	}
}

// csrfTokenFromForm returns a `csrfTokenExtractor` that extracts token from the
// provided form parameter.
func csrfTokenFromForm(param string) csrfTokenExtractor {
	return func(c *gin.Context) (string, error) {
		token := c.Param(param)
		if token == "" {
			return "", errors.New("missing csrf token in the form parameter")
		}
		return token, nil
	}
}

// csrfTokenFromQuery returns a `csrfTokenExtractor` that extracts token from the
// provided query parameter.
func csrfTokenFromQuery(param string) csrfTokenExtractor {
	return func(c *gin.Context) (string, error) {
		token := c.Param(param)
		if token == "" {
			return "", errors.New("missing csrf token in the query string")
		}
		return token, nil
	}
}

func validateCSRFToken(token, clientToken string) bool {
	return subtle.ConstantTimeCompare([]byte(token), []byte(clientToken)) == 1
}
