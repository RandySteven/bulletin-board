package models

import "time"

type TaskReward struct {
	ID        uint64
	TaskID    uint64
	RewardID  uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}