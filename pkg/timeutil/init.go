package timeutil

import "time"

// GoTime the different string formats for go dates
const (
	// DefaultFormat       = "2006-01-02 15:04:05"
	Format              = "2006-01-02 15:04:05"
	GoFormat            = "2006-01-02 15:04:05.999999999"
	DateFormat          = "2006-01-02"
	FormattedDateFormat = "Jan 2, 2006"
	HourMinuteFormat    = "15:04"
	HourFormat          = "15"
	DayDateTimeFormat   = "Mon, Aug 2, 2006 3:04 PM"
	CookieFormat        = "Monday, 02-Jan-2006 15:04:05 MST"
	RFC822Format        = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC1036Format       = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC2822Format       = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339Format       = "2006-01-02T15:04:05-07:00"
	RSSFormat           = "Mon, 02 Jan 2006 15:04:05 -0700"
	WeekStartDay        = time.Sunday
)

var (
	dtFormats   = []string{"2006-01-02T15:04", "2006-01-02T15:04:05", "2006-01-02 15:04:05", "2006-01-02 15:04"}
	TimeFormats = []string{"1/2/2006", "1/2/2006 15:4:5", "2006", "2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4", "2006-1-2 15:4:5", "1-2", "15:4:5", "15:4", "15", "15:4:5 Jan 2, 2006 MST", "2006-01-02 15:04:05.999999999 -0700 MST", "2006-01-02T15:04:05-07:00"}
)
