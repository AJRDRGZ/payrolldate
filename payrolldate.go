package payrolldate

import "time"

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
func Days360(start, end time.Time) int {
	startYear, startMonth, startDay := start.Date()
	endYear, endMonth, endDay := end.Date()

	fixLastDay(&startDay)
	fixLastDay(&endDay)
	if IsLastDayOfFebruary(end) {
		endDay = 30
	}

	return diffDaysBetweenYears(endYear, startYear) + diffDaysBetweenMonths(endMonth, startMonth) + diffDays(endDay, startDay)
}

// fixLastDay returns 30 if day is 31
func fixLastDay(day *int) {
	if *day == 31 {
		*day = 30
	}
}

// diffDaysBetweenYears returns 360 days times between two years
func diffDaysBetweenYears(end, start int) int {
	return (end - start) * 360
}

// diffDaysBetweenMonths returns 30 days times betwwen two months
func diffDaysBetweenMonths(end, start time.Month) int {
	return int(end - start) * 30
}

// diffDays returns diff days between
func diffDays(end, start int) int {
	return end - start + 1
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
