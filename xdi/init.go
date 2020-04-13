package xdi

import (
	"github.com/pubgo/g/xinit"
	"go.uber.org/dig"
)

type In = dig.In
type Out = dig.Out

var _dig = dig.New()

func Provide(constructor interface{}, opts ...dig.ProvideOption) error {
	return _dig.Provide(constructor, opts...)
}

func Invoke(function interface{}, opts ...dig.InvokeOption) error {
	return _dig.Invoke(function, opts...)
}

func String() string {
	return _dig.String()
}

func InitProvide(constructor interface{}, opts ...dig.ProvideOption) {
	xinit.Init(func() error {
		return _dig.Provide(constructor, opts...)
	})
}

func InitInvoke(function interface{}, opts ...dig.InvokeOption) {
	xinit.Init(func() error {
		return _dig.Invoke(function, opts...)
	})
}
