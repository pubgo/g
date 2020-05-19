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
	//xrr.PanicF(nil, "sssss %#v", a)
	xrr.PanicF(fmt.Errorf("ss"), "sssss %#v", a)

	return
}

func init21(a ...interface{}) (err error) {
	//xrr := WithErr(&err)
	//defer xrr.Recover()

	//fmt.Println(a...)
	//xrr.Panic(fmt.Errorf("ss"))
	//xrr.PanicF(init22(a...), "sssss %#v", a)

	return init22(a...)
}

func TestName(t *testing.T) {
	UnWrap(init21(1, 2, 3)).P()
	//UnWrap(nil).P()
	Exit(init21(1, 2, 3))
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		init21(1, 2, 3)
	}
}
