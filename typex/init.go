package typex

import (
	"reflect"
)

type M map[string]interface{}

// StrOf string slice
func StrOf(s1 string, ss ...string) []string {
	var s = make([]string, 0, len(ss)+1)
	return append(append(s, s1), ss...)
}

// ObjOf object slice
func ObjOf(s1 interface{}, ss ...interface{}) []interface{} {
	var s = make([]interface{}, 0, len(ss)+1)
	return append(append(s, s1), ss...)
}

func ValueOf(s1 reflect.Value, ss ...reflect.Value) []reflect.Value {
	var s = make([]reflect.Value, 0, len(ss)+1)
	return append(append(s, s1), ss...)
}

//https://github.com/emirpasic/gods
