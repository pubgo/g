package pkg

import (
	"errors"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// ErrNoFuncCaller not found error
var ErrNoFuncCaller = errors.New("pkg: no func caller error")

type _Caller struct{}

// Caller init
var Caller = &_Caller{}

//var srcDir = filepath.Join(build.Default.GOPATH, "src") + string(os.PathSeparator)
//var modDir = filepath.Join(build.Default.GOPATH, "pkg", "mod") + string(os.PathSeparator)

func (t *_Caller) FromDepth(callDepth int) string {
	fn, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return ErrNoFuncCaller.Error()
	}

	var buf = _bytesPool.Get().(*strings.Builder)
	defer _bytesPool.Put(buf)
	defer buf.Reset()

	buf.WriteString(file)
	buf.WriteString(":")
	buf.WriteString(strconv.Itoa(line))
	buf.WriteString(" ")

	ma := strings.Split(runtime.FuncForPC(fn).Name(), ".")
	buf.WriteString(ma[len(ma)-1])
	return buf.String()
}

func (t *_Caller) FromFunc(fn reflect.Value) string {
	var _fn = fn.Pointer()
	var _e = runtime.FuncForPC(_fn)
	var file, line = _e.FileLine(_fn)

	var buf = _bytesPool.Get().(*strings.Builder)
	defer _bytesPool.Put(buf)
	defer buf.Reset()

	buf.WriteString(file)
	buf.WriteString(":")
	buf.WriteString(strconv.Itoa(line))
	buf.WriteString(" ")

	ma := strings.Split(_e.Name(), ".")
	buf.WriteString(ma[len(ma)-1])
	return buf.String()
}

var _bytesPool = &sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}
