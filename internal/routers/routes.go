package routers

import (
	"github.com/gorilla/mux"
	"github.com/milwad-dev/go-shop/internal/handlers"
	"gorm.io/gorm"
	"net/http"
)

func InitRouters(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	// Users routes
	userRepo := handlers.SetRepositories(db)
	r.HandleFunc("/users", userRepo.UserIndex).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", userRepo.UserShow).Methods(http.MethodGet)

	// Return router instance
	return r
}
