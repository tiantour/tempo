package tempo

import (
	"fmt"
	"time"
)

// Date
func (i *iTime) Date(tempus time.Time) string {
	year, mon, day := tempus.Date()
	return fmt.Sprintf("%02d-%02d-%02d", year, mon, day)
}

// Clock
func (i *iTime) Clock(tempus time.Time) string {
	hour, min, sec := tempus.Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}

// Year
func (i *iTime) Year(tempus time.Time) int {
	return tempus.Year()
}

// MonthString
func (i *iTime) MonthString(tempus time.Time) string {
	return tempus.Month().String()
}

// Month
func (i *iTime) Month(tempus time.Time) int {
	return int(tempus.Month())
}

// Day
func (i *iTime) Day(tempus time.Time) int {
	return tempus.Day()
}

// Hour
func (i *iTime) Hour(tempus time.Time) int {
	return tempus.Hour()
}

// Minute
func (i *iTime) Minute(tempus time.Time) int {
	return tempus.Minute()
}

// Second
func (i *iTime) Second(tempus time.Time) int {
	return tempus.Second()
}

// WeekDayString
func (i *iTime) WeekDayString(tempus time.Time) string {
	return tempus.Weekday().String()
}

// WeekDay
func (i *iTime) WeekDay(tempus time.Time) int {
	return int(tempus.Weekday())
}
