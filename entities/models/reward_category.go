package models

import "time"

type RewardCategory struct {
	ID         uint64
	RewardID   uint64
	CategoryID uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
