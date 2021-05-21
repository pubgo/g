package byteutil

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	_ = 1.0 << (10 * iota) // ignore first value by assigning to blank identifier
	KB
	MB
	GB
	TB
	PB
	EB
)

type (
	// bytesFormat struct
	bytesFormat struct{}
)

var (
	pattern = regexp.MustCompile(`(?i)^(-?\d+(?:\.\d+)?)([KMGTPE]B?|B?)$`)
	global  bytesFormat
)

// Format formats bytesFormat integer to human readable string.
// For example, 31323 bytesFormat will return 30.59KB.
func (bytesFormat) Format(b int64) string {
	multiple := ""
	value := float64(b)

	switch {
	case b >= EB:
		value /= EB
		multiple = "EB"
	case b >= PB:
		value /= PB
		multiple = "PB"
	case b >= TB:
		value /= TB
		multiple = "TB"
	case b >= GB:
		value /= GB
		multiple = "GB"
	case b >= MB:
		value /= MB
		multiple = "MB"
	case b >= KB:
		value /= KB
		multiple = "KB"
	case b == 0:
		return "0"
	default:
		return strconv.FormatInt(b, 10) + "B"
	}

	return fmt.Sprintf("%.2f%s", value, multiple)
}

// Parse parses human readable bytesFormat string to bytesFormat integer.
// For example, 6GB (6G is also valid) will return 6442450944.
func (bytesFormat) Parse(value string) (i int64, err error) {
	parts := pattern.FindStringSubmatch(value)
	if len(parts) < 3 {
		return 0, fmt.Errorf("error parsing value=%s", value)
	}

	bytesString := parts[1]
	multiple := parts[2]
	bytes, err := strconv.ParseFloat(bytesString, 64)
	if err != nil {
		return
	}

	switch multiple {
	default:
		return int64(bytes), nil
	case "K", "KB":
		return int64(bytes * KB), nil
	case "M", "MB":
		return int64(bytes * MB), nil
	case "G", "GB":
		return int64(bytes * GB), nil
	case "T", "TB":
		return int64(bytes * TB), nil
	case "P", "PB":
		return int64(bytes * PB), nil
	case "E", "EB":
		return int64(bytes * EB), nil
	}
}

// Format wraps global bytesFormat's Format function.
func Format(b int64) string {
	return global.Format(b)
}

// Parse wraps global bytesFormat's Parse function.
func Parse(val string) (int64, error) {
	return global.Parse(val)
}
