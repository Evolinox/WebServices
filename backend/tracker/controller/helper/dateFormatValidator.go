package helper

import "regexp"

func IsValidDateFormat(date string) bool {
	datePattern := `^\d{4}-\d{2}-\d{2}$`
	matched, _ := regexp.MatchString(datePattern, date)
	return matched
}
