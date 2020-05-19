package xerror

import (
	"fmt"
	"testing"
)

func init22(a ...interface{}) (err error) {
	xrr := WithErr(&err)
	defer xrr.Recover()

	//fmt.Println(a...)
	//xrr.Panic(fmt.Errorf("ss"))
	//Exit(New(""))
	//_ = fmt.Sprintf("ss")
	//_ = fmt.Errorf("ss")
	//_ = "ss" + "sss"
	//xrr.Panic(nil)
	xrr.PanicF(nil, "sssss %#v", a)
	//xrr.PanicF(fmt.Errorf("ss"), "sssss %#v", a)
	return
}

func init21(a ...interface{}) (err error) {
	xrr := WithErr(&err)
	defer xrr.Recover()
	//
	//fmt.Println(a...)
	//xrr.Panic(fmt.Errorf("ss"))
	//xrr.PanicF(init22(a...), "sssss %#v", a)
	return init22(a...)
}

func TestName(t *testing.T) {
	Debug = false
	UnWrap(init21(1, 2, 3)).P()
	UnWrap(fmt.Errorf("")).P()
	//Exit(init21(1, 2, 3))
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		init22(1, 2, 3)
		//init21(1, 2, 3)
	}
}
