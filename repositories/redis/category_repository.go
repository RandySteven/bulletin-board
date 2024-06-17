package redis_repositories

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"time"
)

type categoryRepository struct {
	client *redis.Client
}

func (c *categoryRepository) Save(ctx context.Context, request *models.Category) (result *uint64, err error) {
	key := fmt.Sprintf("category:%d", request.ID)
	var i time.Duration
	err = c.client.Set(ctx, key, request, i).Err()
	if err != nil {
		return nil, err
	}
	return &request.ID, nil
}

func (c *categoryRepository) FindAll(ctx context.Context) (result []*models.Category, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *categoryRepository) Find(ctx context.Context, id uint64) (result *models.Category, err error) {
	err = c.client.Get(ctx, fmt.Sprintf("category:%d", id)).Scan(&result)
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

func NewCategoryRepository(client *redis.Client) *categoryRepository {
	return &categoryRepository{
		client: client,
	}
}

var _ repositories.ICategoryRepository = &categoryRepository{}
