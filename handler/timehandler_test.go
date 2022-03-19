package handler

import (
	"log"
	"testing"
	"time"
)

func Test_timeHandler_getDateToday(t *testing.T) {
	th := NewTimeHandler(10)
	d, m, y := time.Now().Date()
	want := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	got := th.getDateToday()
	if got.Equal(want) {
		t.Fatal("test failed")
	}
}

func Test_timeHandler_getSalaryDateRaw(t *testing.T) {
	th := timeHandler{
		time.Date(2022, 3, 10, 0, 0, 0, 0, time.UTC),
		21,
	}
	sdRaw := th.getSalaryDateRaw()
	log.Print(sdRaw)
	if sdRaw.Day() != 21 {
		t.Fatal("test failed")
	}
}
func Test_timeHandler_getSalaryDate(t *testing.T) {
	th := timeHandler{
		time.Date(2022, 3, 10, 0, 0, 0, 0, time.UTC),
		20,
	}
	sdRaw := th.getSalaryDate()
	log.Print(sdRaw)
	want := 18 // pay day falls on the weekend and the previous friday is on 18
	if sdRaw.Day() != want {
		t.Fatal("test failed")
		log.Println(sdRaw)
	}
}

func Test_timeHandler_getInterval(t *testing.T) {
	th := timeHandler{
		time.Date(2022, 3, 10, 0, 0, 0, 0, time.UTC),
		20,
	}
	got := th.getInterval()
	want := 8 // pay day falls on the weekend and the previous friday is on 18 -> 18 - 10 = 8
	if want != got {
		t.Fatalf("test failed, want: %v, got := %v, ", want, got)
	}

}

func Test_timeHandler_getResponseDate(t *testing.T) {
	tr := NewTimeHandler(26)
	want1 := 8
	want2 := tr.getSalaryDate().String()
	got1, got2 := tr.getResponseDate()

	if want1 != got1 {
		t.Fatalf("interval error. want: %v, got :%v", want1, got1)
	}
	if want2 != got2 {
		t.Fatalf("date error. want:%v, got:%v", want2, got2)

	}

}

func Test_getPayDay(t *testing.T) {
	want := 31
	got := getPayDay(30)
	if got != want {
		t.Logf("got:%v", got)
		t.Error("test failed")
	}
}
