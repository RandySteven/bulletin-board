package usecases

import (
	"context"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
)

type IChatUseCase interface {
	CreateRoom(ctx context.Context, request *requests.RoomRequest) (result *responses.RoomResponse, customErr *apperror.CustomError)
	GetAllLoginUserRooms(ctx context.Context) (results []*responses.RoomListResponse, customErr *apperror.CustomError)
	GetChatRoomDetail(ctx context.Context, id uint64) (result *responses.RoomResponse, customErr *apperror.CustomError)
	SendChat(ctx context.Context, request *requests.ChatRequest) (result *responses.ChatResponse, customErr *apperror.CustomError)
}
