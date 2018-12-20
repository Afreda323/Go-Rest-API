package main

import (
	"log"
	"net/http"
	"os"
	"todo/middleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Panic("Error", err)
	}
}
