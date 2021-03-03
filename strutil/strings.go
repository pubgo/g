package strutil

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"unsafe"
)

var (
	ErrInvalidStartPosition = errors.New("start position is invalid")
	ErrInvalidStopPosition  = errors.New("stop position is invalid")
)

func Contains(list []string, str string) bool {
	for _, each := range list {
		if each == str {
			return true
		}
	}

	return false
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

func HasEmpty(args ...string) bool {
	for _, arg := range args {
		if len(arg) == 0 {
			return true
		}
	}

	return false
}

func NotEmpty(args ...string) bool {
	return !HasEmpty(args...)
}

// Substr returns runes between start and stop [start, stop) regardless of the chars are ascii or utf8
func Substr(str string, start int, stop int) (string, error) {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return "", ErrInvalidStartPosition
	}

	if stop < 0 || stop > length {
		return "", ErrInvalidStopPosition
	}

	return string(rs[start:stop]), nil
}

func TakeOne(valid, or string) string {
	if len(valid) > 0 {
		return valid
	} else {
		return or
	}
}

func TakeWithPriority(fns ...func() string) string {
	for _, fn := range fns {
		val := fn()
		if len(val) > 0 {
			return val
		}
	}

	return ""
}

func Union(first, second []string) []string {
	set := make(map[string]struct{})

	for _, each := range first {
		set[each] = struct{}{}
	}
	for _, each := range second {
		set[each] = struct{}{}
	}

	merged := make([]string, 0, len(set))
	for k := range set {
		merged = append(merged, k)
	}

	return merged
}

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

func Find(ss []string, str string) int {
	for i, s := range ss {
		if str == s {
			return i
		}
	}
	return -1
}

// Contain
// determines whether the str is in the ss.
func Contain(ss []string, str string) bool { return Find(ss, str) > -1 }

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
	excludeMap := make(map[string]struct{})
	for _, s := range exclude {
		excludeMap[s] = struct{}{}
	}

	for _, s := range base {
		if _, ok := excludeMap[s]; !ok {
			result = append(result, s)
		}
	}
	return result
}

// Unique
func Unique(ss []string) (result []string) {
	_map := make(map[string]struct{})
	for _, s := range ss {
		_map[s] = struct{}{}
	}

	result = make([]string, len(_map))
	for s := range _map {
		result = append(result, s)
	}
	return result
}

// Addr2Hex converts address string to hex string, only support ipv4.
func Addr2Hex(str string) (string, error) {
	ipStr, portStr, err := net.SplitHostPort(str)
	if err != nil {
		return "", err
	}

	ip := net.ParseIP(ipStr).To4()
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return "", nil
	}

	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(port))
	ip = append(ip, buf...)

	return hex.EncodeToString(ip), nil
}

// Hex2Addr converts hex string to address.
func Hex2Addr(str string) (string, error) {
	buf, err := hex.DecodeString(str)
	if err != nil {
		return "", err
	}
	if len(buf) < 4 {
		return "", fmt.Errorf("bad hex string length")
	}
	return fmt.Sprintf("%s:%d", net.IP(buf[:4]).String(), binary.BigEndian.Uint16(buf[4:])), nil
}

// Strings ...
type Strings []string

// KickEmpty kick empty elements from ss
func KickEmpty(ss []string) Strings {
	var ret = make([]string, 0)
	for _, str := range ss {
		if str != "" {
			ret = append(ret, str)
		}
	}
	return Strings(ret)
}

// AnyBlank return true if ss has empty element
func AnyBlank(ss []string) bool {
	for _, str := range ss {
		if str == "" {
			return true
		}
	}

	return false
}

// HeadT ...
func (ss Strings) HeadT() (string, Strings) {
	if len(ss) > 0 {
		return ss[0], Strings(ss[1:])
	}

	return "", Strings{}
}

// Head ...
func (ss Strings) Head() string {
	if len(ss) > 0 {
		return ss[0]
	}
	return ""
}

// Head2 ...
func (ss Strings) Head2() (h0, h1 string) {
	if len(ss) > 0 {
		h0 = ss[0]
	}
	if len(ss) > 1 {
		h1 = ss[1]
	}
	return
}

// Head3 ...
func (ss Strings) Head3() (h0, h1, h2 string) {
	if len(ss) > 0 {
		h0 = ss[0]
	}
	if len(ss) > 1 {
		h1 = ss[1]
	}
	if len(ss) > 2 {
		h2 = ss[2]
	}
	return
}

// Head4 ...
func (ss Strings) Head4() (h0, h1, h2, h3 string) {
	if len(ss) > 0 {
		h0 = ss[0]
	}
	if len(ss) > 1 {
		h1 = ss[1]
	}
	if len(ss) > 2 {
		h2 = ss[2]
	}
	if len(ss) > 3 {
		h3 = ss[3]
	}
	return
}

// Split ...
func Split(raw string, sep string) Strings {
	return Strings(strings.Split(raw, sep))
}
