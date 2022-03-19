package handler

import (
	"encoding/json"
	"net/http"
)

// SalaryDay models the response of "/v1/list-how-much" endpoint
type SalaryDay struct {
	NumDays int    `json:"Days-remaining"`
	Date    string `json:"Date"`
}

func (s *handlers) getSalaryDayResponse(th timeHandler) (resp interface{}) {
	numDays, date := th.getResponseDate()
	response := SalaryDay{
		NumDays: numDays,
		Date:    date,
	}
	return response
}

// SalaryDay models the response of "/v1/list-how-many" endpoint
type MonthlyDates struct {
	Dates []string `json:"Dates"`
}

func (s *handlers) getMonthlyDatesResponse(th timeHandler) (resp interface{}) {
	dates := th.getResponseDates()
	response := MonthlyDates{
		Dates: dates,
	}
	return response
}

// sendResponse is serving a get request to the server by writing a JSON struct
// to the ResponseWriter and setting the proper headers
func (s *handlers) sendResponse(w http.ResponseWriter, resp interface{}) error {
	response, err := json.MarshalIndent(resp, "", "\n")
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(response)
	if err != nil {
		return err
	}
	return nil
}
