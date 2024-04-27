package repositories

import (
	"context"
	"task_mission/entities/models"
)

type ITaskRewardRepository interface {
	IRepository[models.TaskReward]
	FindByTaskId(ctx context.Context, taskId uint64) (result *models.TaskReward, err error)
}
