package tempo

import (
	"fmt"
	"time"
)

// Now now
type Now struct{}

// NewNow new now
func NewNow() *Now {
	return &Now{}
}

// Unix get unix
// date 2017-05-22
// author andy.jiang
func (n Now) Unix() int64 {
	return n.Time().Unix()
}

// String get string
// date 2017-05-22
// author andy.jiang
func (n Now) String() string {
	return n.Time().Format("2006-01-02 15:04:05")
}

// Date get date
// date 2017-05-22
// author andy.jiang
func (n Now) Date() string {
	year, mon, day := n.Time().Date()
	return fmt.Sprintf("%d-%02d-%02d", year, mon, day)
}

// Clock get clock
// date 2017-05-22
// author andy.jiang
func (n Now) Clock() string {
	hour, min, sec := n.Time().Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}

// Year get year
// date 2017-05-22
// author andy.jiang
func (n Now) Year() int {
	return n.Time().Year()
}

// MonthString get month string
// date 2017-05-22
// author andy.jiang
func (n Now) MonthString() string {
	return n.Time().Month().String()
}

// Month get month
// date 2017-05-22
// author andy.jiang
func (n Now) Month() int {
	return int(n.Time().Month())
}

// Day get day
// date 2017-05-22
// author andy.jiang
func (n Now) Day() int {
	return n.Time().Day()
}

// Hour get hour
// date 2017-05-22
// author andy.jiang
func (n Now) Hour() int {
	return n.Time().Hour()
}

// Minute get minute
// date 2017-05-22
// author andy.jiang
func (n Now) Minute() int {
	return n.Time().Minute()
}

// Second get second
// date 2017-05-22
// author andy.jiang
func (n Now) Second() int {
	return n.Time().Second()
}

// WeekDayString get weekday string
// date 2017-05-22
// author andy.jiang
func (n Now) WeekDayString() string {
	return n.Time().Weekday().String()
}

// WeekDay get weekday
// date 2017-05-22
// author andy.jiang
func (n Now) WeekDay() int {
	return int(n.Time().Weekday())
}

// Time get time
// date 2017-05-22
// author andy.jiang
func (n Now) Time() time.Time {
	return time.Now()
}
