package repositories

import "task_mission/entities/models"

type ICategoryRepository interface {
	IRepository[models.Category]
}
