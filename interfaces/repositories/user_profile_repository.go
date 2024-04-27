package repositories

import (
	"context"
	"task_mission/entities/models"
)

type IUserProfileRepository interface {
	IRepository[models.UserProfile]
	FindByEmail(ctx context.Context, email string) (result *models.UserProfile, err error)
}
