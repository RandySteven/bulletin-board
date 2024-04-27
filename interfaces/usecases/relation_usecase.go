package usecases

import (
	"context"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
)

type IRelationUsecase interface {
	AddFriend(ctx context.Context, request *requests.FriendRequest) (result *responses.CreateFriendResponse, err error)
	SeeAllFriends(ctx context.Context)
}
