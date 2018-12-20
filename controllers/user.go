package controllers

import (
	"encoding/json"
	"net/http"
	"todo/models"
	"todo/utils"
)

// SignUp - POST /api/v1/users/signup
func SignUp(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user) // pipe body into user model
	resp := user.CreateUser()            // attempt to create user
	utils.Respond(w, resp)               // send res
}

// LogIn - POST /api/v1/users/login
func LogIn(w http.ResponseWriter, r *http.Request) {

}
