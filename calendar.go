package tempo

import (
	"fmt"
	"time"
)

// Calendar calendar
type Calendar struct{}

// NewCalendar new calendar
func NewCalendar() *Calendar {
	return &Calendar{}
}

// Month get month
func (c *Calendar) Month(year, month int) []string {
	num := c.Day(year, month)
	result := []string{}
	for i := 1; i <= num; i++ {
		item := fmt.Sprintf("%d-%02d-%02d", year, month, i)
		result = append(result, item)
	}
	return result
}

// Day get day
func (c *Calendar) Day(year, month int) int {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}

// Date get date
func (c *Calendar) Date(date string, increment int) string {
	now := fmt.Sprintf("%s 00:00:00", date)
	current, _ := NewString().Time(now)
	return NewTime().Date(current.AddDate(0, 0, increment))
}

// Diff get diff
func (c *Calendar) Diff(start, end string) float64 {
	start = fmt.Sprintf("%s 00:00:00", start)
	end = fmt.Sprintf("%s 00:00:00", end)
	startTime, _ := NewString().Time(start)
	endTime, _ := NewString().Time(end)
	return startTime.Sub(endTime).Hours()
}

// Start get start date
func (c *Calendar) Start(t time.Time) string {
	year := NewTime().Year(t)
	month := NewTime().Month(t)
	return fmt.Sprintf("%d-%02d-01", year, month)
}

// End get end date
func (c *Calendar) End(t time.Time) string {
	year := NewTime().Year(t)
	month := NewTime().Month(t)
	day := c.Day(year, month)
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
