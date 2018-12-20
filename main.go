package main

import (
	"log"
	"net/http"
	"os"
	"todo/middleware"
	"todo/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	v1 := r.PathPrefix("/api/v1").Subrouter()
	routes.InitUserRoutes(v1)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Panic("Error", err)
	}
}
