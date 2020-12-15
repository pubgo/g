package pipe

import (
	"encoding/json"
	"github.com/pubgo/x/xerror"
	"reflect"
)

func (t *_func) ToRaw() []reflect.Value {
	return t.params
}

func (t *_func) ToString() string {
	return t.ToJSON()
}

func (t *_func) ToData(fn ...interface{}) interface{} {
	var _t reflect.Type
	for _, _v := range t.params {
		if _v.IsValid() {
			_t = _v.Type()
			break
		}
	}

	if _t == nil {
		return nil
	}

	for i := 0; i < len(t.params); i++ {
		if !t.params[i].IsValid() {
			t.params[i] = reflect.Zero(_t)
		}
	}

	_rst := reflect.MakeSlice(reflect.SliceOf(_t), 0, len(t.params))
	_rst = reflect.Append(_rst, t.params...)

	if len(fn) != 0 && !_isNone(fn[0]) && reflect.TypeOf(fn[0]).Kind() == reflect.Func {
		reflect.ValueOf(fn[0]).Call([]reflect.Value{_rst})
		return nil
	}

	return _rst.Interface()
}

func (t *_func) ToJSON() string {
	var _res []interface{}
	for _, _p := range t.params {
		_res = append(_res, _if(_isNone(_p), "", _p.Interface))
	}

	dt, err := json.Marshal(_res)
	xerror.PanicM(err, "data json error")

	return string(dt)
}
