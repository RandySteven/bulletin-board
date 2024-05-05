package responses

import (
	"task_mission/entities/models"
	"time"
)

type UserTaskResponse struct {
	ID        uint64     `json:"id"`
	TaskID    uint64     `json:"task_id"`
	UserID    uint64     `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewUserTaskResponse(userTask *models.UserTask) *UserTaskResponse {
	return &UserTaskResponse{
		ID:        userTask.ID,
		TaskID:    userTask.TaskID,
		UserID:    userTask.UserID,
		CreatedAt: userTask.CreatedAt.Local(),
		UpdatedAt: userTask.UpdatedAt.Local(),
		DeletedAt: userTask.DeletedAt,
	}
}
