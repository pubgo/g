package timeutil

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// GetSystemCurTime get system current time, nanoseconds value
func GetSystemCurTime() int {
	return int(time.Now().UnixNano() / 1e6)
}

// GetSystemCurDate 格式化日期时间 YYYY-MM-DD
func GetSystemCurDate() string {
	return time.Now().Format("2006-01-02")
}

func GetToday() (start time.Time, end time.Time) {
	loc, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		loc = time.FixedZone("CST", 8*3600)
	}
	time.Local = loc
	now := time.Now()
	start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end = start.AddDate(0, 0, 1)
	return
}

func GetYesterday() (start, end int) {
	loc, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		loc = time.FixedZone("CST", 8*3600)
	}
	time.Local = loc
	now := time.Now()
	endF := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	startF := endF.AddDate(0, 0, -1)
	end = int(endF.UnixNano() / 1e6)
	start = int(startF.UnixNano() / 1e6)
	return
}

func GetMonDay() int {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)

	return int(weekStart.UnixNano() / 1e6)
}

func GetSeveralDaysAgo(day int) (start, end int) {
	loc, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		loc = time.FixedZone("CST", 8*3600)
	}
	time.Local = loc
	now := time.Now()
	endF := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	startF := endF.AddDate(0, 0, -day)
	end = int(endF.UnixNano() / 1e6)
	start = int(startF.UnixNano() / 1e6)
	return
}

// A Time represents a time with nanosecond precision.
//
// This type does not include location information, and therefore does not
// describe a unique moment in time.
//
// This type exists to represent the TIME type in storage-based APIs like BigQuery.
// Most operations on Times are unlikely to be meaningful. Prefer the DateTime type.
type Time struct {
	Hour   int // The hour of the day in 24-hour format; range [0-23]
	Minute int // The minute of the hour; range [0-59]
	Valid  bool
}

// TimeOf returns the Time representing the time of day in which a time occurs
// in that time's location. It ignores the date.
func TimeOf(t time.Time) Time {
	tm := Time{Valid: !t.IsZero()}
	tm.Hour, tm.Minute, _ = t.Clock()
	return tm
}

// ParseTime parses a string and returns the time value it represents.
// ParseTime accepts an extended form of the RFC3339 partial-time format. After
// the HH:MM:SS part of the string, an optional fractional part may appear,
// consisting of a decimal point followed by one to nine decimal digits.
// (RFC3339 admits only one digit after the decimal point).
func ParseTime(s string) (Time, error) {
	t, err := time.Parse("15:04", s)
	if err != nil {
		t, err := time.Parse("15:04:05", s)
		return TimeOf(t), err
	}
	return TimeOf(t), nil
}

// String returns the date in the format described in ParseTime.
// If Valid is not true, it will return empty string
func (t Time) String() string {
	if t.Valid {
		return fmt.Sprintf("%02d:%02d", t.Hour, t.Minute)
	}
	return ""
}

// ToDate converts Time into time.Time
func (t Time) ToDate() time.Time {
	return time.Date(0, 0, 0, t.Hour, t.Minute, 0, 0, time.UTC)
}

// After checks if instance of t is after tm
func (t Time) After(tm Time) bool {
	if t.Hour > tm.Hour {
		return true
	}
	return t.Minute > tm.Minute
}

// Subtract returns difference between t and t2 in minutes
func (t Time) Subtract(t2 Time) int {
	return (t.Hour-t2.Hour)*60 + t.Minute - t2.Minute
}

// MarshalText implements the encoding.TextMarshaler interface.
// The output is the result of d.String().
func (t Time) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time is expected to be a string in a format accepted by ParseTime.
func (t *Time) UnmarshalText(data []byte) error {
	var err error
	*t, err = ParseTime(string(data))
	return err
}

// Value implements valuer interface
func (t Time) Value() (driver.Value, error) {
	if t.Valid {
		return driver.Value(t.String()), nil
	}
	return nil, nil
}

// Scan implements sql scan interface
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		tm, err := ParseTime(string(v))
		if err != nil {
			return err
		}
		*t = tm
		return nil
	case string:
		tm, err := ParseTime(v)
		if err != nil {
			return err
		}
		*t = tm
		return nil
	}
	return fmt.Errorf("Can't convert %T to Time", value)
}

//获取本地时间戳
func GetUtcTime(tm time.Time) int64 {
	return tm.Unix() //- 8*60*60
}

//当前时间向上取整点
func GetHour(timestamp int64) int {
	//	formaTime := time.Format("2006-01-02 15:04:05")
	tm := time.Unix(timestamp, 0)
	return tm.Hour()
}

