package handlers

import (
	"github.com/gorilla/mux"
	"github.com/milwad-dev/go-shop/internal"
	"github.com/milwad-dev/go-shop/internal/repositories"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Handler struct {
	repo    *repositories.UserRepo
	service *services.UserService
}

// SetRepositories => Set repositories
func SetRepositories(db *gorm.DB) *Handler {
	repo := repositories.NewUserRepo(db)

	return &Handler{repo: repo}
}

// UserIndex => get the latest users with return json response
func (handler *Handler) UserIndex(w http.ResponseWriter, r *http.Request) {
	users := handler.repo.GetLatestUsers()

	internal.WriteJsonResponse(w, users)
}

// UserShow => find user by id and show the user data
func (handler *Handler) UserShow(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userId, _ := strconv.Atoi(id)

	user := handler.repo.FindUserById(userId)

	internal.WriteJsonResponse(w, user)
}

// UserDelete => find user by id then delete user
func (handler Handler) UserDelete(w http.ResponseWriter, r *http.Request) {

}
