package bytex

import (
	"strconv"
	"strings"
)

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
