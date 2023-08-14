package services

import (
	"github.com/milwad-dev/go-shop/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type UserService struct {
	db *gorm.DB
}

// NewUserService => set new service
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// DeleteUser => delete user by id
func (service *UserService) DeleteUser(id int) bool {
	var user models.User

	service.db.Delete(&user, id)

	return true
}

// StoreUser => store user by data
func (service *UserService) StoreUser(r *http.Request) *models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte(r.Form.Get("password")), bcrypt.DefaultCost)

	user := models.User{
		Name:     r.Form.Get("name"),
		Email:    r.Form.Get("email"),
		Password: string(password),
	}

	if r.Form.Get("email_verified_at") == "1" {
		user.EmailVerifiedAt = time.Now()
	} // TODO:

	service.db.Select("Name", "Email", "Password").Create(&user)

	return &user
}
