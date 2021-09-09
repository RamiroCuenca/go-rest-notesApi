package handler

import (
	"net/http"
)

// Common function wich will be used to send responses for correct http requests
func SendResponse(w http.ResponseWriter, status int, data []byte) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)

}

// Common function wich will be used to send errors for incorrect http requests
func SendError(w http.ResponseWriter, status int) {

	data := []byte(`{}`)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
