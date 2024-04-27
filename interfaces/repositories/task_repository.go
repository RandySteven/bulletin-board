package repositories

import "task_mission/entities/models"

type ITaskRepository interface {
	IRepository[models.Task]
}
