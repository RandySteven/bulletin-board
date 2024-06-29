package usecases

import (
	"context"
	"fmt"
	"log"
	"sync"
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
	errCh := make(chan error, 1)
	resultsCh := make(chan *responses.RoomListResponse, len(result))
	wg := sync.WaitGroup{}

	for _, room := range result {
		wg.Add(1)
		go func(room *models.Room) {
			defer wg.Done()

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

			resultsCh <- &responses.RoomListResponse{
				ID: room.ID,
				User1: &responses.UserRoomChatResponse{
					ID:       user1.ID,
					UserName: user1.UserName,
				},
				User2: &responses.UserRoomChatResponse{
					ID:       user2.ID,
					UserName: user2.UserName,
				},
			}
		}(room)
	}
	go func() {
		wg.Wait()
		close(resultsCh)
		close(errCh)
	}()

	for res := range resultsCh {
		results = append(results, res)
	}

	select {
	case err := <-errCh:
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to get all user rooms", err)
		}
	default:
	}

	log.Println(results)
	return results, nil
}

func (c *chatUsecase) GetChatRoomDetail(ctx context.Context, id uint64) (result *responses.RoomResponse, customErr *apperror.CustomError) {
	room, err := c.roomRepository.Find(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get room`, err)
	}
	wg := new(sync.WaitGroup)
	user1 := &models.User{ID: room.UserID1}
	user2 := &models.User{ID: room.UserID2}
	errCh := make(chan error, 1)
	wg.Add(2)
	go func() {
		defer wg.Done()
		user1, err = c.userRepository.Find(ctx, user1.ID)
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		user2, err = c.userRepository.Find(ctx, user2.ID)
		if err != nil {
			errCh <- err
			return
		}
	}()

	wg.Wait()

	for err := range errCh {
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get room`, err)
		}
	}

	result = &responses.RoomResponse{
		ID: room.ID,
		User1: &responses.UserRoomChatResponse{
			ID:       user1.ID,
			UserName: user1.UserName,
		},
		User2: &responses.UserRoomChatResponse{
			ID:       user2.ID,
			UserName: user2.UserName,
		},
		CreatedAt: room.CreatedAt,
	}
	return result, nil
}

func (c *chatUsecase) SendChat(ctx context.Context, request *requests.ChatRequest) (result *responses.ChatResponse, customErr *apperror.CustomError) {
	chatId, err := c.chatRepository.Save(ctx, &models.Chat{
		RoomID:  request.RoomID,
		UserID:  ctx.Value(enums.UserID).(uint64),
		Message: request.Message,
	})
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to save chat`, err)
	}
	chat, err := c.chatRepository.Find(ctx, *chatId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get chat`, err)
	}
	user, err := c.userRepository.Find(ctx, chat.UserID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
	}
	result = &responses.ChatResponse{
		UserName:  user.UserName,
		Message:   chat.Message,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}
	return result, nil
}

var _ usecases.IChatUseCase = &chatUsecase{}

func NewChatUsecase(
	chatRepository repositories.IChatRepository,
	userRepository repositories.IUserRepository,
	roomRepository repositories.IRoomRepository,
) *chatUsecase {
	return &chatUsecase{
		chatRepository: chatRepository,
		userRepository: userRepository,
		roomRepository: roomRepository,
	}
}
