package models

import "time"

type (
	Room struct {
		ID        uint64
		UserID1   uint64
		UserID2   uint64
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}

	Chat struct {
		ID        uint64
		RoomID    uint64
		UserID    uint64
		Message   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}
)
