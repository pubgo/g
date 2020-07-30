package xtry_test

import (
	"github.com/pubgo/x/xtry"
	"github.com/pubgo/xerror"
	"testing"
)

func BenchmarkTry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() (err error) {
			defer xerror.RespErr(&err)
			xerror.Panic(xtry.Try(func() error {
				return nil
			}))
			return
		}()

		xtry.Try(func() error {
			return xtry.ErrParamsNotMatch
		})
		xtry.Try(func() error {
			panic(xtry.ErrParamsNotMatch)
			return nil
		})
	}
}

func BenchmarkPipe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xtry.Pipe(func() error {
			return nil
		}, func() error {
			panic(xtry.ErrParamsNotMatch)
			return nil
		})
	}
}
