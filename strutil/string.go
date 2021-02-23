package strutil

import (
	"strconv"
	"unsafe"
)

// ToStr
// converts the specified byte array to a string.
func ToStr(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// ToBytes
// converts the specified str to a byte array.
func ToBytes(str string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Filter(ss []string, str string) (dt []string) {
	for _, s := range ss {
		if s == str {
			continue
		}

		dt = append(dt, s)
	}
	return
}

func Find(ss []string, str string) int {
	for i, s := range ss {
		if str == s {
			return i
		}
	}
	return -1
}

// Contains
// determines whether the str is in the strs.
func Contains(ss []string, str string) bool {
	return Find(ss, str) > -1
}

// Remove
// remove item from slice
func Remove(ss []string, s string) (result []string) {
	for _, item := range ss {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

func ToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func ToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func ToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// SubStr
// 截取字符串
func SubStr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

// Reverse
// returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Diff
// Creates an slice of slice values not included in the other given slice.
func Diff(base, exclude []string) (result []string) {
	excludeMap := make(map[string]bool)
	for _, s := range exclude {
		excludeMap[s] = true
	}
	for _, s := range base {
		if !excludeMap[s] {
			result = append(result, s)
		}
	}
	return result
}

// Unique
func Unique(ss []string) (result []string) {
	_map := make(map[string]bool)
	for _, s := range ss {
		_map[s] = true
	}

	result = make([]string, len(_map))
	for s := range _map {
		result = append(result, s)
	}
	return result
}
