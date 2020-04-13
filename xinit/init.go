package xinit

import "log"

var _inits []func() error

// Init init
func Init(fn func() (err error)) {
	if fn != nil {
		_inits = append(_inits, fn)
	}
}

// Notify init notify
func Notify() error {
	return Start()
}

// Start init start
func Start() (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln(r)
		}
	}()

	for _, fn := range _inits {
		if err = fn(); err != nil {
			return
		}
	}
	return
}
