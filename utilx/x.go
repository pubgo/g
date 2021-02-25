package utilx

import (
	"reflect"
	"unsafe"
)

// If exported
func If(check bool, a, b interface{}) interface{} {
	if check {
		return a
	}
	return b
}

func Default(t, f interface{}) interface{} {
	if IsZero(t) {
		return f
	}
	return t
}

// IsZero
func IsZero(val interface{}) bool {
	return val == nil || reflect.ValueOf(val).IsZero()
}

// #nosec G103
// GetString returns a string pointer without allocation
func UnsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// #nosec G103
// GetBytes returns a byte pointer without allocation
func UnsafeBytes(s string) (bs []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return
}

// CopyString copies a string to make it immutable
func CopyString(s string) string {
	return string(UnsafeBytes(s))
}

// CopyBytes copies a slice to make it immutable
func CopyBytes(b []byte) []byte {
	tmp := make([]byte, len(b))
	copy(tmp, b)
	return tmp
}
