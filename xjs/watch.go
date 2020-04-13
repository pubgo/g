package xjs

import (
	"fmt"
	"github.com/dave/flux"
	. "github.com/siongui/godom"
)

func (t *Vapper) Dispatch(action flux.ActionInterface) chan struct{} {
	return t.dispatcher.Dispatch(action)
}

func (t *Vapper) Watch(key interface{}, f func(done chan struct{})) {
	t.watcher.Watch(key, f)
}

func (t *Vapper) Delete(key interface{}) {
	t.watcher.Delete(key)
}

func (t *Vapper) Fail(err error) {
	if err != nil {
		Window.Alert(err.Error())
	}
}

func (t *Vapper) Debug(message ...interface{}) {
	Window.Get("console").Call("log", message...)
}

func (t *Vapper) Log(message ...interface{}) {
	m := Document.GetElementById("message")
	if len(message) == 0 {
		m.SetInnerHTML("")
		return
	}

	s := fmt.Sprint(message[0])
	if m.InnerHTML() != s {
		requestAnimationFrame()
		m.SetInnerHTML(s)
		requestAnimationFrame()
	}
	t.Debug(message...)
}

func (t *Vapper) Logf(format string, args ...interface{}) {
	t.Log(fmt.Sprintf(format, args...))
}

func requestAnimationFrame() {
	c := make(chan struct{})
	Window.Call("requestAnimationFrame", func() { close(c) })
	<-c
}

func (t *Vapper) watch() {

	n := flux.NewNotifier()
	t.notifier = n
	t.watcher = n

	t.dispatcher = flux.NewDispatcher(
		// Notifier:
		t.notifier,
		// Stores:
		t.stores...
	)
}
