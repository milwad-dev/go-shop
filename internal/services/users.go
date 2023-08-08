package services

import (
	"github.com/milwad-dev/go-shop/internal/models"
	"gorm.io/gorm"
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
