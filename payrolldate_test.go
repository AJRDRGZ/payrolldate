package payrolldate

import (
	"reflect"
	"testing"
	"time"
)

func TestStartDateOfMonth(t *testing.T) {
	tests := []struct {
		args time.Time
		want time.Time
	}{
		{Date("2020-01-04"), Date("2020-01-01")},
		{Date("2020-02-09"), Date("2020-02-01")},
		{Date("2020-02-29"), Date("2020-02-01")},
		{Date("2020-03-14"), Date("2020-03-01")},
		{Date("2020-04-07"), Date("2020-04-01")},
		{Date("2020-05-25"), Date("2020-05-01")},
		{Date("2020-06-10"), Date("2020-06-01")},
		{Date("2020-07-23"), Date("2020-07-01")},
		{Date("2020-07-01"), Date("2020-07-01")},
		{Date("2020-08-10"), Date("2020-08-01")},
		{Date("2020-09-11"), Date("2020-09-01")},
		{Date("2020-10-09"), Date("2020-10-01")},
		{Date("2020-11-28"), Date("2020-11-01")},
		{Date("2020-12-30"), Date("2020-12-01")},
	}
	for _, tt := range tests {
		if got := StartDateOfMonth(tt.args); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("StartDateOfMonth(%v) = %v, want %v", tt.args, got, tt.want)
		}
	}
}

func TestEndDateOfMonth(t *testing.T) {
	tests := []struct {
		args time.Time
		want time.Time
	}{
		{Date("2020-01-04"), Date("2020-01-31")},
		{Date("2020-02-09"), Date("2020-02-29")},
		{Date("2019-02-10"), Date("2019-02-28")},
		{Date("2020-03-14"), Date("2020-03-31")},
		{Date("2020-04-07"), Date("2020-04-30")},
		{Date("2020-05-25"), Date("2020-05-31")},
		{Date("2020-06-10"), Date("2020-06-30")},
		{Date("2020-07-23"), Date("2020-07-31")},
		{Date("2020-07-01"), Date("2020-07-31")},
		{Date("2020-08-10"), Date("2020-08-31")},
		{Date("2020-09-11"), Date("2020-09-30")},
		{Date("2020-10-09"), Date("2020-10-31")},
		{Date("2020-11-28"), Date("2020-11-30")},
		{Date("2020-12-30"), Date("2020-12-31")},
	}
	for _, tt := range tests {
		got := EndDateOfMonth(tt.args)
		if got != tt.want {
			t.Errorf("EndDateOfMonth(%v) = %v, want %v", tt.args, got, tt.want)
		}
	}
}

func TestIsLastDayOfFebruary(t *testing.T) {
	tests := []struct {
		args time.Time
		want bool
	}{
		{Date("2020-02-04"), false},
		{Date("2020-02-29"), true},
		{Date("2020-02-10"), false},
		{Date("2019-02-28"), true},
		{Date("2020-02-28"), false},
		{Date("2020-05-25"), false},
		{Date("2020-06-10"), false},
		{Date("2020-07-23"), false},
		{Date("2020-08-10"), false},
		{Date("2020-12-30"), false},
	}
	for _, tt := range tests {
		if got := IsLastDayOfFebruary(tt.args); got != tt.want {
			t.Errorf("IsLastDayOfFebruary(%v) = %v, want %v", tt.args, got, tt.want)
		}
	}
}

