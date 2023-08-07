package internal

import (
	"encoding/json"
	"net/http"
)

// WriteJsonResponse => set header content-type to json and return json response
func WriteJsonResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(data)
}
