package helper

import "regexp"

func IsValidDateFormat(date string) bool {
	datePattern := `^\d{4}-\d{2}-\d{2}$`
	matched, _ := regexp.MatchString(datePattern, date)
	return matched
}

func IsValidTimeFormat(time string) bool {
	//check time format from 00:00 to 23:59
	timePattern := `^(?:[01]\d|2[0-3]):[0-5]\d$`
	matched, _ := regexp.MatchString(timePattern, time)
	return matched
}
