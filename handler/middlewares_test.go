package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handlers_getGetRequestOnly(t *testing.T) {
	var n int
	h := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		n++
	})
	handler := NewHandlers(log.Default())
	h = handler.getGetRequestOnly(h)

	// test: not get request
	r := httptest.NewRequest("POST", "/v1/list-how-many?salary_day=15", nil)
	h.ServeHTTP(httptest.NewRecorder(), r)
	isEqual(t, n, 0)

	// test: valid get request
	r = httptest.NewRequest("GET", "/v1/list-how-many?salary_day=15", nil)
	h.ServeHTTP(httptest.NewRecorder(), r)
	isEqual(t, n, 1)
}

func Test_handlers_getValidQueryOnly(t *testing.T) {
	var n int
	h := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		n++
	})
	handler := NewHandlers(log.Default())
	h = handler.getValidQueryOnly(h)

	// test: not valid query
	r := httptest.NewRequest("GET", "/v1/list-how-many?salary_day=32", nil)
	h.ServeHTTP(httptest.NewRecorder(), r)
	isEqual(t, n, 0)

	// test: valid query
	r = httptest.NewRequest("POST", "/v1/list-how-many?salary_day=15", nil)
	h.ServeHTTP(httptest.NewRecorder(), r)
	isEqual(t, n, 1)
}

func isEqual(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got:%v, want:%v", got, want)
	}

}
