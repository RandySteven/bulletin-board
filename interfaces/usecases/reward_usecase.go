package usecases

import (
	"context"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
	"task_mission/entities/models"
)

type IRewardUsecase interface {
	CreateReward(ctx context.Context, request *requests.CreateRewardRequest) (result *models.Reward, customerr *apperror.CustomError)
	GetAllRewards(ctx context.Context) (result []*responses.RewardListResponse, customErr *apperror.CustomError)
	GetRewardByID(ctx context.Context, id uint64) (result *models.Reward, customErr *apperror.CustomError)
}
