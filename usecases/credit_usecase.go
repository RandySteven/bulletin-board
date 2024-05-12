package usecases

import (
	"context"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
	"task_mission/entities/models"
	"task_mission/enums"
	"task_mission/interfaces/repositories"
	"task_mission/interfaces/usecases"
)

type creditUsecase struct {
	userRepository       repositories.IUserRepository
	creditRepository     repositories.ICreditRepository
	userCreditRepository repositories.IUserCreditRepository
}

func (c *creditUsecase) GiveCredit(ctx context.Context, request *requests.CreditRequest) (response *responses.UserCreditResponse, customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)
	user, err := c.userCreditRepository.Find(ctx, userId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
	}
	credit := &models.Credit{
		FromUserID:  user.ID,
		ToUserID:    request.ToUserID,
		Credit:      request.Credit,
		Description: request.Description,
	}
	cId, err := c.creditRepository.Save(ctx, credit)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to insert credit`, err)
	}
	credit, err = c.creditRepository.Find(ctx, *cId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to find credit`, err)
	}
	response = &responses.UserCreditResponse{
		Credit: credit.Credit,
		UserID: user.ID,
	}
	return response, nil
}

func (c *creditUsecase) SeeUserCredit(ctx context.Context, userId uint64) (response *responses.UserCreditResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (c *creditUsecase) SeeAllCredits(ctx context.Context) (result []*responses.UserCreditResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

var _ usecases.ICreditUseCase = &creditUsecase{}

func NewCreditUsecase(
	userRepository repositories.IUserRepository,
	creditRepository repositories.ICreditRepository,
	userCreditRepository repositories.IUserCreditRepository) *creditUsecase {
	return &creditUsecase{
		userRepository:       userRepository,
		creditRepository:     creditRepository,
		userCreditRepository: userCreditRepository,
	}
}
