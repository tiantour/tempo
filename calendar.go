package tempo

import (
	"fmt"
	"time"
)

// Month
func (i *iCalendar) Month(year, month int) []string {
	num := i.Day(year, month)
	result := []string{}
	for i := 1; i <= num; i++ {
		item := fmt.Sprintf("%d-%02d-%02d", year, month, i)
		result = append(result, item)
	}
	return result
}

// Day
func (i *iCalendar) Day(year, month int) int {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}

// Date
func (i *iCalendar) Date(date string, increment int) string {
	now := fmt.Sprintf("%s 00:00:00", date)
	current, _ := String.Time(now)
	return Time.Date(current.AddDate(0, 0, increment))
}

// Diff
func (i *iCalendar) Diff(start, end string) float64 {
	start = fmt.Sprintf("%s 00:00:00", start)
	end = fmt.Sprintf("%s 00:00:00", end)
	startTime, _ := String.Time(start)
	endTime, _ := String.Time(end)
	return startTime.Sub(endTime).Hours()
}

// StartDate
func (i *iCalendar) Start(t time.Time) string {
	year := Time.Year(t)
	month := Time.Month(t)
	return fmt.Sprintf("%d-%02d-01", year, month)
}

// EndDate
func (i *iCalendar) End(t time.Time) string {
	year := Time.Year(t)
	month := Time.Month(t)
	day := i.Day(year, month)
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
