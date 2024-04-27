package models

import "time"

type UserTask struct {
	ID        uint64
	UserID    uint64
	TaskID    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
