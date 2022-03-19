package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func Test_handlers_getSalaryDayResponse(t *testing.T) {
	th := timeHandler{
		time.Date(2022, 3, 10, 0, 0, 0, 0, time.UTC),
		21,
	}
	want := SalaryDay{
		11,
		"",
	}
	got := handler.getSalaryDayResponse(th)

	if want.NumDays != got.(SalaryDay).NumDays {
		t.Errorf("want: %v, got: %v", want.NumDays, got.(SalaryDay).NumDays)
	}

}

func Test_handlers_getMonthlyDatesResponse(t *testing.T) {
	th := timeHandler{
		time.Date(2022, 3, 10, 0, 0, 0, 0, time.UTC),
		21,
	}

	want := 10 // 10 months left including this one because we start on 10th of march and next salary day is 21 March
	got := handler.getMonthlyDatesResponse(th)

	if want != len(got.(MonthlyDates).Dates) {
		t.Errorf("want: %v, got: %v", want, len(got.(MonthlyDates).Dates))

	}
}

func Test_handlers_sendResponse(t *testing.T) {
	monthsResponse := MonthlyDates{
		Dates: []string{
			"11.10.2021, Luni",
			"11.12.2021, Luni",
			"01.01.2022, Marti",
		},
	}
	rw := httptest.NewRecorder()
	err := handler.sendResponse(rw, monthsResponse)
	if err != nil {
		t.Errorf("err: %v", err)
	}

	resp := rw.Result()
	gotBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	wantBody, _ := json.Marshal(monthsResponse)
	if !reflect.DeepEqual(gotBody, wantBody) {
		t.Errorf("got: %v want :%v", gotBody, wantBody)
	}

	gotStatus := resp.StatusCode
	wantStatus := http.StatusOK
	isEqual(t, gotStatus, wantStatus)

	if ctype := rw.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, "application/json")
	}
}
