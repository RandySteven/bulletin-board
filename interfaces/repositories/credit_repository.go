package repositories

import "task_mission/entities/models"

type ICreditRepository interface {
	IRepository[models.Credit]
}
