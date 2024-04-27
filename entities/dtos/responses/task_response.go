package responses

import "time"

type TaskResponse struct {
	ID        uint64     `json:"id"`
	TaskID    uint64     `json:"task_id"`
	UserID    uint64     `json:"user_id"`
	RewardID  uint64     `json:"reward_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type TaskListResponse struct {
	ID          uint64     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ExpiryDate  string     `json:"expiry_date"`
	RewardID    uint64     `json:"reward_id"`
	Reward      string     `json:"reward"`
	UserID      uint64     `json:"user_id"`
	UserName    string     `json:"user_name"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
