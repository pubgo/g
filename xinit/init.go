package xinit

import (
	slog "log"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

var log = slog.New(os.Stderr, "", slog.LstdFlags|slog.Lshortfile|slog.Ltime)

var inits []func() error

// Load init
func Load(fn func() error) {
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
		if err := fn(); err != nil {
			log.Println(funcForPC(fn))
			log.Fatalln(err)
		}
	}
	return
}

func funcForPC(fn func() error) string {
	var _fn = reflect.ValueOf(fn).Pointer()
	var _e = runtime.FuncForPC(_fn)
	var file, line = _e.FileLine(_fn)

	var buf = &strings.Builder{}
	defer buf.Reset()

	buf.WriteString(file)
	buf.WriteString(":")
	buf.WriteString(strconv.Itoa(line))
	buf.WriteString(" ")

	ma := strings.Split(_e.Name(), ".")
	buf.WriteString(ma[len(ma)-1])
	return buf.String()
}
