package tempo

import (
	"fmt"
	"time"
)

// Unix
func (i *iNow) Unix() int64 {
	return i.Time().Unix()
}

// String
func (i *iNow) String() string {
	return i.Time().Format("2006-01-02 15:04:05")
}

// Date
func (i *iNow) Date() string {
	year, mon, day := i.Time().Date()
	return fmt.Sprintf("%d-%02d-%02d", year, mon, day)
}

// Clock
func (i *iNow) Clock() string {
	hour, min, sec := i.Time().Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}

// Year
func (i *iNow) Year() int {
	return i.Time().Year()
}

// MonthString
func (i *iNow) MonthString() string {
	return i.Time().Month().String()
}

// Month
func (i *iNow) Month() int {
	return int(i.Time().Month())
}

// Day
func (i *iNow) Day() int {
	return i.Time().Day()
}

// Hour
func (i *iNow) Hour() int {
	return i.Time().Hour()
}

// Minute
func (i *iNow) Minute() int {
	return i.Time().Minute()
}

// Second
func (i *iNow) Second() int {
	return i.Time().Second()
}

// WeekDayString
func (i *iNow) WeekDayString() string {
	return i.Time().Weekday().String()
}

// WeekDay
func (i *iNow) WeekDay() int {
	return int(i.Time().Weekday())
}

// Time
func (i *iNow) Time() time.Time {
	return time.Now()
}
