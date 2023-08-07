package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	ID              int       `gorm:"primarykey" json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	CreatedAt       time.Time `json:"created_at"`
}
