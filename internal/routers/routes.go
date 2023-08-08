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
	userHandler := handlers.SetUserHandler(db)

	r.HandleFunc("/users", userHandler.UserIndex).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", userHandler.UserShow).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", userHandler.UserDelete).Methods(http.MethodDelete)

	// Return router instance
	return r
}
