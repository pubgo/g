package byteutil

import (
	"strconv"
	"strings"
)

// ToLowerBytes is the equivalent of bytes.ToLower
func ToLowerBytes(b []byte) []byte {
	for i := 0; i < len(b); i++ {
		b[i] = toLowerTable[b[i]]
	}
	return b
}

// ToUpperBytes is the equivalent of bytes.ToUpper
func ToUpperBytes(b []byte) []byte {
	for i := 0; i < len(b); i++ {
		b[i] = toUpperTable[b[i]]
	}
	return b
}

// TrimRightBytes is the equivalent of bytes.TrimRight
func TrimRightBytes(b []byte, c byte) []byte {
	lenStr := len(b)
	for lenStr > 0 && b[lenStr-1] == c {
		lenStr--
	}
	return b[:lenStr]
}

// TrimLeftBytes is the equivalent of bytes.TrimLeft
func TrimLeftBytes(b []byte, c byte) []byte {
	lenStr, start := len(b), 0
	for start < lenStr && b[start] == c {
		start++
	}
	return b[start:]
}

// TrimBytes is the equivalent of bytes.Trim
func TrimBytes(b []byte, c byte) []byte {
	i, j := 0, len(b)-1
	for ; i < j; i++ {
		if b[i] != c {
			break
		}
	}
	for ; i < j; j-- {
		if b[j] != c {
			break
		}
	}

	return b[i : j+1]
}

// EqualFold the equivalent of bytes.EqualFold
func EqualFoldBytes(b, s []byte) (equals bool) {
	n := len(b)
	equals = n == len(s)
	if equals {
		for i := 0; i < n; i++ {
			if equals = b[i]|0x20 == s[i]|0x20; !equals {
				break
			}
		}
	}
	return
}


const (
	uByte = 1 << (10 * iota)
	uKilobyte
	uMegabyte
	uGigabyte
	uTerabyte
	uPetabyte
	uExabyte
)

// ByteSize returns a human-readable byte string of the form 10M, 12.5K, and so forth.
// The unit that results in the smallest number greater than or equal to 1 is always chosen.
func ByteSize(bytes uint64) string {
	unit := ""
	value := float64(bytes)
	switch {
	case bytes >= uExabyte:
		unit = "EB"
		value = value / uExabyte
	case bytes >= uPetabyte:
		unit = "PB"
		value = value / uPetabyte
	case bytes >= uTerabyte:
		unit = "TB"
		value = value / uTerabyte
	case bytes >= uGigabyte:
		unit = "GB"
		value = value / uGigabyte
	case bytes >= uMegabyte:
		unit = "MB"
		value = value / uMegabyte
	case bytes >= uKilobyte:
		unit = "KB"
		value = value / uKilobyte
	case bytes >= uByte:
		unit = "B"
	default:
		return "0B"
	}
	result := strconv.FormatFloat(value, 'f', 1, 64)
	result = strings.TrimSuffix(result, ".0")
	return result + unit
}
