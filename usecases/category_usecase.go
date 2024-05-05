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

type categoryUsecase struct {
	categoryRepository       repositories.ICategoryRepository
	rewardRepository         repositories.IRewardRepository
	rewardCategoryRepository repositories.IRewardCategoryRepository
}

func (c *categoryUsecase) CreateCategory(ctx context.Context, request *requests.CategoryRequest) (result *responses.CategoryResponse, customErr *apperror.CustomError) {
	category := &models.Category{
		Category: request.Category,
	}
	id, err := c.categoryRepository.Save(ctx, category)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create category`, err)
	}
	category, err = c.categoryRepository.Find(ctx, *id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to find category`, err)
	}
	result = responses.NewCategoryListResponse(category)
	return result, nil
}

func (c *categoryUsecase) GetAllCategories(ctx context.Context) (result []*responses.CategoryResponse, customErr *apperror.CustomError) {
	categories, err := c.categoryRepository.FindAll(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to find categories`, err)
	}
	result = responses.NewCategoryListResponses(categories)
	return result, nil
}

func (c *categoryUsecase) GetCategoryById(ctx context.Context, categoryId uint64) (result *responses.CategoryDetailResponse, customErr *apperror.CustomError) {
	category, err := c.categoryRepository.Find(ctx, categoryId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to find category`, err)
	}

	categoryRewards, err := c.rewardCategoryRepository.FindByCategoryID(ctx, category.ID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to find reward category`, err)
	}

	var rewards []*models.Reward
	for _, categoryReward := range categoryRewards {
		reward, err := c.rewardRepository.Find(ctx, categoryReward.RewardID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to find category`, err)
		}
		rewards = append(rewards, reward)
	}

	result = responses.NewCategoryDetailResponse(category, rewards)
	return result, nil
}

var _ usecases.ICategoryUsecase = &categoryUsecase{}

func NewCategoryUsecase(
	categoryRepository repositories.ICategoryRepository,
	rewardRepository repositories.IRewardRepository,
	rewardCategoryRepository repositories.IRewardCategoryRepository) *categoryUsecase {
	return &categoryUsecase{
		categoryRepository:       categoryRepository,
		rewardRepository:         rewardRepository,
		rewardCategoryRepository: rewardCategoryRepository,
	}
}
