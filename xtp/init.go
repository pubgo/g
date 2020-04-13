package xtp

import "reflect"

type StrM map[string]interface{}
type IntM map[int]interface{}

// StrOf string slice
func StrOf(s ...string) []string {
	return s
}

// IntOf int slice
func IntOf(i ...int) []int {
	return i
}

// IntOf int slice
func Int64Of(i ...int64) []int64 {
	return i
}

func ObjOf(i ...interface{}) []interface{} {
	return i
}

func Slice(data interface{}, i ...int) interface{} {
	_dt := reflect.ValueOf(data)
	if _dt.Len() == 0 {
		return nil
	}

	_rst := reflect.MakeSlice(reflect.SliceOf(_dt.Index(0).Type()), 0, _dt.Len())
	_rst = reflect.AppendSlice(_rst, _dt)
	return _rst.Interface()
}

//https://github.com/emirpasic/gods
