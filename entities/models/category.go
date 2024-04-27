package models

import "time"

type Category struct {
	ID        uint64
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
