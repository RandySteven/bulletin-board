package usecases

import (
	"context"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
)

type IRelationUsecase interface {
	AddFriend(ctx context.Context, request *requests.FriendRequest) (result *responses.CreateFriendResponse, customErr *apperror.CustomError)
	SeeAllFriends(ctx context.Context)
	SeeAllFollowers(ctx context.Context) (result []*responses.FollowerResponse, customErr *apperror.CustomError)
	SeeAllFollowings(ctx context.Context) (result []*responses.FollowingResponse, customErr *apperror.CustomError)
}
