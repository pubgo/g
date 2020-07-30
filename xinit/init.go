package xinit

import (
	slog "log"
	"os"
)

var log = slog.New(os.Stderr, "", slog.LstdFlags|slog.Lshortfile|slog.Ltime)

var inits []func() error

// Go init
func Go(fn func() error) {
	if fn == nil {
		log.Fatalln("fn is nil")
	}
	inits = append(inits, fn)
}

// Start init start
func Start() {
	defer func() {
		if err1 := recover(); err1 != nil {
			log.Fatalln(err1)
		}
	}()

	for _, fn := range inits {
		fn()
	}
	return
}
