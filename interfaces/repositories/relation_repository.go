package repositories

import (
	"context"
	"task_mission/entities/models"
)

type IRelationRepository interface {
	IRepository[models.Relation]
	FindRelationWithUser(ctx context.Context, userId uint) (result []*models.Relation, err error)
	FindRelationWithFriend(ctx context.Context, userId uint, friendId uint) (result *models.Relation, err error)
}
