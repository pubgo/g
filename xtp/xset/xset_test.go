package xset_test

import (
	"fmt"
	"github.com/pubgo/x/xtp/xset"
	"testing"
)

func TestA(t *testing.T) {
	_set := xset.NewSet()
	_set.Add(1)
	_set.Add(1)
	fmt.Println(_set.String())
}
