package tempo

import (
	"fmt"
	"time"
)

// Time time
type Time struct{}

// NewTime new time
func NewTime() *Time {
	return &Time{}
}

// Date get date
// date 2017-05-22
// author andy.jiang
func (t Time) Date(tempus time.Time) string {
	year, mon, day := tempus.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, mon, day)
}

// Clock get clock
// date 2017-05-22
// author andy.jiang
func (t Time) Clock(tempus time.Time) string {
	hour, min, sec := tempus.Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}

// Year get year
// date 2017-05-22
// author andy.jiang
func (t Time) Year(tempus time.Time) int {
	return tempus.Year()
}

// MonthString get month string
// date 2017-05-22
// author andy.jiang
func (t Time) MonthString(tempus time.Time) string {
	return tempus.Month().String()
}

// Month get month
// date 2017-05-22
// author andy.jiang
func (t Time) Month(tempus time.Time) int {
	return int(tempus.Month())
}

// Day get day
// date 2017-05-22
// author andy.jiang
func (t Time) Day(tempus time.Time) int {
	return tempus.Day()
}

// Hour get hour
// date 2017-05-22
// author andy.jiang
func (t Time) Hour(tempus time.Time) int {
	return tempus.Hour()
}

// Minute get minute
// date 2017-05-22
// author andy.jiang
func (t Time) Minute(tempus time.Time) int {
	return tempus.Minute()
}

// Second get second
// date 2017-05-22
// author andy.jiang
func (t Time) Second(tempus time.Time) int {
	return tempus.Second()
}

// WeekDayString get weekday string
// date 2017-05-22
// author andy.jiang
func (t Time) WeekDayString(tempus time.Time) string {
	return tempus.Weekday().String()
}

// WeekDay get weekday
// date 2017-05-22
// author andy.jiang
func (t Time) WeekDay(tempus time.Time) int {
	return int(tempus.Weekday())
}
