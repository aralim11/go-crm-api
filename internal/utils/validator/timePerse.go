package validator

import "time"

func ParseDate(value string) time.Time {
	// timezone
	loc, _ := time.LoadLocation("Asia/Dhaka")

	// try main format first, value, timezone
	convertTime, err := time.ParseInLocation("2006-01-02", value, loc)
	if err == nil {
		return convertTime
	}

	// fallback format
	return convertTime
}
