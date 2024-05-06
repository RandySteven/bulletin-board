package usecases

import (
	"context"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
)

type ICreditUseCase interface {
	GiveCredit(ctx context.Context, request *requests.CreditRequest) (response *responses.UserCreditResponse, err error)
	SeeUserCredit(ctx context.Context, userId uint64) (response *responses.UserCreditResponse, err error)
	SeeAllCredits(ctx context.Context) (result []*responses.UserCreditResponse, err error)
}
