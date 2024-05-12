package usecases

import (
	"context"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
)

type ICreditUseCase interface {
	GiveCredit(ctx context.Context, request *requests.CreditRequest) (response *responses.UserCreditResponse, customErr *apperror.CustomError)
	SeeUserCredit(ctx context.Context, userId uint64) (response *responses.UserCreditResponse, customErr *apperror.CustomError)
	SeeAllCredits(ctx context.Context) (result []*responses.UserCreditResponse, customErr *apperror.CustomError)
}
