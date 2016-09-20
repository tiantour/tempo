package tempo

import (
	"fmt"
	"time"
)

// Day
func (i *iCalendar) Day(year, month int) int {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}

// Month
func (i *iCalendar) Month(year, month int) []map[string]interface{} {
	num := i.Day(year, month)
	result := []map[string]interface{}{}
	for day := 1; day <= num; day++ {
		item := map[string]interface{}{
			"date": fmt.Sprintf("%d-%02d-%02d", year, month, day),
		}
		result = append(result, item)
	}
	return result
}
