package handlers

import (
	"github.com/gorilla/mux"
	"github.com/milwad-dev/go-shop/internal"
	"github.com/milwad-dev/go-shop/internal/repositories"
	"github.com/milwad-dev/go-shop/internal/services"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Handler struct {
	repo    *repositories.UserRepo
	service *services.UserService
}

// SetUserHandler => set user handlers
func SetUserHandler(db *gorm.DB) *Handler {
	repo := repositories.NewUserRepo(db)
	service := services.NewUserService(db)

	return &Handler{
		repo:    repo,
		service: service,
	}
}

// UserIndex => get the latest users with return json response
func (handler *Handler) UserIndex(w http.ResponseWriter, r *http.Request) {
	users := handler.repo.GetLatestUsers()

	internal.WriteJsonResponse(w, users, 200)
}

// UserStore => get the latest users with return json response
func (handler *Handler) UserStore(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user := handler.service.StoreUser(r)

	internal.WriteJsonResponse(w, user, 201)
}

// UserShow => find user by id and show the user data
func (handler *Handler) UserShow(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userId, _ := strconv.Atoi(id)

	user := handler.repo.FindUserById(userId)

	internal.WriteJsonResponse(w, user, 200)
}

func (handler *Handler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := mux.Vars(r)["id"]
	userId, _ := strconv.Atoi(id)

	user := handler.service.UpdateUser(userId, r)

	internal.WriteJsonResponse(w, user, 201)
}

// UserDelete => find user by id then delete user
func (handler *Handler) UserDelete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userId, _ := strconv.Atoi(id)

	handler.service.DeleteUser(userId)

	data := make(map[string]any)
	data["message"] = "user delete successfully."

	internal.WriteJsonResponse(w, data, 200)
}
