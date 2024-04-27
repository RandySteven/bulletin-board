package repositories

import "task_mission/entities/models"

type IRoleRepository interface {
	IRepository[models.Role]
}
