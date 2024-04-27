package models

import "time"

type UserProfile struct {
	ID        uint64
	Email     string
	Password  string
	Image     string
	UserID    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
