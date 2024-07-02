package repositories__test

import (
	"context"
	"github.com/stretchr/testify/mock"
	"task_mission/entities/models"
	"task_mission/enums"
	"task_mission/mocks/repositorymocks"
	"testing"
)

func TestCreateUser(t *testing.T) {
	repo := &repositorymocks.IUserRepository{}
	ctx := context.Background()
	request := &models.User{
		Name:     "Randy Steven",
		UserName: "randy_steven",
		Gender:   enums.Male,
	}
	repo.On("Save", ctx, request).
		Return(1, mock.AnythingOfType("error")).
		Once()

}
