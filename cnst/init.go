package cnst

import (
	"fmt"
	"net/http"
)

// ErrHTTPTag err tags
var ErrHTTPTag = struct {
	ErrUnsupportedMediaType        string
	ErrNotFound                    string
	ErrUnauthorized                string
	ErrForbidden                   string
	ErrMethodNotAllowed            string
	ErrStatusRequestEntityTooLarge string
	ErrTooManyRequests             string
	ErrBadRequest                  string
	ErrBadGateway                  string
	ErrInternalServerError         string
	ErrRequestTimeout              string
	ErrServiceUnavailable          string
	ErrValidatorNotRegistered      string
	ErrRendererNotRegistered       string
	ErrInvalidRedirectCode         string
	ErrCookieNotFound              string
	ErrInvalidCertOrKeyType        string
}{
	ErrUnsupportedMediaType:        fmt.Sprintf("http.code.%d", http.StatusUnsupportedMediaType),
	ErrNotFound:                    fmt.Sprintf("http.code.%d", http.StatusNotFound),
	ErrUnauthorized:                fmt.Sprintf("http.code.%d", http.StatusUnauthorized),
	ErrForbidden:                   fmt.Sprintf("http.code.%d", http.StatusForbidden),
	ErrMethodNotAllowed:            fmt.Sprintf("http.code.%d", http.StatusMethodNotAllowed),
	ErrStatusRequestEntityTooLarge: fmt.Sprintf("http.code.%d", http.StatusRequestEntityTooLarge),
	ErrTooManyRequests:             fmt.Sprintf("http.code.%d", http.StatusTooManyRequests),
	ErrBadRequest:                  fmt.Sprintf("http.code.%d", http.StatusBadRequest),
	ErrBadGateway:                  fmt.Sprintf("http.code.%d", http.StatusBadGateway),
	ErrInternalServerError:         fmt.Sprintf("http.code.%d", http.StatusInternalServerError),
	ErrRequestTimeout:              fmt.Sprintf("http.code.%d", http.StatusRequestTimeout),
	ErrServiceUnavailable:          fmt.Sprintf("http.code.%d", http.StatusServiceUnavailable),
	ErrValidatorNotRegistered:      "validator not registered",
	ErrRendererNotRegistered:       "renderer not registered",
	ErrInvalidRedirectCode:         "invalid redirect status code",
	ErrCookieNotFound:              "cookie not found",
	ErrInvalidCertOrKeyType:        "invalid cert or key type, must be string or []byte",
}
