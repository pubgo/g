package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xconfig/xconfig_log"
	"github.com/pubgo/x/xdi"
	"github.com/pubgo/xerror"
	"net/http"
	"regexp"
	"time"
)

// SetUp
// initializes the logging middleware.
func SetUp() gin.HandlerFunc {
	var (
		skip             map[string]struct{}
		_isUTC           = false
		_skipPathRegex   *regexp.Regexp
		_isSkipPathRegex func(path string) bool
		_logger          xconfig_log.Log
	)
	xerror.Panic(xdi.Invoke(func(logger xconfig_log.Log, config *xconfig.Config) {
		_logger = logger
		_logCfg := config.Web.Logger
		_isUTC = _logCfg.UTC
		_skipPathRegex = xerror.ExitErr(regexp.Compile(_logCfg.SkipPathRegex)).(*regexp.Regexp)
		_isSkipPathRegex = func(path string) bool {
			return _logCfg.SkipPathRegex != "" && _skipPathRegex.MatchString(path)
		}

		if length := len(_logCfg.SkipPath); length > 0 {
			skip = make(map[string]struct{}, length)
			for _, path := range _logCfg.SkipPath {
				skip[path] = struct{}{}
			}
		}
	}))

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		c.Next()

		if _, ok := skip[path]; ok {
			return
		}

		if _isSkipPathRegex(path) {
			return
		}

		end := time.Now()
		latency := end.Sub(start)
		if _isUTC {
			end = end.UTC()
		}

		msg := "Request"
		if len(c.Errors) > 0 {
			msg = c.Errors.String()
		}

		dumpLogger := _logger.With().
			Int("status", c.Writer.Status()).
			Str("remote_addr", c.Request.RemoteAddr).
			Str("uri", c.Request.RequestURI).
			Str("referer", c.Request.Referer()).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Str("user-agent", c.Request.UserAgent()).
			Logger()

		switch {
		case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
			dumpLogger.Warn().Msg(msg)
		case c.Writer.Status() >= http.StatusInternalServerError:
			dumpLogger.Error().Msg(msg)
		default:
			dumpLogger.Info().Msg(msg)
		}
	}
}
