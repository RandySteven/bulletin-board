package repositories

import (
	"context"
	"task_mission/entities/models"
)

type IUserRepository interface {
	IRepository[models.User]
	FindVerifyUser(ctx context.Context, id uint64, isVerified bool) (result *models.User, err error)
}
