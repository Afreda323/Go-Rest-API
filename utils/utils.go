package utils

import (
	"encoding/json"
	"net/http"
)

// Message - generate a response to be sent with json
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
	}
}

// Respond - send a JSON response to user
func Respond(w http.ResponseWriter, res map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
