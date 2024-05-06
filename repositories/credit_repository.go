package repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type creditRepository struct {
	db *sql.DB
}

func (c *creditRepository) Save(ctx context.Context, request *models.Credit) (result *uint64, err error) {
	return utils.Save[models.Credit](ctx, c.db, queries.InsertCredit, request.FromUserID, request.ToUserID, request.Credit, request.Description)
}

func (c *creditRepository) FindAll(ctx context.Context) (result []*models.Credit, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *creditRepository) Find(ctx context.Context, id uint64) (result *models.Credit, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *creditRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (c *creditRepository) Update(ctx context.Context, request *models.Credit) (result *models.Credit, err error) {
	//TODO implement me
	panic("implement me")
}

func NewCreditRepository(db *sql.DB) *creditRepository {
	return &creditRepository{
		db: db,
	}
}

var _ repositories.ICreditRepository = &creditRepository{}
