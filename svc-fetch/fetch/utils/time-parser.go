package utils

import (
	"strings"
	"time"
)

// parse date string to time format 2006-01-02
func ParseDateString(date string) time.Time {
	if strings.Contains(date, " ") {
		split := strings.Split(date, " ")
		date = split[0]
	} else if strings.Contains(date, "T") {
		split := strings.Split(date, "T")
		date = split[0]
	}

	resDate, _ := time.Parse("2006-01-02", date)
	return resDate
}
