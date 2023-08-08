package internal

import (
	"encoding/json"
	"net/http"
)

// WriteJsonResponse => set header content-type to json and return json response
func WriteJsonResponse(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application-json")

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}
