package usecases

import (
	"context"
	"fmt"
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
	roomExist, err := c.roomRepository.CheckExistsRoom(ctx, room)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get room`, err)
	}
	if roomExist {
		return nil, apperror.NewCustomError(apperror.ErrBadRequest, `room already exists`, fmt.Errorf("room already exists"))
	}
	roomID, err := c.roomRepository.Save(ctx, room)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to save room`, err)
	}
	room, err = c.roomRepository.Find(ctx, *roomID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get room`, err)
	}
	user2, err := c.userRepository.Find(ctx, room.UserID2)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user2`, err)
	}
	result = &responses.RoomResponse{
		ID: room.ID,
		User1: &responses.UserRoomChatResponse{
			ID:       user.ID,
			UserName: user.UserName,
		},
		User2: &responses.UserRoomChatResponse{
			ID:       user2.ID,
			UserName: user2.UserName,
		},
	}
	return result, nil
}

func (c *chatUsecase) GetAllLoginUserRooms(ctx context.Context) (results []*responses.RoomListResponse, customErr *apperror.CustomError) {
	result, err := c.roomRepository.FindAllUserRooms(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get all user rooms`, err)
	}
	errCh := make(chan error, len(result))
	go func() {
		for _, room := range result {
			user1, err := c.userRepository.Find(ctx, room.UserID1)
			if err != nil {
				errCh <- err
				return
			}
			user2, err := c.userRepository.Find(ctx, room.UserID2)
			if err != nil {
				errCh <- err
				return
			}
			results = append(results, &responses.RoomListResponse{
				ID: room.ID,
				User1: &responses.UserRoomChatResponse{
					ID:       user1.ID,
					UserName: user1.UserName,
				},
				User2: &responses.UserRoomChatResponse{
					ID:       user2.ID,
					UserName: user2.UserName,
				},
			})
		}
	}()
	if err := <-errCh; err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get all user rooms`, err)
	}
	return results, nil
}

func (c *chatUsecase) GetChatRoomDetail(ctx context.Context, id uint64) (result *responses.RoomResponse, customErr *apperror.CustomError) {
	return
}

func (c *chatUsecase) SendChat(ctx context.Context, request *requests.ChatRequest) (result *responses.ChatResponse, customErr *apperror.CustomError) {
	return
}

var _ usecases.IChatUseCase = &chatUsecase{}
