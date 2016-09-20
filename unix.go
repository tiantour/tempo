package tempo

import (
	"fmt"
	"time"
)

// String
func (i *iUnix) String(n int64) string {
	return i.Time(n).Format("2006-01-02 15:04:05")
}

// Date
func (i *iUnix) Date(n int64) string {
	year, mon, day := i.Time(n).Date()
	return fmt.Sprintf("%02d-%02d-%02d", year, mon, day)
}

// Clock
func (i *iUnix) Clock(n int64) string {
	hour, min, sec := i.Time(n).Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}

// Year
func (i *iUnix) Year(n int64) int {
	return i.Time(n).Year()
}

// MonthString
func (i *iUnix) MonthString(n int64) string {
	return i.Time(n).Month().String()
}

// Month
func (i *iUnix) Month(n int64) int {
	return int(i.Time(n).Month())
}

// Day
func (i *iUnix) Day(n int64) int {
	return i.Time(n).Day()
}

// Hour
func (i *iUnix) Hour(n int64) int {
	return i.Time(n).Hour()
}

// Minute
func (i *iUnix) Minute(n int64) int {
	return i.Time(n).Minute()
}

// Second
func (i *iUnix) Second(n int64) int {
	return i.Time(n).Second()
}

// TimeToWeekDay
func (i *iUnix) WeekDay(n int64) int {
	return int(i.Time(n).Weekday())
}

// WeekDayString
func (i *iUnix) WeekDayString(n int64) string {
	return i.Time(n).Weekday().String()
}

// Time
func (i *iUnix) Time(n int64) time.Time {
	return time.Unix(n, 0)
}
