package responses

import "time"

type UserTaskResponse struct {
	ID        uint64     `json:"id"`
	TaskID    uint64     `json:"task_id"`
	UserID    uint64     `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
