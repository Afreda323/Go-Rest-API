package routes

import (
	"todo/controllers"

	"github.com/gorilla/mux"
)

// InitTodoRoutes - Create routes for CRUDing todos
func InitTodoRoutes(r *mux.Router) {
	s := r.PathPrefix("/todos").Subrouter()

	s.HandleFunc("/", controllers.GetTodo).Methods("GET")
	s.HandleFunc("/{id}", controllers.GetTodos).Methods("GET")
	s.HandleFunc("/", controllers.CreateTodo).Methods("POST")
	s.HandleFunc("/{id}", controllers.EditTodo).Methods("PATCH")
	s.HandleFunc("/{id}", controllers.DeleteTodo).Methods("DELETE")
}
