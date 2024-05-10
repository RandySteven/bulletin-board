package repositories

import (
	"context"
	"task_mission/entities/models"
)

type IRelationRepository interface {
	IRepository[models.Relation]
	FindUserFollowings(ctx context.Context, userId uint64) (result []*models.Relation, err error)
	FindRelationWithFriend(ctx context.Context, userId uint64, friendId uint64) (result *models.Relation, err error)
	FindUserFollowers(ctx context.Context, userId uint64) (result []*models.Relation, err error)
}
