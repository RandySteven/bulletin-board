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

func (c *creditRepository) GetUserCredits(ctx context.Context, userId uint64) (result []*models.Credit, err error) {
	rows, err := c.db.QueryContext(ctx, queries.SelectUserCredits.ToString(), userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		credit := &models.Credit{}
		err = rows.Scan(credit.ID, credit.FromUserID, credit.ToUserID, credit.Credit, credit.Description, credit.CreatedAt, credit.UpdatedAt, credit.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, credit)
	}
	return result, nil
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
