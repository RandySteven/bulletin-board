package models

import (
	"task_mission/enums"
	"time"
)

type User struct {
	ID          uint64
	Name        string
	UserName    string
	DateOfBirth time.Time
	Gender      enums.UserGender
	IsVerified  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
