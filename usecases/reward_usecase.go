package usecases

import (
	"context"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/interfaces/usecases"
)

type rewardUsecase struct {
	uow                repositories.UnitOfWork
	rewardRepo         repositories.IRewardRepository
	taskRepo           repositories.ITaskRepository
	userRepo           repositories.IUserRepository
	categoryRepo       repositories.ICategoryRepository
	rewardCategoryRepo repositories.IRewardCategoryRepository
	taskRewardRepo     repositories.ITaskRewardRepository
}

func (r *rewardUsecase) CreateReward(ctx context.Context, request *requests.CreateRewardRequest) (result *models.Reward, customerr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (r *rewardUsecase) GetAllRewards(ctx context.Context) (result []*responses.RewardListResponse, customErr *apperror.CustomError) {
	rewards, err := r.rewardRepo.FindAll(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, ``, err)
	}
	result = responses.NewRewardListResponses(rewards)
	return result, nil
}

func (r *rewardUsecase) GetRewardByID(ctx context.Context, id uint64) (result *models.Reward, customErr *apperror.CustomError) {
	result, err := r.rewardRepo.Find(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, ``, err)
	}
	return result, nil
}

var _ usecases.IRewardUsecase = &rewardUsecase{}

func NewRewardUsecase(
	uow repositories.UnitOfWork,
	rewardRepo repositories.IRewardRepository,
	taskRepo repositories.ITaskRepository,
	userRepo repositories.IUserRepository,
	categoryRepo repositories.ICategoryRepository,
	rewardCategoryRepo repositories.IRewardCategoryRepository,
	taskRewardRepo repositories.ITaskRewardRepository,
) *rewardUsecase {
	return &rewardUsecase{
		uow:                uow,
		rewardRepo:         rewardRepo,
		taskRepo:           taskRepo,
		userRepo:           userRepo,
		categoryRepo:       categoryRepo,
		rewardCategoryRepo: rewardCategoryRepo,
		taskRewardRepo:     taskRewardRepo,
	}
}
