package usecases

import (
	"context"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
	"task_mission/entities/models"
)

type IUserUsecase interface {
	RegisterUser(ctx context.Context, register *requests.UserRegisterRequest) (result *models.User, customErr *apperror.CustomError)
	LoginUser(ctx context.Context, login *requests.UserLoginRequest) (result *responses.UserLoginResponse, customErr *apperror.CustomError)
	VerifyUser(ctx context.Context, id uint64) (result *models.User, customErr *apperror.CustomError)
	UserDetail(ctx context.Context, id uint64) (result *responses.UserDetailResponse, customErr *apperror.CustomError)
}
