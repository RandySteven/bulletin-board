package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/models"
	"task_mission/handlers"
	"task_mission/mocks/usecasemocks"
	"testing"
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
		FirstName:   "User",
		LastName:    "User 2",
		Email:       "user@email.com",
		Password:    "password",
		UserName:    "user_name123",
		DateOfBirth: "2001-01-01",
	}
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", nil)
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		s.T().Fatal(err) // Handle marshalling error
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	expectedUser := &models.User{ID: 1, Name: "User User 2"}
	s.usecases.On("RegisterUser", mock.Anything, payload).Return(expectedUser, nil)

	handler := handlers.NewUserHandler(&s.usecases)
	handler.RegisterHandler(recorder, req)

	s.Equal(http.StatusCreated, recorder.Code)
}

func (s *UserHandlerSuite) TestUserLoginSuccess() {
	payload := &requests.UserLoginRequest{
		Email:    "user@email.com",
		Password: "password",
	}
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/auth/login", nil)
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		s.T().Fatal(err) // Handle marshalling error
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	s.usecases.On("LoginUser", mock.Anything, payload).Return(mock.AnythingOfType("*responses.UserLoginResponse"), nil)

	handler := handlers.NewUserHandler(&s.usecases)
	handler.LoginHandler(recorder, req)

	s.Equal(http.StatusCreated, recorder.Code)
}

func (s *UserHandlerSuite) TearDownSuite() {}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerSuite))
}
