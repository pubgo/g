package byteutil

import (
	"unsafe"
)

// Copy copies a slice to make it immutable
func Copy(b []byte) []byte {
	tmp := make([]byte, len(b))
	copy(tmp, b)
	return tmp
}

// #nosec G103
// ToStr returns a string pointer without allocation
func ToStr(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
