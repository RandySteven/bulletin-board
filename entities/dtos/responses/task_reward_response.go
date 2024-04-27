package responses

import "time"

type TaskRewardResponse struct {
	Task struct {
		ID          uint64     `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Image       string     `json:"image"`
		ExpiryDate  string     `json:"expiry_date"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
		DeletedAt   *time.Time `json:"deleted_at"`
	} `json:"task"`
	Reward struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       string `json:"image"`
	} `json:"reward"`
}
