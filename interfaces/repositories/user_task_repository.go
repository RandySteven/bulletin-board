package repositories

import "task_mission/entities/models"

type IUserTaskRepository interface {
	IRepository[models.UserTask]
}
