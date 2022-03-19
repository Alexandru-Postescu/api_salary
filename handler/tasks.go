package handler

import (
	"net/http"
	"strconv"
)

// handleSalaryDay handles the request to the "/v1/list-how-much" endpoint
func (h *handlers) handleSalaryDay(rw http.ResponseWriter, r *http.Request) {
	in, err := h.getInput(r)
	if err != nil {
		h.logger.Println("500 Server error")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	th := NewTimeHandler(in)
	output := h.getSalaryDayResponse(th)
	h.logger.Println("handling salary day request")
	err = h.sendResponse(rw, output)
	if err != nil {
		h.logger.Println("500 Server error")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

}

// handleSalaryDay handles the request to the "/v1/list-how-many" endpoint
func (h *handlers) handleMonths(rw http.ResponseWriter, r *http.Request) {
	in, err := h.getInput(r)
	if err != nil {
		h.logger.Println("500 Server error")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	th := NewTimeHandler(in)
	output := h.getMonthlyDatesResponse(th)
	h.logger.Println(" handling monthly dates request")

	err = h.sendResponse(rw, output)
	if err != nil {
		h.logger.Println("500 Server error")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

}

// getInput reads the input from the request
func (h *handlers) getInput(req *http.Request) (int, error) {
	values := req.URL.Query()

	numDay, err := strconv.Atoi(values.Get("salary_day"))
	if err != nil {
		return 0, err
	}
	return numDay, nil
}
