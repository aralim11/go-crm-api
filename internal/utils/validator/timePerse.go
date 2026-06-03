package validator

import (
	"time"
)

func ParseDate(value string) time.Time {
	loc, _ := time.LoadLocation("Asia/Dhaka")

	formats := []string{
		"1/2/2006",
		"01/02/2006",
		"2006-01-02",
		"02-01-2006",
	}

	for _, f := range formats {
		if t, err := time.ParseInLocation(f, value, loc); err == nil {
			return t
		}
	}

	// fallback (zero value)
	return time.Time{}
}
