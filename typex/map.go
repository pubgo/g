package typex

import (
	"reflect"

	"github.com/pubgo/x/fx"
	"github.com/pubgo/xerror"
)

var NotFound = new(interface{})

type Map map[string]interface{}

func (t Map) Each(fn interface{}) (err error) {
	defer xerror.RespErr(&err)

	xerror.Assert(fn == nil, "[fn] should not be nil")

	vfn := fx.WrapRaw(fn)
	onlyKey := reflect.TypeOf(fn).NumIn() == 1
	t.data.Range(func(key, value interface{}) bool {
		if onlyKey {
			_ = vfn(key)
			return true
		}

		_ = vfn(key, value)
		return true
	})

	return nil
}

func (t *Map) MapTo(data interface{}) (err error) {
	defer xerror.RespErr(&err)

	vd := reflect.ValueOf(data)
	if vd.Kind() == reflect.Ptr {
		vd = vd.Elem()
		vd.Set(reflect.MakeMap(vd.Type()))
	}

	// var data = make(map[string]int); MapTo(data)
	// var data map[string]int; MapTo(&data)
	xerror.Assert(!vd.IsValid() || vd.IsNil(), "[data] type error")

	t.data.Range(func(key, value interface{}) bool {
		vd.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
		return true
	})

	return nil
}

func (t *Map) Set(key, value interface{}) {
	_, ok := t.data.LoadOrStore(key, value)
	if !ok {
		t.count.Inc()
	}
}

func (t *Map) Load(key interface{}) (value interface{}, ok bool) { return t.data.Load(key) }
func (t *Map) Range(f func(key, value interface{}) bool)         { t.data.Range(f) }
func (t *Map) Len() int                                          { return int(t.count.Load()) }
func (t *Map) Delete(key interface{})                            { t.data.Delete(key); t.count.Dec() }
func (t Map) Get(key string) interface{} {
	value, ok := t[key]
	if ok {
		return value
	}
	return NotFound
}

func (t Map) Has(key string) bool {
	_, ok := t[key]
	return ok
}
