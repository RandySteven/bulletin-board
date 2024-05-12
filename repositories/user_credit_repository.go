package repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type userCreditRepository struct {
	db *sql.DB
}

func (u userCreditRepository) Save(ctx context.Context, request *models.UserCredit) (result *uint64, err error) {
	return utils.Save[models.UserCredit](ctx, u.db, queries.InsertUserCredit, &request.UserID)
}

func (u userCreditRepository) FindAll(ctx context.Context) (result []*models.UserCredit, err error) {
	//TODO implement me
	panic("implement me")
}

func (u userCreditRepository) Find(ctx context.Context, id uint64) (result *models.UserCredit, err error) {
	return
}

func (u userCreditRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (u userCreditRepository) Update(ctx context.Context, request *models.UserCredit) (result *models.UserCredit, err error) {
	//TODO implement me
	panic("implement me")
}

var _ repositories.IUserCreditRepository = &userCreditRepository{}

func NewUserCreditRepository(db *sql.DB) *userCreditRepository {
	return &userCreditRepository{db: db}
}
