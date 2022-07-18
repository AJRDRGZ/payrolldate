package payrolldate

import (
	"fmt"
	"time"
)

const (
	days365 = time.Hour * 24 * 365
	days364 = time.Hour * 24 * 364
)

// Days360 returns the difference between two days based on the 360 day year.
//
// Algorithm based in https://en.wikipedia.org/wiki/360-day_calendar, this function
// use a similar algorithm to European method (30E/360), however is modified for
// apply the rules for the colombian payroll so:
//
// 	- If either date A or B falls on the 31st of the month, that date will be
// 	  changed to the 30th.
// 	- Where date B falls on the last day of February, that date will be changed
//	  to the 30th.
// 	- All months are considered to last 30 days and hence a full year has 360 days
//  - Always add one day for sum the last day of endDate
func Days360(startDate, endDate time.Time) int {
	startYear, startMonth, startDay := startDate.Date()
	endYear, endMonth, endDay := endDate.Date()

	if startDay == 31 {
		startDay = 30
	}

	if endDay == 31 {
		endDay = 30
	}

	if IsLastDayOfFebruary(endDate) {
		endDay = 30
	}

	return ((endYear - startYear) * 360) + (int(endMonth-startMonth) * 30) + (endDay - startDay) + 1
}

// IsLastDayOfFebruary returns true if the date is the last day of february
func IsLastDayOfFebruary(date time.Time) bool {
	return date.Month() == 2 && date.Day() == EndDateOfMonth(date).Day()
}

// EndDateOfMonth returns the end date of the month
func EndDateOfMonth(date time.Time) time.Time {
	e := StartDateOfMonth(date).AddDate(0, 1, 0).Add(-time.Nanosecond)
	y, m, d := e.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, date.Location())
}

// StartDateOfMonth returns the start date of the month
func StartDateOfMonth(date time.Time) time.Time {
	y, m, _ := date.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, date.Location())
}

// Date returns a time.Time of a string date in the format yyyy-mm-dd
func Date(date string) time.Time {
	layoutISO := "2006-01-02"
	d, err := time.Parse(layoutISO, date)
	if err != nil {
		panic(err)
	}
	return d
}

// PreviousDate360 returns the start date of a date counting previous 360 days
func PreviousDate360(date time.Time) time.Time {
	year, month, day := date.Date()

	if month == 12 && (day == 30 || day == 31) {
		return Date(fmt.Sprintf("%d-01-01", year))
	}

	// Cuando la fecha es 28 de febrero de un año festivo, siempre retornaremos el 01 de marzo del año pasado.
	// Sabemos que la diferencia de días es 358 y no 360, pero es imposible llegar a 360 días.
	// Por esta razón se decidió devolver esta fecha y asumir que este es el cálculo correcto.
	if IsLeapYear(year) && month == 2 && day == 28 {
		return Date(fmt.Sprintf("%d-03-01", year-1))
	}

	if IsLastDayOfFebruary(date) || day == 30 || day == 31 {
		return Date(fmt.Sprintf("%d-%02d-01", year-1, month+1))
	}

	day++
	if month == 2 && !IsLeapYear(year-1) && day > 28 {
		day = 28
	}

	return Date(fmt.Sprintf("%d-%02d-%02d", year-1, month, day))
}

// IsLeapYear returns true if is leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
