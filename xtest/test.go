package xtest

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pubgo/x/pkg/colorutil"
	"github.com/pubgo/x/xerror"
)

// Test Test
type Test struct {
	name string
	desc string
	fn   interface{}
	args []interface{}
}

// In input
func (t *Test) In(args ...interface{}) *Test {
	return &Test{fn: t.fn, args: args, desc: t.desc, name: t.name}
}

func (t *Test) _Err(b bool, fn ...interface{}) {
	fmt.Printf("  [Desc func %s] [%s]\n", colorutil.GreenBg(t.name+" start"), t.desc)
	_err := _try(t.fn)(t.args...)(fn...)
	_isErr := _isNone(_err)

	if (b && !_isErr) || (!b && _isErr) {
		fmt.Printf("  [Desc func %s] --> %s\n", colorutil.GreenBg(t.name+" ok"), _caller.FromDepth(3))
	} else {
		fmt.Printf("  [Desc func %s] --> %s\n", colorutil.RedBg(t.name+" fail"), _caller.FromDepth(3))
	}

	xerror.PanicTT((b && _isErr) || (!b && !_isErr), func(err xerror.IErr) {
		err.SetErr(fmt.Errorf("%s test error", t.name))
		err.M("input", t.args)
		err.M("desc", t.desc)
		err.M("func_name", t.name)
	})
}

// IsErr
// check if result is error
// fn is result callback
func (t *Test) IsErr(fn ...interface{}) {
	t._Err(true, fn...)
}

// IsNil
// check if result is nil
// fn is result callback
func (t *Test) IsNil(fn ...interface{}) {
	t._Err(false, fn...)
}

// Run
// run test
func Run(fn interface{}, desc func(desc func(string) *Test)) {
	_fn := reflect.ValueOf(fn)
	xerror.PanicM(xerror.AssertFn(_fn), "input is not a function")

	_name := strings.Split(_caller.FromFunc(_fn), " ")[1]
	_funcName := strings.Split(_caller.FromFunc(_fn), " ")[1] + strings.TrimLeft(_fn.Type().String(), "func")
	_path := strings.Split(_caller.FromFunc(reflect.ValueOf(desc)), " ")[0]

	fmt.Printf("[Test func %s] [%s] --> %s\n", colorutil.GreenBg(_name+" start"), _funcName, _path)
	_err := _try(desc)(func(s string) *Test {
		return &Test{desc: s, fn: fn, name: _name}
	})()
	if _err != nil {
		fmt.Printf("[Test func %s]\n", colorutil.RedBg(_name+" fail"))
	} else {
		fmt.Printf("[Test func %s]\n", colorutil.GreenBg(_name+" success"))
	}
	xerror.Panic(_err)
}

func init() {
	Run(nil, func(desc func(string) *Test) {
		desc("").In()
	})
}
