package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/api-salary/server"
)

var handler = NewHandlers(log.Default())
var mux = http.ServeMux{}

func Test_handlers_handleSalaryDay(t *testing.T) {
	handler.SetupRoutes(&mux)
	srv := server.New(&mux, ":8080")
	r := httptest.NewRequest("POST", "/v1/list-how-many?salary_day=15", nil)
	w := httptest.NewRecorder()
	srv.Handler.ServeHTTP(w, r)

	isStatusCodeEqual(t, w.Result().StatusCode, http.StatusMethodNotAllowed)
}

func Test_handlers_handleMonths(t *testing.T) {
	handler.SetupRoutes(&mux)
	srv := server.New(&mux, ":8080")
	r := httptest.NewRequest("GET", "/v1/list-how-many?salary_day=15", nil)
	w := httptest.NewRecorder()
	srv.Handler.ServeHTTP(w, r)

	isStatusCodeEqual(t, w.Result().StatusCode, http.StatusOK)
}

func isStatusCodeEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}
