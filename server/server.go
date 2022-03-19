package server

import (
	"net/http"
)

func New(mux *http.ServeMux, serverAdress string) *http.Server {
	srv := &http.Server{
		Addr:    serverAdress,
		Handler: mux,
	}
	return srv
}
