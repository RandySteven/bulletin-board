package repositories

import "task_mission/entities/models"

type IRelationRepository interface {
	IRepository[models.Relation]
}
