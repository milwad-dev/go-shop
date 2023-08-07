package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`

	gorm.Model
}
