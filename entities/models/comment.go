package models

import "time"

type Comment struct {
	ID        uint64
	Comment   string
	ParentID  uint64
	UserID    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
