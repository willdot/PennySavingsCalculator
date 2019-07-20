package calculator

import (
	"errors"
	"time"
)

// ErrStartDateAfterEndDate is for when user tries to use a start date that is after the end date
var ErrStartDateAfterEndDate = errors.New("Start date can't be after end date")

// CalculateHowMuchToSaveBetweenDays takes a start date and an end date and returns how much to save for this period
func CalculateHowMuchToSaveBetweenDays(start, end time.Time) (int, error) {

	if end.Before(start) {
		return 0, ErrStartDateAfterEndDate
	}

	yearStart := calculateStartDateOfYear(start)

	startDayFromFirstOfYear := calculateDaysBetween(yearStart, start)

	endDayFromFirstOfYear := calculateDaysBetween(yearStart, end)

	totalToSave := calculateCostOfDays(startDayFromFirstOfYear, endDayFromFirstOfYear)

	return totalToSave, nil
}

func calculateStartDateOfYear(currentDate time.Time) time.Time {

	year, _, _ := currentDate.Date()
	return time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
}

func calculateDaysBetween(start, end time.Time) int {

	days := end.Sub(start).Hours() / 24

	return int(days + 1)
}

func calculateCostOfDays(start, end int) int {

	result := 0

	for i := start; i <= end; i++ {
		result += i
	}

	return result
}
