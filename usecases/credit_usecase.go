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
	"task_mission/utils"
)

type creditUsecase struct {
	userRepository       repositories.IUserRepository
	creditRepository     repositories.ICreditRepository
	userCreditRepository repositories.IUserCreditRepository
}

func (c *creditUsecase) GiveCredit(ctx context.Context, request *requests.CreditRequest) (response *responses.UserCreditResponse, customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)
	user, err := c.userRepository.Find(ctx, userId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
	}
	toUser, err := c.userRepository.Find(ctx, request.ToUserID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
	}
	credit := &models.Credit{
		FromID:      user.ID,
		ToID:        toUser.ID,
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
		UserID: toUser.ID,
	}
	return response, nil
}

func (c *creditUsecase) SeeUserCredit(ctx context.Context, userId uint64) (response *responses.UserCreditResponse, customErr *apperror.CustomError) {
	user, err := c.userRepository.Find(ctx, userId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
	}
	credits, err := c.creditRepository.GetUserCredits(ctx, user.ID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user credits`, err)
	}
	creditAvg := utils.CreditsAverage(credits)
	response = &responses.UserCreditResponse{
		UserID: user.ID,
		Credit: creditAvg,
	}
	return response, nil
}

func (c *creditUsecase) SeeAllCredits(ctx context.Context) (result []*responses.UserCreditResponse, customErr *apperror.CustomError) {
	credits, err := c.creditRepository.FindAll(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user credits`, err)
	}
	creditMap := make(map[uint64][]*models.Credit)
	for _, credit := range credits {
		creditMap[credit.ToID] = append(creditMap[credit.ToID], credit)
	}
	for toID, credits := range creditMap {
		creditAvg := utils.CreditsAverage(credits)
		response := &responses.UserCreditResponse{
			UserID: toID,
			Credit: creditAvg,
		}
		result = append(result, response)
	}
	return result, nil
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
