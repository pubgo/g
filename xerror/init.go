package xerror

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
)

// func caller depth
const (
	callDepth = 3
	DebugKey  = "debug"
)

var (
	// ErrDone done
	ErrDone                = errors.New("DONE")
	ErrBadRequest          = New(400, http.StatusText(400))
	ErrUnauthorized        = New(401, http.StatusText(401))
	ErrForbidden           = New(403, http.StatusText(403))
	ErrNotFound            = New(404, http.StatusText(404))
	ErrMethodNotAllowed    = New(405, http.StatusText(405))
	ErrTimeout             = New(408, http.StatusText(408))
	ErrConflict            = New(409, http.StatusText(409))
	ErrInternalServerError = New(500, http.StatusText(500))
	isErrNil               = func(err error) bool {
		return err == nil || err == ErrDone
	}
	Debug  bool
	logger = log.New(os.Stdout, "[xerror] ", log.LstdFlags|log.Lshortfile)
)

func init() {
	Debug = true
	if b, _ := strconv.ParseBool(os.Getenv(DebugKey)); !b {
		Debug = false
	}
}
