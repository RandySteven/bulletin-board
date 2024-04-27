package usecases

import (
	"context"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/models"
)

type IUserCreditUsecase interface {
	SubmitCreditUser(ctx context.Context, request *requests.CreditRequest) (result *models.Credit, err error)
	GetUserCredit(ctx context.Context) (result *models.Credit, err error)
	GetUserCreditDetail(ctx context.Context, id uint64) (result *models.Credit, err error)
}
