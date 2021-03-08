package typex

import "reflect"

type M map[string]interface{}

// StrOf string slice
func StrOf(s ...string) []string                 { return s }
func ObjOf(i ...interface{}) []interface{}       { return i }
func ValueOf(v ...reflect.Value) []reflect.Value { return v }

//https://github.com/emirpasic/gods
