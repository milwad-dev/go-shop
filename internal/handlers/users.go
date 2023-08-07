package handlers

import (
	"encoding/json"
	"github.com/milwad-dev/go-shop/internal/repositories"
	"net/http"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	users := repositories.GetLatestUsers()

	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(users)
}
