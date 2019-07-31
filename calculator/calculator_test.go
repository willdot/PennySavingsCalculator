package calculator

import (
	"testing"
	"time"
)

func TestCalculateHowMuchToSaveBetweenDays(t *testing.T) {

	cases := []struct {
		Name          string
		Start         time.Time
		End           time.Time
		Expected      int
		ExpectedError error
	}{
		{
			"1st Jan to 2nd Jan",
			time.Date(2019, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
			time.Date(2019, time.Month(1), 2, 0, 0, 0, 0, time.UTC),
			3,
			nil,
		},
		{
			"1st Feb to 28th feb",
			time.Date(2019, time.Month(2), 1, 0, 0, 0, 0, time.UTC),
			time.Date(2019, time.Month(2), 28, 0, 0, 0, 0, time.UTC),
			1274,
			nil,
		},
		{
			"1st Jan to 31st Dec 2019",
			time.Date(2019, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
			time.Date(2019, time.Month(12), 31, 0, 0, 0, 0, time.UTC),
			66795,
			nil,
		},
		{
			"1st Jan to 31st Dec 2020 (leap year)",
			time.Date(2020, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.Month(12), 31, 0, 0, 0, 0, time.UTC),
			67161,
			nil,
		},
		{
			"Start date after end date",
			time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.Month(12), 31, 0, 0, 0, 0, time.UTC),
			0,
			ErrStartDateAfterEndDate,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got, err := CalculateHowMuchToSaveBetweenDays(test.Start, test.End)

			if err != test.ExpectedError {
				t.Errorf("got error %v, want %v", got, test.ExpectedError)
			}

			if got != test.Expected {
				t.Errorf("got %v, want %v", got, test.Expected)
			}
		})
	}
}

func TestCalculateDaysBetween(t *testing.T) {

	cases := []struct {
		Name     string
		Start    time.Time
		End      time.Time
		Expected int
	}{
		{
			"1st Jan to 2nd Jan",
			time.Date(2019, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
			time.Date(2019, time.Month(1), 2, 0, 0, 0, 0, time.UTC),
			2,
		},
		{
			"1st Feb to 28th feb",
			time.Date(2019, time.Month(2), 1, 0, 0, 0, 0, time.UTC),
			time.Date(2019, time.Month(2), 28, 0, 0, 0, 0, time.UTC),
			28,
		},
		{
			"1st Jan to 31st dec - leap year",
			time.Date(2020, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.Month(12), 31, 0, 0, 0, 0, time.UTC),
			366,
		},
		{
			"1st Jan to 31st dec",
			time.Date(2019, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
			time.Date(2019, time.Month(12), 31, 0, 0, 0, 0, time.UTC),
			365,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := calculateDaysBetween(test.Start, test.End)

			if got != test.Expected {
				t.Errorf("got %v, want %v", got, test.Expected)
			}
		})
	}
}

func TestCalculateCostOfDays(t *testing.T) {

	cases := []struct {
		Name     string
		Start    int
		End      int
		Expected int
	}{
		{
			"1st Jan to 2nd Jan",
			1,
			2,
			3,
		},
		{
			"1st Feb to 28th feb",
			32,
			59,
			1274,
		},
		{
			"Whole year",
			1,
			365,
			66795,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := calculateCostOfDays(test.Start, test.End)

			if got != test.Expected {
				t.Errorf("got %v, want %v", got, test.Expected)
			}
		})
	}
}

func TestCalculateStartDateOfYear(t *testing.T) {
	cases := []struct {
		Name        string
		CurrentYear time.Time
		Expected    time.Time
	}{
		{
			"2019",
			time.Date(2019, time.Month(4), 28, 0, 0, 0, 0, time.UTC),
			time.Date(2019, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"2022",
			time.Date(2022, time.Month(4), 28, 0, 0, 0, 0, time.UTC),
			time.Date(2022, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := calculateStartDateOfYear(test.CurrentYear)

			if got != test.Expected {
				t.Errorf("got %v, want %v", got, test.Expected)
			}
		})
	}
}
