package handler

import (
	"time"
)

var now = func() time.Time { return time.Now() }

type timeHandler struct {
	createdAt time.Time
	payDay    int
}

func NewTimeHandler(payDay int) timeHandler {
	timeHandler := timeHandler{
		now(),
		getPayDay(payDay),
	}
	return timeHandler
}

func (t timeHandler) getDateToday() time.Time {
	return t.createdAt
}

func (t timeHandler) getSalaryDateRaw() time.Time {
	todayDate := t.getDateToday()
	y, m, _ := todayDate.Date()

	var sdRaw time.Time

	if t.payDay < todayDate.Day() {
		sdRaw = time.Date(y, m, t.payDay, 0, 0, 0, 0, time.UTC)
		sdRaw = sdRaw.AddDate(0, 1, 0)
		return sdRaw
	}
	sdRaw = time.Date(y, m, t.payDay, 0, 0, 0, 0, time.UTC)
	return sdRaw
}

func (t timeHandler) getVerifiedDay(day time.Time) time.Time {
	_, _, d := day.Date()

	var shift int
	var verifiedDay time.Time

	switch day.Weekday() {
	case 0:
		shift = 5
		if (d+shift)-t.getDateToday().Day() > 7 {
			shift = -2
		}
	case 6:
		shift = 6
		if (d+shift)-t.getDateToday().Day() > 7 {
			shift = -1
		}
	default:
		shift = 0
	}
	verifiedDay = day.AddDate(0, 0, shift)
	return verifiedDay
}

func (t timeHandler) getSalaryDate() time.Time {
	salaryDate := t.getVerifiedDay(t.getSalaryDateRaw())
	return salaryDate

}

func (t timeHandler) getInterval() int {
	return int(t.getSalaryDate().Sub(t.getDateToday()).Hours() / 24)
}

func (t timeHandler) getResponseDate() (int, string) {
	return t.getInterval(), t.getSalaryDate().String()
}

func (t timeHandler) getResponseDates() []string {
	var dates []string
	salaryDay := t.getSalaryDate()
	y, m, d := salaryDay.Date()

	dates = append(dates, salaryDay.String())
	for i := m + 1; i <= 12; i++ {
		salaryDay = time.Date(y, i, d, 0, 0, 0, 0, time.UTC)
		salaryDay = t.getVerifiedDay(salaryDay)
		dates = append(dates, salaryDay.String())
	}
	return dates
}

// We are sure that the request is valided so getPayDay doesn't need to throw an error
func getPayDay(payDate int) int {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	pd := min(payDate, lastOfMonth.Day())
	return pd
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