func TestDays360(t *testing.T) {
	tests := []struct {
		startDate time.Time
		endDate   time.Time
		want      int
	}{
		{Date("2020-01-31"), Date("2020-02-15"), 16},
		{Date("2020-02-05"), Date("2020-03-03"), 29},
		{Date("2020-02-29"), Date("2020-03-03"), 5},
		{Date("2020-02-01"), Date("2020-02-29"), 30},
		{Date("2020-12-01"), Date("2020-12-31"), 30},
		{Date("2020-01-01"), Date("2020-06-30"), 180},
		{Date("2020-07-01"), Date("2020-12-31"), 180},
		{Date("2019-02-25"), Date("2019-03-05"), 11},
		{Date("2020-01-30"), Date("2020-02-05"), 6},
		{Date("2020-01-31"), Date("2020-02-05"), 6},
		{Date("2020-02-29"), Date("2020-02-29"), 2},
		{Date("2020-01-01"), Date("2020-12-31"), 360},
		{Date("2019-02-28"), Date("2019-02-28"), 3},
		{Date("2020-12-01"), Date("2021-11-30"), 360},
		{Date("2023-03-01"), Date("2024-02-29"), 360},
		{Date("2022-03-01"), Date("2023-02-28"), 360},
		{Date("2006-02-28"), Date("2007-02-27"), 360},
		{Date("2024-01-01"), Date("2024-12-31"), 360},
		{Date("2024-01-01"), Date("2024-12-30"), 360},
		{Date("2020-11-01"), Date("2021-10-31"), 360},
		{Date("2020-11-01"), Date("2021-10-30"), 360},
		{Date("2019-06-01"), Date("2020-05-31"), 360},
		{Date("2024-03-01"), Date("2025-02-28"), 360},
		{Date("2020-09-19"), Date("2021-09-18"), 360},
		{Date("2019-01-26"), Date("2020-01-25"), 360},
		{Date("2019-12-25"), Date("2020-12-24"), 360},
		{Date("2021-07-14"), Date("2022-07-13"), 360},
		{Date("2021-04-08"), Date("2022-04-07"), 360},
		{Date("2023-02-20"), Date("2024-02-19"), 360},
		{Date("2023-02-21"), Date("2024-02-20"), 360},
		{Date("2023-02-22"), Date("2024-02-21"), 360},
		{Date("2023-02-23"), Date("2024-02-22"), 360},
		{Date("2023-02-24"), Date("2024-02-23"), 360},
		{Date("2023-02-25"), Date("2024-02-24"), 360},
		{Date("2023-02-26"), Date("2024-02-25"), 360},
		{Date("2023-02-27"), Date("2024-02-26"), 360},
		{Date("2023-02-28"), Date("2024-02-27"), 360},
		{Date("2023-03-01"), Date("2024-02-28"), 358},
		{Date("2023-03-01"), Date("2024-02-29"), 360},
		{Date("2023-03-02"), Date("2024-03-01"), 360},
		{Date("2023-03-03"), Date("2024-03-02"), 360},
		{Date("2023-02-28"), Date("2024-02-27"), 360},
	}
	for _, tt := range tests {
		if got := Days360(tt.startDate, tt.endDate); got != tt.want {
			t.Errorf("Days360(%v, %v) = %v, want %v", tt.startDate, tt.endDate, got, tt.want)
		}
	}
}

func TestPreviousDate360(t *testing.T) {
	tests := []struct {
		date time.Time
		want time.Time
	}{
		{Date("2021-11-30"), Date("2020-12-01")},
		{Date("2024-02-29"), Date("2023-03-01")},
		// fix initial date
		{Date("2024-02-28"), Date("2023-03-01")},
		{Date("2028-02-28"), Date("2027-03-01")},
		{Date("2024-12-31"), Date("2024-01-01")}, // complete year
		{Date("2024-12-30"), Date("2024-01-01")}, // complete year
		{Date("2021-10-31"), Date("2020-11-01")},
		{Date("2021-10-30"), Date("2020-11-01")},
		{Date("2020-05-31"), Date("2019-06-01")}, // Año bisiesto
		{Date("2021-09-18"), Date("2020-09-19")},
		{Date("2020-01-25"), Date("2019-01-26")},
		{Date("2020-12-24"), Date("2019-12-25")},
		{Date("2022-07-13"), Date("2021-07-14")},
		{Date("2022-04-07"), Date("2021-04-08")},
	}
	for _, tt := range tests {
		if got := PreviousDate360(tt.date); got != tt.want {
			t.Errorf("PreviousDate360(%v) = %v, want %v", tt.date, got, tt.want)
		}
	}
}
