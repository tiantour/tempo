package tempo

import (
	"fmt"
	"time"
)

// String string
type String struct{}

// NewString new string
func NewString() *String {
	return &String{}
}

// Unix get unix
func (s String) Unix(str string) (int64, error) {
	tempo, err := s.Time(str)
	return tempo.Unix(), err
}

// Date get date
func (s String) Date(str string) (string, error) {
	tempo, err := s.Time(str)
	year, mon, day := tempo.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, mon, day), err
}

// Clock get clock
func (s String) Clock(str string) (string, error) {
	tempo, err := s.Time(str)
	hour, min, sec := tempo.Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec), err
}

// Year get year
func (s String) Year(str string) (int, error) {
	tempo, err := s.Time(str)
	return tempo.Year(), err
}

// MonthString get month string
func (s String) MonthString(str string) (string, error) {
	tempo, err := s.Time(str)
	return tempo.Month().String(), err
}

// Month get month
func (s String) Month(str string) (int, error) {
	tempo, err := s.Time(str)
	return int(tempo.Month()), err
}

// Day get day
func (s String) Day(str string) (int, error) {
	tempo, err := s.Time(str)
	return tempo.Day(), err
}

// Hour get hour
func (s String) Hour(str string) (int, error) {
	tempo, err := s.Time(str)
	return tempo.Hour(), err
}

// Minute get minute
func (s String) Minute(str string) (int, error) {
	tempo, err := s.Time(str)
	return tempo.Minute(), err
}

// Second get second
func (s String) Second(str string) (int, error) {
	tempo, err := s.Time(str)
	return tempo.Second(), err
}

// WeekDay get weekday
func (s String) WeekDay(str string) (int, error) {
	tempo, err := s.Time(str)
	return int(tempo.Weekday()), err
}

// WeekDayString get weekday string
func (s String) WeekDayString(str string) (string, error) {
	tempo, err := s.Time(str)
	return tempo.Weekday().String(), err
}

// Time get time
func (s String) Time(str string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
}
