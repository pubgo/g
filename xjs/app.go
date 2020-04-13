// +build js

package xjs

import (
	"github.com/dave/flux"
	"github.com/pubgo/g/xerror"
	dom "github.com/siongui/godom"
	"reflect"
	"sync"
)

var _app sync.Once
var _vapper *Vapper

func Default() *Vapper {
	_app.Do(func() {
		_vapper = &Vapper{}
	})
	return _vapper
}

type Vapper struct {
	dispatcher flux.DispatcherInterface
	watcher    flux.WatcherInterface
	notifier   flux.NotifierInterface

	stores []flux.StoreInterface

	routes        []*route
	notFoundRoute *route
	cfg           reflect.Value

	// ShouldInterceptLinks tells the router whether or not to intercept click events
	// on links and call the Navigate method instead of the default behavior.
	// If it is set to true, the router will automatically intercept links when
	// Start, Navigate, or Back are called, or when the onpopstate event is triggered.
	ShouldInterceptLinks bool
	// ForceHashURL tells the router to use the hash component of the url to
	// represent different routes, even if history.pushState is supported.
	ForceHashURL bool
	// Verbose determines whether or not the router will log to console.log.
	// If true, the router will log a message if, e.g., a match cannot be found for
	// a particular path.
	Verbose bool
	// listener is the js.Object representation of a listener callback.
	// It is required in order to use the RemoveEventListener method
	listener func(event dom.Event)
}

// handleInject inject config and stores
func (t *Vapper) handleInject(_in interface{}) {
	defer xerror.Assert()

	_hn := reflect.ValueOf(_in)
	xerror.PanicT(!_hn.IsValid() || _hn.IsNil(), "func inject error")

	_Init := _hn.MethodByName("Init")
	if !_Init.IsValid() || _Init.IsNil() {
		return
	}

	var args []reflect.Value
	for i := 0; i < _Init.Type().NumIn(); i++ {
		_t := _Init.Type().In(i)

		if _t == reflect.TypeOf(t) {
			args = append(args, reflect.ValueOf(t))
		}

		if _t == t.cfg.Type() {
			args = append(args, t.cfg)
		}

		for j := 0; j < len(t.stores); j++ {
			if _t == reflect.TypeOf(t.stores[j]) {
				args = append(args, reflect.ValueOf(t.stores[j]))
				break
			}
		}
	}

	xerror.PanicT(_Init.Type().NumIn() != len(args), "inject params not match")
	_Init.Call(args)
}

// Start causes the router to listen for changes to window.location and
// trigger the appropriate handler whenever there is a change.
func (t *Vapper) Start() {

	// inject app,store,config
	for _, d := range t.routes {
		t.handleInject(d.handler)
	}

	for _, d := range t.stores {
		t.handleInject(d)
	}

	// watch store
	t.watch()

	// watch router
	if browserSupportsPushState && !t.ForceHashURL {
		t.pathChanged(getPath(), true)
		t.watchHistory()
	} else {
		t.setInitialHash()
		t.watchHash()
	}
	if t.ShouldInterceptLinks {
		t.InterceptLinks()
	}
}

// Stop causes the router to stop listening for changes, and therefore
// the router will not trigger any more router.Handler functions.
func Stop() {
	if browserSupportsPushState && !_vapper.ForceHashURL {
		dom.Window.RemoveEventListener("onpopstate", func(event dom.Event) {
			event.PreventDefault()

		})
	} else {
		dom.Window.RemoveEventListener("onhashchange", func(event dom.Event) {
			event.PreventDefault()
		})
	}
}
