package repositories

import "task_mission/entities/models"

type IRewardCategoryRepository interface {
	IRepository[models.RewardCategory]
}
