package repositories

import "task_mission/entities/models"

type IUserRepository interface {
	IRepository[models.User]
}
