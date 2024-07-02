package repositories

import (
	"context"
	"task_mission/entities/models"
)

type ITaskRepository interface {
	IRepository[models.Task]
	UpdateTasksExpiryDate(ctx context.Context) error
}
