package models

import (
	"task_mission/enums"
	"time"
)

type Task struct {
	ID          uint64
	Title       string
	Description string
	Image       string
	Status      enums.TaskStatus
	UserID      uint64
	ExpiredDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
