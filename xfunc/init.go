package xfunc

import (
	"errors"
	"github.com/pubgo/x/pkg"
	"log"
	"reflect"
	"time"
)

func WithTimeout(dur time.Duration, fn func() error) error {
	var ch = make(chan error, 1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case error:
					ch <- err.(error)
				default:
					log.Fatalln(err)
				}
			}
		}()
		ch <- fn()
	}()

	select {
	case err := <-ch:
		return err
	case <-time.After(dur):
		return errors.New("函数执行超时: " + pkg.Caller.FromFunc(reflect.ValueOf(fn)))
	}
}
