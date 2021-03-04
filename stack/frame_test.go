package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	var fn = func(i int) error { return nil }
	fmt.Println(Func(fn))
	vfn:=reflect.TypeOf(fn)
	fmt.Printf("%s\n",vfn)
}
