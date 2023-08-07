package repositories

import (
	"fmt"
	"github.com/milwad-dev/go-shop/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo => set new repo
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// GetLatestUsers => get the latest users
func (repo *UserRepo) GetLatestUsers() []models.User {
	var users []models.User

	repo.db.Find(&users)

	return users
}

// FindUserById => find user by id
func (repo *UserRepo) FindUserById(id int) models.User {
	var user models.User

	repo.db.First(&user, id)
	fmt.Println(user.Name)
	return user
}
