package xerror_test

import (
	"errors"
	"fmt"
	"github.com/pubgo/g/retry"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/pubgo/g/pkg"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/g/xinit"
	"github.com/pubgo/g/xtest"
	"github.com/pubgo/g/xtry"
)

func init() {
	xerror.Exit(xinit.Start(), "init error")
}

var _try = xtry.Try
var _retry = retry.Retry

func init21(a ...interface{}) (b string, err error) {
	fmt.Println(a...)
	return "22", fmt.Errorf("ss")
}

func TestT1(t *testing.T) {
	defer xerror.Debug()

	fmt.Println(xerror.PanicErr(init21(1, 2)).(string))
}

func TestT(t *testing.T) {
	xtest.Run(xerror.PanicT, func(desc func(string) *xtest.Test) {
		desc("params is true").In(true, "test t").IsErr()
		desc("params is false").In(false, "test t").IsNil()
	})
}

func TestRetry(t *testing.T) {
	defer xerror.Assert()

	xtest.Run(_retry, func(desc func(string) *xtest.Test) {
		desc("retry(3)").In(3, func() {
			xerror.PanicT(true, "test t")
		}).IsErr(func(err error) {
			xerror.PanicM(err, "test Retry error")
		})
	})
}

func TestIf(t *testing.T) {
	defer xerror.Assert()

	xerror.PanicT(pkg.If(true, "test true", "test false").(string) != "test true", "")
}

func TestTT(t *testing.T) {
	defer xerror.Assert()

	_fn := func(b bool) {
		xerror.PanicTT(b, func(err xerror.IErr) {
			err.M("k", "v")
		})
	}

	xtest.Run(_fn, func(desc func(string) *xtest.Test) {
		desc("true params 1").In(true).IsErr()
		desc("true params 2").In(true).IsErr()
		desc("true params 3").In(true).IsErr()
		desc("false params").In(false).IsNil()
	})
}

func TestWrap(t *testing.T) {
	defer xerror.Assert()

	xerror.PanicM(errors.New("test"), "test")
}

func TestWrapM(t *testing.T) {
	defer xerror.Assert()

	xerror.PanicM(errors.New("dd"), "test")
}

func testFunc2() {
	xerror.PanicMM(errors.New("testFunc_1"), func(err xerror.IErr) {
		err.M("ss", 1)
		err.M("input", 2)
	})
}

func testFunc1() {
	testFunc2()
}

func testFunc() {
	xerror.PanicM(_try(testFunc1), "xerror.Wrap")
}

func TestErrLog(t *testing.T) {
	defer xerror.Assert()

	xtest.Run(testFunc, func(desc func(string) *xtest.Test) {
		desc("test func").In().IsErr()
	})
}

func init11() {
	xerror.PanicT(true, "test tt")
}

func TestT2(t *testing.T) {
	defer xerror.Assert()

	xtest.Run(init11, func(desc func(string) *xtest.Test) {
		desc("simple test").In().IsErr()
	})
}

func TestTry(t *testing.T) {
	defer xerror.Assert()

	xerror.Panic(_try(xerror.PanicT)(true, "sss"))
}

func TestTask(t *testing.T) {
	defer xerror.Assert()

	xerror.PanicM(_try(func() {
		xerror.PanicM(errors.New("dd"), "err ")
	}), "test wrap")
}

func TestHandle(t *testing.T) {
	defer xerror.Assert()

	func() {
		xerror.PanicM(errors.New("hello error"), "sss")
	}()

}

func TestErrHandle(t *testing.T) {
	defer xerror.Assert()

	xerror.ErrHandle(_try(func() {
		xerror.PanicT(true, "test T")
	}), func(err xerror.IErr) {
		err.P()
	})

	xerror.ErrHandle("ttt", func(err xerror.IErr) {
		err.P()
	})
	xerror.ErrHandle(errors.New("eee"), func(err xerror.IErr) {
		err.P()
	})
	xerror.ErrHandle([]string{"dd"}, func(err xerror.IErr) {
		err.P()
	})
}

func TestIsZero(t *testing.T) {
	//defer xerror.Log()

	var ss = func() map[string]interface{} {
		return make(map[string]interface{})
	}

	var ss1 = func() map[string]interface{} {
		return nil
	}

	var s = 1
	var ss2 map[string]interface{}
	xerror.PanicT(pkg.IsZero(reflect.ValueOf(1)), "")
	xerror.PanicT(pkg.IsZero(reflect.ValueOf(1.2)), "")
	xerror.PanicT(!pkg.IsZero(reflect.ValueOf(nil)), "")
	xerror.PanicT(pkg.IsZero(reflect.ValueOf("ss")), "")
	xerror.PanicT(pkg.IsZero(reflect.ValueOf(map[string]interface{}{})), "")
	xerror.PanicT(pkg.IsZero(reflect.ValueOf(ss())), "")
	xerror.PanicT(!pkg.IsZero(reflect.ValueOf(ss1())), "")
	xerror.PanicT(pkg.IsZero(reflect.ValueOf(&s)), "")
	xerror.PanicT(!pkg.IsZero(reflect.ValueOf(ss2)), "")
}

