package handler

import (
	"log"
	"net/http"
	"time"
)

type handlers struct {
	logger *log.Logger
}

func NewHandlers(logger *log.Logger) *handlers {
	return &handlers{
		logger,
	}
}

// loggerHandler is logging the duration of a incoming request to the router
func (h *handlers) loggerHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(rw, r)
		h.logger.Printf("request processed in %s\n", time.Since(startTime))
	}
}

// SetupRoutes is setting up the router by mapping the endpoints to their corresponding handler
func (h *handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/v1/list-how-much", h.loggerHandler(h.getGetRequestOnly(h.getValidQueryOnly(h.handleSalaryDay))))
	mux.HandleFunc("/v1/list-how-many", h.loggerHandler(h.getGetRequestOnly(h.getValidQueryOnly(h.handleMonths))))
}
