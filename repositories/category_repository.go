package repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type categoryRepository struct {
	db *sql.DB
}

func (c *categoryRepository) Save(ctx context.Context, request *models.Category) (result *uint64, err error) {
	return utils.Save[models.Category](ctx, c.db, queries.InsertIntoCategory, request.Category)
}

func (c *categoryRepository) FindAll(ctx context.Context) (result []*models.Category, err error) {
	category := &models.Category{}
	result, err = utils.FindAll[models.Category](ctx, c.db, queries.SelectAllCategories, category)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *categoryRepository) Find(ctx context.Context, id uint64) (result *models.Category, err error) {
	result = &models.Category{}
	err = utils.FindByID[models.Category](ctx, c.db, queries.SelectCategoryByID, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *categoryRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryRepository) Update(ctx context.Context, request *models.Category) (result *models.Category, err error) {
	//TODO implement me
	panic("implement me")
}

func NewCategoryRepository(db *sql.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

var _ repositories.ICategoryRepository = &categoryRepository{}
