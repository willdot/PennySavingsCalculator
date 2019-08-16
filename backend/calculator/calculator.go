package calculator

import (
	"errors"
	"time"
)

// ErrStartDateAfterEndDate is for when user tries to use a start date that is after the end date
var ErrStartDateAfterEndDate = errors.New("Start date can't be after end date")

// ErrDatesHaveDifferentYears is for when the user tries to provide 2 dates from different years
var ErrDatesHaveDifferentYears = errors.New("The start and end date must be in the same year")

// CalculateHowMuchToSaveBetweenDays takes a start date and an end date and returns how much to save for this period
func CalculateHowMuchToSaveBetweenDays(start, end time.Time) (int, error) {

	if end.Before(start) {
		return 0, ErrStartDateAfterEndDate
	}

	if err := makeSureBothDatesAreInSameYear(start, end); err != nil {
		return 0, err
	}

	yearStart := calculateStartDateOfYear(start)

	startDayFromFirstOfYear := calculateDaysBetween(yearStart, start)

	endDayFromFirstOfYear := calculateDaysBetween(yearStart, end)

	totalToSave := calculateCostOfDays(startDayFromFirstOfYear, endDayFromFirstOfYear)

	return totalToSave, nil
}

func makeSureBothDatesAreInSameYear(start, end time.Time) error {

	startYear, _, _ := start.Date()
	endYear, _, _ := end.Date()

	if startYear != endYear {
		return ErrDatesHaveDifferentYears
	}

	return nil
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
