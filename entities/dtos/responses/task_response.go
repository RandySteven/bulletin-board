package responses

import (
	"task_mission/entities/models"
	"task_mission/enums"
	"time"
)

type (
	TaskResponse struct {
		ID        uint64     `json:"id"`
		TaskID    uint64     `json:"task_id"`
		UserID    uint64     `json:"user_id"`
		RewardID  uint64     `json:"reward_id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	TaskListResponse struct {
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
		Status    enums.TaskStatus `json:"status"`
		CreatedAt time.Time        `json:"created_at"`
		UpdatedAt time.Time        `json:"updated_at"`
		DeletedAt *time.Time       `json:"deleted_at"`
	}

	TaskDetailResponse struct {
		ID          uint64 `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		ExpiryDate  string `json:"expiry_date"`
		Reward      struct {
			ID          uint64 `json:"id"`
			Name        string `json:"name"`
			Image       string `json:"image"`
			Description string `json:"description"`
		} `json:"reward"`
		User struct {
			ID       uint64 `json:"id"`
			Name     string `json:"name"`
			UserName string `json:"user_name"`
			Image    string `json:"image"`
		} `json:"user"`
		Status    enums.TaskStatus `json:"status"`
		CreatedAt time.Time        `json:"created_at"`
		UpdatedAt time.Time        `json:"updated_at"`
		DeletedAt *time.Time       `json:"deleted_at"`
	}
)

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
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		DeletedAt: task.DeletedAt,
	}
}

func NewTaskDetailResponse(task *models.Task, user *models.User, profile *models.UserProfile, reward *models.Reward) *TaskDetailResponse {
	return &TaskDetailResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		ExpiryDate:  task.ExpiredDate.Format(time.RFC3339),
		Reward: struct {
			ID          uint64 `json:"id"`
			Name        string `json:"name"`
			Image       string `json:"image"`
			Description string `json:"description"`
		}{ID: reward.ID, Name: reward.Name, Image: reward.Image, Description: reward.Description},
		User: struct {
			ID       uint64 `json:"id"`
			Name     string `json:"name"`
			UserName string `json:"user_name"`
			Image    string `json:"image"`
		}{ID: user.ID, Name: user.UserName, UserName: user.UserName, Image: profile.Image},
		Status:    task.Status,
		CreatedAt: task.CreatedAt.Local(),
		UpdatedAt: task.UpdatedAt.Local(),
		DeletedAt: task.DeletedAt,
	}
}
