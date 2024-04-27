package models

import "time"

type UserCredit struct {
	ID        uint64
	Credit    float64
	UserID    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
