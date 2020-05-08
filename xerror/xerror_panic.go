package xerror

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

var _funcCaller = func(callDepth int) []string {
	var buf = make([]string, 2)
	if debug {
		buf[0] = _caller.FromDepth(callDepth)
		buf[1] = _caller.FromDepth(callDepth + 1)
	}
	return buf
}

// Panic panic
func Panic(err error) {
	if isErrNil(err) {
		return
	}

	_e := _handle(err)
	_e.Caller(_funcCaller(callDepth + 1)...)
	panic(_e)
}

// PanicM error assert
func PanicM(err interface{}, msg interface{}, args ...interface{}) {
	if _isNone(err) {
		return
	}

	if _e := _handle(err); !_isNone(_e) {
		_e1 := &_Err{caller: _funcCaller(callDepth + 1)}
		_e1.SetErr(msg, args...)
		_e.node(_e1)
		panic(_e)
	}

}

// PanicMM error assert
func PanicMM(err interface{}, fn func(err IErr)) {
	if _isNone(err) {
		return
	}

	if _e := _handle(err); !_isNone(_e) {
		_e1 := &_Err{caller: _funcCaller(callDepth + 1)}
		fn(_e1)

		if _e1.Err() == nil {
			log.Fatalln("please set error")
		}
		_e.node(_e1)
		panic(_e)
	}
}

// Wrap error wrap
func Wrap(err interface{}, msg interface{}, args ...interface{}) error {
	if _isNone(err) {
		return nil
	}

	if _e := _handle(err); !_isNone(_e) {
		_e1 := &_Err{caller: _funcCaller(callDepth + 1)}
		_e1.SetErr(msg, args...)
		_e.node(_e1)
		return _e
	}
	return nil
}

// WrapM error wrap
func WrapM(err interface{}, fn func(err IErr)) error {
	if _isNone(err) {
		return nil
	}

	if _e := _handle(err); !_isNone(_e) {
		_e1 := &_Err{caller: _funcCaller(callDepth + 1)}
		fn(_e1)

		if _e1.Err() == nil {
			log.Fatalln("please set error")
		}
		_e.node(_e1)
		return _e
	}

	return nil
}

// PanicErr panic with func()(interface{},error)
func PanicErr(d1 interface{}, err error) interface{} {
	if _isNone(err) {
		return d1
	}

	if _e := _handle(err); !_isNone(_e) {
		_e.Caller(_funcCaller(callDepth + 1)...)
		panic(_e)
	}

	return d1
}

func PanicBytes(d1 interface{}, err error) []byte {
	if _isNone(err) {
		return d1.([]byte)
	}

	if _e := _handle(err); !_isNone(_e) {
		_e.Caller(_funcCaller(callDepth + 1)...)
		panic(_e)
	}

	return d1.([]byte)
}

func PanicStr(d1 interface{}, err error) string {
	if _isNone(err) {
		return d1.(string)
	}

	if _e := _handle(err); !_isNone(_e) {
		_e.Caller(_funcCaller(callDepth + 1)...)
		panic(_e)
	}

	return d1.(string)
}

// ExitErr exit when panic else return value
func ExitErr(d1 interface{}, err error) interface{} {
	if _isNone(err) {
		return d1
	}

	if _e := _handle(err); !_isNone(_e) {
		_e.Caller(_funcCaller(callDepth + 1)...)
		_e.P()
		os.Exit(-1)
	}
	return d1
}

// Exit os exit when panic
func ExitF(err interface{}, msg interface{}, args ...interface{}) {
	if _isNone(err) {
		return
	}

	if _e := _handle(err); !_isNone(_e) {
		_e1 := &_Err{caller: _funcCaller(callDepth + 1)}
		_e1.SetErr(msg, args...)
		_e.node(_e1)
		_e.P()
		os.Exit(-1)
	}
}

func Exit(err interface{}) {
	if _isNone(err) {
		return
	}

	if _e := _handle(err); !_isNone(_e) {
		_e1 := &_Err{caller: _funcCaller(callDepth + 1)}
		_e.node(_e1)
		_e.P()
		os.Exit(-1)
	}
}

// P err
func P(err error, msg string, args ...interface{}) {
	if !debug {
		return
	}

	if _e := _handle(err); !_isNone(_e) {
		_e1 := &_Err{caller: _funcCaller(callDepth + 1)}
		_e1.SetErr(msg, args...)
		_e.node(_e1)
		fmt.Println(_e.Error())
	}
}

// Debug err
func Debug(d ...interface{}) {
	if !debug {
		return
	}

	fmt.Print(colorize(fmt.Sprintf("%s, %s\n", time.Now().Format("2006/01/02 - 15:04:05"), _caller.FromDepth(callDepth)), colorRed))
	for _, i := range d {
		if i == nil || _isNone(i) {
			continue
		}

		switch _i := i.(type) {
		case string, []byte:
			fmt.Printf("%s\n", _i)
			continue
		}

		switch reflect.ValueOf(i).Kind() {
		case reflect.Struct, reflect.Map, reflect.Ptr:
			fmt.Printf("%s\n", PanicErr(json.MarshalIndent(i, "", "\t")))
		default:
			fmt.Printf("%#v\n", i)
		}
	}
}

// AssertFn errors
func AssertFn(fn reflect.Value) error {
	if _isZero(fn) || fn.Kind() != reflect.Func {
		return fmt.Errorf("the func is nil[%#v] or not func type[%s]", fn, fn.Kind())
	}
	return nil
}

// PanicT bool assert
func PanicT(b bool, msg interface{}, args ...interface{}) {
	if b {
		_e := &_Err{caller: _funcCaller(callDepth + 1)}
		_e.SetErr(msg, args...)
		panic(_e)
	}
}

// PanicTT bool assert
func PanicTT(b bool, fn func(err IErr)) {
	if b {
		_err := &_Err{caller: _funcCaller(callDepth + 1)}
		fn(_err)

		if _err.Err() == nil {
			log.Fatalln("please set error")
		}
		panic(_err)
	}
}
