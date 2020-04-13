package pkg

import (
	"reflect"
	"time"
)

// If exported
func If(b bool, t, f interface{}) interface{} {
	if b {
		return t
	}
	return f
}

func Default(t, f interface{}) interface{} {
	if IsNone(t) {
		return f
	}
	return t
}

// IsNone exported
func IsNone(val interface{}) bool {
	return val == nil || IsZero(reflect.ValueOf(val))
}

// IsZero exported
func IsZero(val reflect.Value) bool {
	if !val.IsValid() {
		return true
	}

	switch val.Kind() {
	case reflect.String:
		return val.Len() == 0
	case reflect.Bool:
		return val.Bool() == false
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint() == 0
	case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface, reflect.Slice, reflect.Map:
		return val.IsNil()
	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if !IsZero(val.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Struct:
		if t, ok := val.Interface().(time.Time); ok {
			return t.IsZero()
		}

		valid := val.FieldByName("Valid")
		if valid.IsValid() {
			va, ok := valid.Interface().(bool)
			return ok && !va
		}

		return reflect.DeepEqual(val.Interface(), reflect.Zero(val.Type()).Interface())
	default:
		return reflect.DeepEqual(val.Interface(), reflect.Zero(val.Type()).Interface())
	}
}

// LCS gets the longest common substring of s1 and s2.
// 最长公共子串
// Refers to http://en.wikibooks.org/wiki/Algorithm_Implementation/Strings/Longest_common_substring.
func LCS(s1 string, s2 string) string {
	var m = make([][]int, 1+len(s1))

	for i := 0; i < len(m); i++ {
		m[i] = make([]int, 1+len(s2))
	}

	longest := 0
	xLongest := 0

	for x := 1; x < 1+len(s1); x++ {
		for y := 1; y < 1+len(s2); y++ {
			if s1[x-1] == s2[y-1] {
				m[x][y] = m[x-1][y-1] + 1
				if m[x][y] > longest {
					longest = m[x][y]
					xLongest = x
				}
			} else {
				m[x][y] = 0
			}
		}
	}

	return s1[xLongest-longest : xLongest]
}
