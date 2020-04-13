package main

import (
	"fmt"
	"github.com/gen2brain/dlgs"
	"github.com/go-vgo/robotgo"
	"github.com/pubgo/g/xerror"
	hook "github.com/robotn/gohook"
	"time"
)

func main() {

	//xerror.Panic(robotgo.ActiveName("WeChat"))

	s := robotgo.Start()
	defer robotgo.End()

	//go func() {
	//	for {
	//		robotgo.AddEvent("command")
	//		ok = true
	//
	//		if len(evs) > 0 {
	//			_st := evs[0]
	//			_end := evs[len(evs)-1]
	//			fmt.Println(_st, _end)
	//		}
	//	}
	//}()

	var ok = true
	x1, x2, x3, x4 := 0, 0, 0, 0
	for ev := range s {
		if !ok {
			continue
		}

		if ev.Kind == hook.KeyUp {
			ok = true
		}

		if ev.Kind == hook.MouseMove {
			continue
		}

		if ev.Kind == hook.MouseHold {
			x1, x2 = int(ev.X), int(ev.Y)
		}

		if ev.Kind == hook.MouseDown {
			x3, x4 = int(ev.X), int(ev.Y)
			ok = false

			_name := fmt.Sprintf("%d.png", time.Now().Unix())
			robotgo.SaveCapture(_name, x1, x2, x3-x1, x4-x2)
			robotgo.Sleep(2)
			if xerror.PanicErr(dlgs.Question("dd", "is ok", false)).(bool) {
				ok = true
			}
		}
	}
}
