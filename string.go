package tempo

import (
	"fmt"
	"time"
)

// Unix
func (i *iString) Unix(s string) (int64, error) {
	tempo, err := i.Time(s)
	return tempo.Unix(), err
}

// Date
func (i *iString) Date(s string) (string, error) {
	tempo, err := i.Time(s)
	year, mon, day := tempo.Date()
	return fmt.Sprintf("%02d-%02d-%02d", year, mon, day), err
}

// Clock
func (i *iString) Clock(s string) (string, error) {
	tempo, err := i.Time(s)
	hour, min, sec := tempo.Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec), err
}

// Year
func (i *iString) Year(s string) (int, error) {
	tempo, err := i.Time(s)
	return tempo.Year(), err
}

// MonthString
func (i *iString) MonthString(s string) (string, error) {
	tempo, err := i.Time(s)
	return tempo.Month().String(), err
}

// Month
func (i *iString) Month(s string) (int, error) {
	tempo, err := i.Time(s)
	return int(tempo.Month()), err
}

// Day
func (i *iString) Day(s string) (int, error) {
	tempo, err := i.Time(s)
	return tempo.Day(), err
}

// Hour
func (i *iString) Hour(s string) (int, error) {
	tempo, err := i.Time(s)
	return tempo.Hour(), err
}

// Minute
func (i *iString) Minute(s string) (int, error) {
	tempo, err := i.Time(s)
	return tempo.Minute(), err
}

// Second
func (i *iString) Second(s string) (int, error) {
	tempo, err := i.Time(s)
	return tempo.Second(), err
}

// TimeToWeekDay
func (i *iString) WeekDay(s string) (int, error) {
	tempo, err := i.Time(s)
	return int(tempo.Weekday()), err
}

// WeekDayString
func (i *iString) WeekDayString(s string) (string, error) {
	tempo, err := i.Time(s)
	return tempo.Weekday().String(), err
}

// Time
func (i *iString) Time(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
}
