package main

import (
	"github.com/pubgo/x/xlog"
)

func main() {
	log := xlog.GetLog()
	ss := log.With()
	ss.With()
	log.Error("hello",
		xlog.Skip(),
		xlog.Any("hss", "ss"),
	)
}
