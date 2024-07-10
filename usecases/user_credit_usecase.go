package usecases

import "task_mission/interfaces/repositories"

type userCreditUseCase struct {
	uow            repositories.UnitOfWork
	userRepository repositories.IUserRepository
	userCreditRepo repositories.ICreditRepository
}

func NewUserCreditUseCase(
	uow repositories.UnitOfWork,
	userRepository repositories.IUserRepository,
	userCreditRepo repositories.ICreditRepository) *userCreditUseCase {
	return &userCreditUseCase{
		uow:            uow,
		userRepository: userRepository,
		userCreditRepo: userCreditRepo,
	}
}
