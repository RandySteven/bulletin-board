package repositories

import "task_mission/entities/models"

type IRewardRepository interface {
	IRepository[models.Reward]
}
