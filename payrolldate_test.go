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
		if (got.Year() != tt.want.Year()) || (got.Month() != tt.want.Month()) || (got.Day() != tt.want.Day()) {
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
	}
	for _, tt := range tests {
		if got := Days360(tt.startDate, tt.endDate); got != tt.want {
			t.Errorf("Days360(%v, %v) = %v, want %v", tt.startDate, tt.endDate, got, tt.want)
		}
	}
}
