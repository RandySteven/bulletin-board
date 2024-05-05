package responses

import (
	"task_mission/entities/models"
	"time"
)

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
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ExpiryDate  string `json:"expiry_date"`
	Reward      struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	} `json:"reward"`
	User struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewTaskListResponse(task *models.Task, user *models.User, reward *models.Reward) *TaskListResponse {
	return &TaskListResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		ExpiryDate:  task.ExpiredDate.Format(time.RFC3339),
		Reward: struct {
			ID   uint64 `json:"id"`
			Name string `json:"name"`
		}{ID: reward.ID, Name: reward.Name},
		User: struct {
			ID   uint64 `json:"id"`
			Name string `json:"name"`
		}{ID: user.ID, Name: user.UserName},
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		DeletedAt: task.DeletedAt,
	}
}
