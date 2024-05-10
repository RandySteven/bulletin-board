package repositories

import (
	"context"
	"task_mission/entities/models"
)

type ICreditRepository interface {
	IRepository[models.Credit]
	GetUserCredits(ctx context.Context, userId uint64) ([]*models.Credit, error)
}
