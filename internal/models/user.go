package models

import "time"

type User struct {
	ID              int
	Name            string
	Email           string
	Password        string
	EmailVerifiedAt time.Time
	CreatedAt       time.Time
}
