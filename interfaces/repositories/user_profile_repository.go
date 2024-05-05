package repositories

import (
	"context"
	"task_mission/entities/models"
)

type IUserProfileRepository interface {
	IRepository[models.UserProfile]
	FindByEmail(ctx context.Context, email string) (result *models.UserProfile, err error)
	FindByUserID(ctx context.Context, userID uint64) (result *models.UserProfile, err error)
}
