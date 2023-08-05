package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	ID              int
	Name            string
	Email           string
	Password        string
	EmailVerifiedAt time.Time
	CreatedAt       time.Time
}
