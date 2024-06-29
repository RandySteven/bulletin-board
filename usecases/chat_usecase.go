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

type chatUsecase struct {
	roomRepository repositories.IRoomRepository
	chatRepository repositories.IChatRepository
	userRepository repositories.IUserRepository
}

func (c *chatUsecase) CreateRoom(ctx context.Context, request *requests.RoomRequest) (result *responses.RoomResponse, customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)
	user, err := c.userRepository.Find(ctx, userId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
	}
	room := &models.Room{
		UserID1: user.ID,
		UserID2: request.UserID,
	}
	c.roomRepository.Save(ctx, room)
	return
}

func (c *chatUsecase) GetAllLoginUserRooms(ctx context.Context) (results []*responses.RoomListResponse, customErr *apperror.CustomError) {
	return
}

func (c *chatUsecase) GetChatRoomDetail(ctx context.Context, id uint64) (result *responses.RoomResponse, customErr *apperror.CustomError) {
	return
}

func (c *chatUsecase) SendChat(ctx context.Context, request *requests.ChatRequest) (result *responses.ChatResponse, customErr *apperror.CustomError) {
	return
}

var _ usecases.IChatUseCase = &chatUsecase{}
