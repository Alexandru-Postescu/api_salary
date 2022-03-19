package main

import (
	"log"
	"net/http"
	"os"

	"example.com/api-salary/handler"
	"example.com/api-salary/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER ", log.LstdFlags|log.Lshortfile)

	h := handler.NewHandlers(logger)
	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, ":8000")
	logger.Println("Server starting on 8000...")
	err := srv.ListenAndServe()
	logger.Fatalf("err:%v", err)

}
