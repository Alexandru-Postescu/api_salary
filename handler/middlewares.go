package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s *handlers) getGetRequestOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if !isRequestValid(r) {
			s.logger.Println("Request not allowed")
			rw.Header().Set("Access-Control-Allow-Methods", "GET")
			rw.WriteHeader(http.StatusMethodNotAllowed)
			var message = []byte(`{"Acces Allowed Only": "GET"}`)
			rw.Write(message)
			http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(rw, r)
	}
}

func (s *handlers) getValidQueryOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if !isQueryValid(r) {
			s.logger.Println("Invalid query")
			var message = []byte(`{"Not found": "Invalid Query/URL"}`)
			rw.Write(message)

			http.NotFound(rw, r)
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
