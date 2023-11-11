package services

import (
	"time"
)

func ConvertTextExpireDate(year int) string {
	currentTime := time.Now()
	addDate := currentTime.AddDate(year, 0, 0)
	textDate := addDate.Format("2006-01-02")

	return textDate
}

func CheckExpireDate(formattedDate time.Time) (string, error) {
	// Convert timestamp to string in "YYYY-MM-DD" format
	dateString := formattedDate.Format("2006-01-02")

	// Check if the formatted date is before or equal to the current time
	isBeforeOrEqual := formattedDate.Before(time.Now()) || formattedDate.Equal(time.Now())

	if isBeforeOrEqual {
		// Add 3 years to the formatted date
		currentTime := time.Now()
		newDate := currentTime.AddDate(3, 0, 0)
		newDateString := newDate.Format("2006-01-02")
		return newDateString, nil
	}

	return dateString, nil
}
