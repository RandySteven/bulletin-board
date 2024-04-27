package repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type rewardCategoryRepository struct {
	db *sql.DB
}

func (r rewardCategoryRepository) Save(ctx context.Context, request *models.RewardCategory) (result *uint64, err error) {
	return utils.Save[models.RewardCategory](ctx, r.db, queries.InsertIntoRewardCategory, &request.RewardID, &request.CategoryID)
}

func (r rewardCategoryRepository) FindAll(ctx context.Context) (result []*models.RewardCategory, err error) {
	//TODO implement me
	panic("implement me")
}

func (r rewardCategoryRepository) Find(ctx context.Context, id uint64) (result *models.RewardCategory, err error) {
	//TODO implement me
	panic("implement me")
}

func (r rewardCategoryRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r rewardCategoryRepository) Update(ctx context.Context, request *models.RewardCategory) (result *models.RewardCategory, err error) {
	//TODO implement me
	panic("implement me")
}

func NewRewardCategoryRepository(db *sql.DB) *rewardCategoryRepository {
	return &rewardCategoryRepository{db: db}
}

var _ repositories.IRewardCategoryRepository = &rewardCategoryRepository{}
