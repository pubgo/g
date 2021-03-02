package future

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/pubgo/x/abc"
	"github.com/stretchr/testify/assert"
)

type a1 struct {
	a int
}

func a1Test() *a1 {
	return &a1{}
}

func a2Test() *a1 {
	return nil
}

func a3Test() a1 {
	return a1{}
}

func TestAsync(t *testing.T) {
	val := Async(http.Get, "https://www.cnblogs.com")
	assert.Nil(t, val.Err())
	assert.NotNil(t, val.Get())
	assert.Nil(t, val.Value(func(resp *http.Response, err error) {
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	}))

	val = Async(func(i int) {}, 1)
	assert.Nil(t, val.Err())
	assert.Nil(t, val.Get())
	assert.Nil(t, val.Value(func() {}))

	a1Val := Await(Async(a1Test), func(a *a1) *a1 { a.a = 1; return a })
	assert.Nil(t, a1Val.Err())
	assert.NotNil(t, a1Val.Get().(*a1))
	assert.Equal(t, a1Val.Get().(*a1).a, 1)

	a1Val = Await(Async(a1Test), func(a *a1) *a1 { a.a = 1; return nil })
	assert.Nil(t, a1Val.Err())
	assert.Nil(t, a1Val.Get())
	assert.Nil(t, a1Val.Value(func(a *a1) {
		fmt.Println("a", a, a == nil)
	}))

	a2Val := Await(Async(a2Test), func(a *a1) *a1 {
		if a == nil {
			return nil
		}

		a.a = 1
		return a
	})
	assert.Nil(t, a2Val.Err())
	assert.Nil(t, a2Val.Get())

	a3Val := Await(Async(a3Test), func(a a1) a1 {
		a.a = 1
		return a
	})
	assert.Nil(t, a3Val.Err())
	assert.Equal(t, a3Val.Get().(a1).a, 1)
}

func TestAwait(t *testing.T) {
	val := Async(http.Get, "https://www.cnblogs.com")
	head := Await(val, func(resp *http.Response, err error) http.Header {
		assert.Nil(t, err)
		resp.Header.Add("aa", "11")
		return resp.Header
	})
	assert.Nil(t, head.Err())
	assert.NotNil(t, head.Get())
	assert.Equal(t, head.Get().(http.Header).Get("aa"), "11")
}

func promise1() abc.IPromise {
	return Promise(func(g abc.Future) {
		for i := 0; i < 10; i++ {
			val := Async(http.Get, "https://www.cnblogs.com")
			g.Yield(val)

			val = Async(http.Get, "https://www.cnblogs.com")
			val = Await(val, func(resp *http.Response, err error) (*http.Response, error) {
				resp.Header.Set("a", "b")
				return resp, err
			})

			g.YieldFn(val, func(resp *http.Response, err error) (*http.Response, error) {
				resp.Header.Set("b", "c")
				return resp, err
			})
		}
	})
}

func TestPromise(t *testing.T) {
	p := promise1()
	heads := p.Map(func(resp *http.Response, err error) http.Header { return resp.Header })
	assert.Nil(t, heads.Err())

	for _, h := range heads.Value().([]http.Header) {
		fmt.Println(h)
	}
}
