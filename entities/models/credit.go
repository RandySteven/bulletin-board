package models

import "time"

type Credit struct {
	ID          uint64
	FromID      uint64 //yg kasik credit
	ToID        uint64 //yg nerima credit
	Credit      float32
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
