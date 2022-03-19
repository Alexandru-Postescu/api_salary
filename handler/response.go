package handler

import (
	"encoding/json"
	"net/http"
)

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

func (s *handlers) sendResponse(w http.ResponseWriter, resp interface{}) error {
	response, err := json.Marshal(resp)
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
