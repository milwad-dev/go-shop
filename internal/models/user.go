package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	ID              int            `gorm:"primaryKey" json:"id"`
	Name            string         `json:"name"`
	Email           string         `json:"email"`
	Password        string         `json:"password"`
	EmailVerifiedAt time.Time      `json:"email_verified_at"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