//获取offset天的现在时间:注意时区
func GetLastDayCurrentTime(timestamp int64, offset int) time.Time {
	tm := time.Unix(timestamp, 0)
	yesDay := tm.AddDate(0, 0, 1*offset)
	return yesDay
}

//获取给定时间的星期
func GetTimeWeek(timestamp int64) int {
	tm := time.Unix(timestamp, 0)
	weekDay := tm.Weekday().String()
	var week int = 0
	switch weekDay {
	case "Monday":
		week = 1
	case "Tuesday":
		week = 2
	case "Wednesday":
		week = 3
	case "Thursday":
		week = 4
	case "Friday":
		week = 5
	case "Saturday":
		week = 6
	default:
		week = 0
	}
	return week
}

//获取向上整时时间
func GetHour0(timestamp int64) time.Time {
	tm := time.Unix(timestamp, 0)
	tStr := tm.Format("2006-01-02 15") + ":00:00"
	return StrToTime(tStr, "2006-01-02 15:04:05", nil)
}

//获取给定日期的零点时间
func GetDay0(timestamp int64) time.Time {
	tm := time.Unix(timestamp, 0)
	tStr := tm.Format("2006-01-02") + " 00:00:00"

	return StrToTime(tStr, "2006-01-02 15:04:05", nil)
}

//获取offset 0点时间
func GetUtcDay0(now time.Time, timeZone *time.Location) int64 {
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return tm.Unix()
}

//字符串转时间
func StrToTime(tStr, format string, timeZone *time.Location) time.Time {
	if len(format) == 0 {
		format = "2006-01-02 15:04:05"
	}
	if timeZone == nil {
		//chinaLocal, _ := time.LoadLocation("Local")
		timeZone = time.Local
	}

	ti, _ := time.ParseInLocation(format, tStr, timeZone)
	return ti
}

/*
	给定字符串时间转换成本地时间戳
*/
func StringTimetoUnix(timestr string) int64 {
	return StrToTime(timestr, "2006-01-02 15:04:05", time.Local).Unix()
}

//获取最近上个星期天的零点日期
func GetWeek0(timestamp int64) time.Time {
	weekday := GetTimeWeek(timestamp)
	tm0 := GetDay0(timestamp)
	return tm0.AddDate(0, 0, -1*weekday)
}

//获取最近上个星期天的零点日期
func GetUtcWeek0(timestamp int64) int64 {
	weekday := GetTimeWeek(timestamp)
	tm0 := GetDay0(timestamp)
	tm0 = tm0.AddDate(0, 0, -1*weekday)

	return tm0.Unix()
}

/*
	获取给定时间的当月1号零点时间
*/
func GetMonth0(timestamp int64) time.Time {

	tm0 := GetDay0(timestamp)
	month0 := tm0.Day() - 1
	tm0 = tm0.AddDate(0, 0, -1*month0) //这个月1号
	return tm0
}

//整点执行操作
func TimerByHour(f func()) {
	for {
		now := time.Now()
		// 计算下一个整点
		next := now.Add(time.Hour * 1)
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
		//以下为定时执行的操作
		f()
	}
}

//时间戳转换为time
func UnixToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

//获取本地时间
func GetLocalTime(tm time.Time) time.Time {
	local, _ := time.LoadLocation("Local")
	return tm.In(local)
	//return tm.Add(8 * 60 * 60 * time.Second)
}

//获取系统时间的格式
func GetSysTimeLayout() string {
	t := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	strLayout := strings.Replace(t.String(), "+0000 UTC", "", -1)
	return strings.TrimSpace(strLayout)
}

func FormatTime(tm time.Time, for_str string) string {
	return tm.Format(for_str)
}

func GetTimeStr(tm time.Time) string {
	return FormatTime(tm, "2006-01-02 15:04:05")
}

func GetDayStr(tm time.Time) string {
	return FormatTime(tm, "2006-01-02")
}

// 秒时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// 毫秒时间戳
func NowTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// 毫秒时间戳
func Timestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 秒时间戳转时间
func TimeFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

// 毫秒时间戳转时间
func TimeFromTimestamp(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(time.Millisecond))
}

// 时间格式化
func TimeFormat(time time.Time, layout string) string {
	return time.Format(layout)
}

// 字符串时间转时间类型
func TimeParse(timeStr, layout string) (time.Time, error) {
	return time.Parse(layout, timeStr)
}

// return yyyyMMdd
func GetDay(time time.Time) int {
	ret, _ := strconv.Atoi(time.Format("20060102"))
	return ret
}

// 返回指定时间当天的开始时间
func WithTimeAsStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
