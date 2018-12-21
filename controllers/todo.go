package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/models"
	"todo/utils"

	"github.com/gorilla/mux"
)

// CreateTodo - POST /api/v1/todos
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	todo := &models.Todo{}

	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error parsing request"))
	}
	todo.UserID = userID

	resp := todo.CreateTodo()

	utils.Respond(w, resp)
}

// GetTodos - GET /api/v1/todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)

	data := models.GetTodos(userID)
	resp := utils.Message(true, "Success")
	resp["data"] = data

	utils.Respond(w, resp)
}

// GetTodo - GET /api/v1/todos/{id}
func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, _ := strconv.ParseUint(vars["id"], 10, 64)
	userID := r.Context().Value("user").(uint)

	resp := models.GetTodo(uint(todoID), userID)

	utils.Respond(w, resp)
}

// EditTodo - PATCH /api/v1/todos/{id}
func EditTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, _ := strconv.ParseUint(vars["id"], 10, 64)
	userID := r.Context().Value("user").(uint)
	todo := &models.Todo{}

	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error parsing request"))
	}
	todo.UserID = userID
	todo.ID = uint(todoID)

	resp := todo.EditTodo()

	utils.Respond(w, resp)
}

// DeleteTodo - DELETE /api/v1/todos/{id}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, _ := strconv.ParseUint(vars["id"], 10, 64)
	userID := r.Context().Value("user").(uint)

	resp := models.RemoveTodo(uint(todoID), userID)

	utils.Respond(w, resp)

}
