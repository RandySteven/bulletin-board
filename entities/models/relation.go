package models

import (
	"task_mission/enums"
	"time"
)

type Relation struct {
	ID             uint64
	UserID         uint64
	FriendID       uint64
	RelationStatus enums.RelationStatus
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
