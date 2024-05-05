package repositories

import (
	"context"
	"task_mission/entities/models"
)

type IRewardCategoryRepository interface {
	IRepository[models.RewardCategory]
	FindByRewardID(ctx context.Context, rewardID uint64) ([]*models.RewardCategory, error)
	FindByCategoryID(ctx context.Context, categoryID uint64) ([]*models.RewardCategory, error)
}
