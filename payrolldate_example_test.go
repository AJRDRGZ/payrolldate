package payrolldate

import "fmt"

func ExampleStartDateOfMonth() {
	date := Date("2020-01-04")
	start := StartDateOfMonth(date)
	fmt.Println(start.Format("2006-01-02"))

	// Output:
	// 2020-01-01
}

func ExampleEndDateOfMonth() {
	date1 := Date("2020-01-05")
	date2 := Date("2020-02-15")
	date3 := Date("2020-10-31")
	end1 := EndDateOfMonth(date1)
	end2 := EndDateOfMonth(date2)
	end3 := EndDateOfMonth(date3)
	fmt.Println(end1.Format("2006-01-02"))
	fmt.Println(end2.Format("2006-01-02"))
	fmt.Println(end3.Format("2006-01-02"))

	// Output:
	// 2020-01-31
	// 2020-02-29
	// 2020-10-31
}

func ExampleIsLastDayOfFebruary() {
	date1 := Date("2020-02-04")
	date2 := Date("2020-07-23")
	date3 := Date("2020-02-29")
	isLastDay1 := IsLastDayOfFebruary(date1)
	isLastDay2 := IsLastDayOfFebruary(date2)
	isLastDay3 := IsLastDayOfFebruary(date3)
	fmt.Println(isLastDay1)
	fmt.Println(isLastDay2)
	fmt.Println(isLastDay3)

	// Output:
	// false
	// false
	// true
}

func ExampleDays360() {
	start := Date("2020-01-31")
	end := Date("2020-02-15")
	days := Days360(start, end)
	fmt.Println(days)

	// Output:
	// 16
}