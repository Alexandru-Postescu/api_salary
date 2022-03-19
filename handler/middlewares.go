package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

// getRequestOnly is validating only GET requests
func (s *handlers) getGetRequestOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if !isRequestValid(r) {
			s.logger.Println("Request not allowed")
			http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
			rw.Header().Set("Access-Control-Allow-Methods", "GET")

			var message = []byte(`{"Acces Allowed Only": "GET"}`)
			rw.Write(message)
			return
		}
		next.ServeHTTP(rw, r)
	}
}

// getValidQueryOnly is validating if an URL is valid. An URL is invalid
// if it misses salary_day value or if that value is invalid
func (s *handlers) getValidQueryOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if !isQueryValid(r) {
			s.logger.Println("Invalid query")
			var message = []byte(`{"Not found": "Invalid Query/URL"}`)
			http.NotFound(rw, r)
			rw.Write(message)

			return
		}
		next.ServeHTTP(rw, r)
	}
}

func isRequestValid(req *http.Request) bool {
	return req.Method == "GET"
}

func isQueryValid(req *http.Request) bool {
	values := req.URL.Query()
	fmt.Printf("%v", values.Get("salary_day"))
	if !values.Has("salary_day") {
		return false
	}
	numDay, err := strconv.Atoi(values.Get("salary_day"))
	if err != nil || numDay < 1 || numDay > 31 {
		return false
	}
	return true
}
