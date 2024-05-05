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

func (r *rewardCategoryRepository) FindByRewardID(ctx context.Context, rewardID uint64) ([]*models.RewardCategory, error) {
	return r.findByConditionId(ctx, queries.SelectByRewardID, rewardID)
}

func (r *rewardCategoryRepository) FindByCategoryID(ctx context.Context, categoryID uint64) ([]*models.RewardCategory, error) {
	return r.findByConditionId(ctx, queries.SelectByCategoryID, categoryID)
}

func (r *rewardCategoryRepository) Save(ctx context.Context, request *models.RewardCategory) (result *uint64, err error) {
	return utils.Save[models.RewardCategory](ctx, r.db, queries.InsertIntoRewardCategory, &request.RewardID, &request.CategoryID)
}

func (r *rewardCategoryRepository) FindAll(ctx context.Context) (result []*models.RewardCategory, err error) {
	rewardCategory := &models.RewardCategory{}
	result, err = utils.FindAll[models.RewardCategory](ctx, r.db, queries.SelectFromRewardCategory, rewardCategory)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *rewardCategoryRepository) Find(ctx context.Context, id uint64) (result *models.RewardCategory, err error) {
	return
}

func (r *rewardCategoryRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r *rewardCategoryRepository) Update(ctx context.Context, request *models.RewardCategory) (result *models.RewardCategory, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *rewardCategoryRepository) findByConditionId(ctx context.Context, query queries.GoQuery, conditionId uint64) ([]*models.RewardCategory, error) {
	rows, err := r.db.QueryContext(ctx, query.ToString(), conditionId)
	if err != nil {
		return nil, err
	}
	rewardCategoryList := make([]*models.RewardCategory, 0)
	for rows.Next() {
		rewardCategory := &models.RewardCategory{}
		err = rows.Scan(&rewardCategory.ID, &rewardCategory.RewardID, &rewardCategory.CategoryID, &rewardCategory.CreatedAt, &rewardCategory.UpdatedAt, &rewardCategory.DeletedAt)
		if err != nil {
			return nil, err
		}
		rewardCategoryList = append(rewardCategoryList, rewardCategory)
	}
	return rewardCategoryList, nil
}

func NewRewardCategoryRepository(db *sql.DB) *rewardCategoryRepository {
	return &rewardCategoryRepository{db: db}
}

var _ repositories.IRewardCategoryRepository = &rewardCategoryRepository{}
