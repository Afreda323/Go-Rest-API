package routes

import (
	"todo/controllers"

	"github.com/gorilla/mux"
)

// InitUserRoutes - add routes that belong to user
func InitUserRoutes(r *mux.Router) {
	s := r.PathPrefix("/users").Subrouter()
	s.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	s.HandleFunc("/login", controllers.LogIn).Methods("POST")
}