func TestResp(t *testing.T) {
	defer xerror.Assert()

	xtest.Run(xerror.Resp, func(desc func(string) *xtest.Test) {
		desc("resp ok").In(func(err xerror.IErr) {
			err.Caller(pkg.Caller.FromDepth(2))
		}).IsNil()
	})

}

func TestTicker(t *testing.T) {
	defer xerror.Assert()

	retry.Ticker(func(i int) time.Duration {
		fmt.Println(i)
		return time.Second
	})
}

func TestRetryAt(t *testing.T) {
	retry.At(time.Second*2, func(i int) {
		fmt.Println(i)

		xerror.PanicT(true, "test RetryAt")
	})
}

func TestErr(t *testing.T) {
	xerror.ErrHandle(_try(func() {
		xerror.ErrHandle(_try(func() {
			xerror.PanicT(true, "90999 error")
		}), func(err xerror.IErr) {
			xerror.PanicM(err, "wrap")
		})
	}), func(err xerror.IErr) {
		err.P()
	})
}

func _GetCallerFromFn2() {
	xerror.PanicMM(errors.New("test 123"), func(err xerror.IErr) {
		err.M("ss", "dd")
	})
}

func _GetCallerFromFn1(fn func()) {
	xerror.Panic(xerror.AssertFn(reflect.ValueOf(fn)))
	fn()
}

func TestGetCallerFromFn(t *testing.T) {
	defer xerror.Assert()

	fmt.Println(pkg.Caller.FromFunc(reflect.ValueOf(_GetCallerFromFn1)))

	xtest.Run(_GetCallerFromFn1, func(desc func(string) *xtest.Test) {
		desc("GetCallerFromFn ok").In(_GetCallerFromFn2).IsErr()
		desc("GetCallerFromFn nil").In(nil).IsErr()
	})
}

func TestTest(t *testing.T) {
	defer xerror.Assert()

	xtest.Run(xerror.AssertFn, func(desc func(string) *xtest.Test) {
		desc("params is func 1").
			In(reflect.ValueOf(func() {})).
			IsNil(func(err error) {
				xerror.PanicM(err, "check error")
			})

		desc("params is func 2").
			In(reflect.ValueOf(func() {})).
			IsNil(func(err error) {
				xerror.PanicM(err, "check error")
			})

		desc("params is func 3").
			In(reflect.ValueOf(func() {})).
			IsNil(func(err error) {
				xerror.PanicM(err, "check error")
			})

		desc("params is nil").
			In(reflect.ValueOf(nil)).
			IsErr(func(err error) {
				xerror.PanicM(err, "check error ok")
			})
	})
}

func TestLoadEnv(t *testing.T) {
	xerror.Panic(xenv.LoadFile("../.env"))
	xerror.PanicT(os.Getenv("a") != "1", "env error")
}

func init2() (err error) {
	defer xerror.RespErr(&err)

	xerror.PanicTT(true, func(err xerror.IErr) {
		err.SetErr("ok sss %d", 23)
	})
	return
}

func TestSig(t *testing.T) {
	defer xerror.Assert()

	xerror.Panic(init2())
}

func TestIsNone(t *testing.T) {
	defer xerror.Debug()

	xtest.Run(pkg.IsNone, func(desc func(string) *xtest.Test) {
		desc("is null").In(nil).IsNil(func(b bool) {
			xerror.PanicT(b != true, "error")
		})
		desc("is ok").In("ok").IsNil(func(b bool) {
			xerror.PanicT(b == false, "error")
		})
	})
}

func OOnit() (string, error) {
	return "test", errors.New("ss")
	//return "test", nil
}

func TestResultHandle2(t *testing.T) {
	//xerror.ResHandle(xerror.ErrDone)
	//xerror.ResHandle("ss", xerror.ErrDone)
	//xerror.ResHandle("ss", xerror.ErrDone, func(err xerror.IErr) {})

	//xerror.ResHandle("ss", xerror.ErrDone, func(s string, err xerror.IErr) {})

}

var YErr = xerror.NewXErr("YErr")
var ErrDone = YErr.New("Done")
var ErrDone1 = YErr.New("Done1")

func TestYErr(t *testing.T) {
	var s = func(a interface{}) {
		_a := reflect.ValueOf(a)
		fmt.Println(_a.Kind(), _a.String(), _a.Type(), _a.Type().String(), _a.Type().Kind(), _a.Type().Name(), a)
		print("\n\n")
	}

	defer xerror.Resp(func(err xerror.IErr) {
		fmt.Println(err.Is(ErrDone1))
	})

	//ErrDone.Panic()

	s(ErrDone)
	s(ErrDone1)
	s(ErrDone1.Err("ss"))
	s(ErrDone1)
	xerror.Panic(ErrDone1)
}

func TestDebug(t *testing.T) {
	//xenv.SetDebug()
	//xinit.Start()
	xerror.Debug("hello", map[string]interface{}{})
}
