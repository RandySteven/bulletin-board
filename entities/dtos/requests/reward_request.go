package requests

type CreateRewardRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	TaskID      uint64 `json:"task_id"`
}
