package main

import (
	. "fmt"
	"time"
)

const DATE_LAYOUT = "2006-01-02"
const MAX_VACATION_DAYS_IN_ONE_PERIOD = 10

func main() {
	experience, usedDays, startDate := requestParameters()

	allowedVacationDays := 15 + experience
	vacationDaysInMonth := float64(allowedVacationDays) / 12
	maxDurationInSpecifiedDate := int(vacationDaysInMonth*float64(startDate.Month())) - usedDays
	if maxDurationInSpecifiedDate > MAX_VACATION_DAYS_IN_ONE_PERIOD {
		maxDurationInSpecifiedDate = MAX_VACATION_DAYS_IN_ONE_PERIOD
	}

	Println("=====================")
	Printf("Experience in company: %d years\nSelected Month: %s\n", experience, startDate.Month())

	Printf("Allowed Vacation in year: %d working days (%d already used)\n", allowedVacationDays, usedDays)

	remain := allowedVacationDays - maxDurationInSpecifiedDate - usedDays
	Printf("Maximum vacation duration in specified date: %d working days (%d will remain)\n", maxDurationInSpecifiedDate, remain)

	Println("Press Enter to close... (Windows crutch)")
	Scanln()
}

func requestParameters() (experience, usedDays int, startDate time.Time) {
	var startDateInput string
	var parseError error

	Print("Company experience (years): ")
	Scanln(&experience)
	for startDateInput == "" || parseError != nil {
		Print("Vacation start date (example: 2020-12-31): ")
		Scanln(&startDateInput)
		startDate, parseError = time.Parse(DATE_LAYOUT, startDateInput)
		if parseError != nil {
			Println("Invalid date format")
		}
	}
	Print("Already used days this year: ")
	Scanln(&usedDays)

	return
}
