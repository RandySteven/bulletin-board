package models

import "time"

type Reward struct {
	ID          uint64
	Name        string
	Image       string
	Description string
	UserID      uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
