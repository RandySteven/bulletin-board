package usecases

import (
	"context"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/models"
)

type IUserUsecase interface {
	RegisterUser(ctx context.Context, register *requests.UserRegisterRequest) (result *models.User, customErr *apperror.CustomError)
	LoginUser(ctx context.Context, login *requests.UserLoginRequest) (result *models.User, customErr *apperror.CustomError)
}
