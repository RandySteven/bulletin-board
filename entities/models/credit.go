package models

import "time"

type Credit struct {
	ID          uint64
	FromUserID  uint64 //yg kasik credit
	ToUserID    uint64 //yg nerima credit
	Credit      float64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
