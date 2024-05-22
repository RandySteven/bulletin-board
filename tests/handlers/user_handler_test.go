package handlers_test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"task_mission/entities/dtos/requests"
	"task_mission/mocks/usecasemocks"
)

type UserHandlerSuite struct {
	suite.Suite
	usecases usecasemocks.IUserUsecase
}

func (s *UserHandlerSuite) SetupSuite() {
	s.usecases = usecasemocks.IUserUsecase{}
}

func (s *UserHandlerSuite) TestUserRegisterSuccess() {
	payload := &requests.UserRegisterRequest{
		FirstName: "User",
		LastName:  "User 2",
		Email:     "user@email.com",
		Password:  "password",
		UserName:  "user_name123",
	}
	_, _ = s.usecases.RegisterUser(context.Background(), payload)
	s.Assert()
}

func (s *UserHandlerSuite) TearDownSuite() {}
