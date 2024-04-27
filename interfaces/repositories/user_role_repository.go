package repositories

import (
	"context"
	"task_mission/entities/models"
)

type IUserRoleRepository interface {
	IRepository[models.UserRole]
	FindByUserID(ctx context.Context, userId uint64) (result *models.UserRole, err error)
}
