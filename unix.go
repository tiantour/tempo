package tempo

import (
	"fmt"
	"time"
)

// Unix unix
type Unix struct{}

// NewUnix new unix
func NewUnix() *Unix {
	return &Unix{}
}

// String
func (u Unix) String(n int64) string {
	return u.Time(n).Format("2006-01-02 15:04:05")
}

// Date get date
func (u Unix) Date(n int64) string {
	year, mon, day := u.Time(n).Date()
	return fmt.Sprintf("%d-%02d-%02d", year, mon, day)
}

// Clock get clock
func (u Unix) Clock(n int64) string {
	hour, min, sec := u.Time(n).Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}

// Year get year
func (u Unix) Year(n int64) int {
	return u.Time(n).Year()
}

// MonthString get month string
func (u Unix) MonthString(n int64) string {
	return u.Time(n).Month().String()
}

// Month get month
func (u Unix) Month(n int64) int {
	return int(u.Time(n).Month())
}

// Day get day
func (u Unix) Day(n int64) int {
	return u.Time(n).Day()
}

// Hour get hour
func (u Unix) Hour(n int64) int {
	return u.Time(n).Hour()
}

// Minute get minute
func (u Unix) Minute(n int64) int {
	return u.Time(n).Minute()
}

// Second get second
func (u Unix) Second(n int64) int {
	return u.Time(n).Second()
}

// WeekDay get time weekday
func (u Unix) WeekDay(n int64) int {
	return int(u.Time(n).Weekday())
}

// WeekDayString get weekday string
func (u Unix) WeekDayString(n int64) string {
	return u.Time(n).Weekday().String()
}

// Time get time
func (u Unix) Time(n int64) time.Time {
	return time.Unix(n, 0)
}
