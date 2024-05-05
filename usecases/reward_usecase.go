package usecases

import (
	"context"
	"log"
	"sync"
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
	userProfileRepo    repositories.IUserProfileRepository
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
	for _, reward := range rewards {
		response := responses.NewRewardListResponse(reward)
		result = append(result, response)
	}
	return result, nil
}

func (r *rewardUsecase) GetRewardByID(ctx context.Context, id uint64) (result *responses.RewardDetailResponse, customErr *apperror.CustomError) {
	reward, err := r.rewardRepo.Find(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, ``, err)
	}

	var (
		wg              sync.WaitGroup
		errCh           = make(chan *apperror.CustomError)
		categoryRewards = make([]*models.RewardCategory, 0)
		category        = &models.Category{}
		categories      = make([]*models.Category, 0)
		task            = &models.Task{}
		user            = &models.User{}
		profile         = &models.UserProfile{}
		taskReward      = &models.TaskReward{}
	)

	wg.Add(3)

	go func() {
		defer wg.Done()
		categoryRewards, err = r.rewardCategoryRepo.FindByRewardID(ctx, reward.ID)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get category reward`, err)
			return
		}

		for _, categoryReward := range categoryRewards {
			log.Println(categoryReward.RewardID, " ", categoryReward.CategoryID)
			category, err = r.categoryRepo.Find(ctx, categoryReward.CategoryID)
			if err != nil {
				errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get category`, err)
				return
			}
			categories = append(categories, category)
		}
	}()

	go func() {
		defer wg.Done()

		taskReward, err = r.taskRewardRepo.FindByRewardId(ctx, id)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get task reward`, err)
			return
		}
		task, err = r.taskRepo.Find(ctx, taskReward.TaskID)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get task`, err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		user, err = r.userRepo.Find(ctx, reward.UserID)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
			return
		}

		profile, err = r.userProfileRepo.FindByUserID(ctx, user.ID)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user profile`, err)
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for errC := range errCh {
		return nil, errC
	}

	result = responses.NewRewardDetailResponse(reward, user, profile, task, categories)

	return result, nil
}

var _ usecases.IRewardUsecase = &rewardUsecase{}

func NewRewardUsecase(
	uow repositories.UnitOfWork,
	rewardRepo repositories.IRewardRepository,
	taskRepo repositories.ITaskRepository,
	userRepo repositories.IUserRepository,
	userProfileRepo repositories.IUserProfileRepository,
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
		userProfileRepo:    userProfileRepo,
		rewardCategoryRepo: rewardCategoryRepo,
		taskRewardRepo:     taskRewardRepo,
	}
}
