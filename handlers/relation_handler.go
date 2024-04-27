package handlers

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"task_mission/entities/dtos/requests"
	"task_mission/enums"
	"task_mission/interfaces/handlers"
	"task_mission/interfaces/usecases"
	"task_mission/utils"
)

type RelationHandler struct {
	usecase usecases.IRelationUsecase
}

func (rh *RelationHandler) GetAllFollowers(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (rh *RelationHandler) GetAllFollowings(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (rh *RelationHandler) AddFriend(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.FriendRequest{}
		dataKey = `relation`
	)
	result, err := rh.usecase.AddFriend(ctx, request)
	if err != nil {
		return
	}
	utils.ResponseHandler(w, http.StatusCreated, `success to add friend`, &dataKey, result, nil)
}

func (rh *RelationHandler) GetAllFriends(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (rh *RelationHandler) GetFriendWithUserID(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ handlers.IRelationHandler = &RelationHandler{}

func NewRelationHandler(usecase usecases.IRelationUsecase) *RelationHandler {
	return &RelationHandler{
		usecase: usecase,
	}
}
