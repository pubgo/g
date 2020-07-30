package xdi

import (
	"github.com/pubgo/x/xinit"
	"go.uber.org/dig"
)

type In = dig.In
type Out = dig.Out

var _dig = dig.New()

func Invoke(function interface{}, opts ...dig.InvokeOption) error {
	return _dig.Invoke(function, opts...)
}

func String() string {
	return _dig.String()
}

func InitProvide(constructor interface{}, opts ...dig.ProvideOption) {
	if err := _dig.Provide(constructor, opts...); err != nil {
		panic(err)
	}
}

func InitInvoke(function interface{}, opts ...dig.InvokeOption) {
	xinit.Go(func() error {
		return _dig.Invoke(function, opts...)
	})
}
