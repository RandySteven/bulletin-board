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
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `followers`
	)
	result, customErr := rh.usecase.SeeAllFollowers(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get followers`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get all followers`, &dataKey, result, nil)
}

func (rh *RelationHandler) GetAllFollowings(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `followers`
	)
	result, customErr := rh.usecase.SeeAllFollowings(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get followings`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get all followings`, &dataKey, result, nil)
}

func (rh *RelationHandler) AddFriend(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.FriendRequest{}
		dataKey = `relation`
	)
	if err := utils.BindJSON(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `fail to parse request`, nil, nil, err)
		return
	}
	result, customErr := rh.usecase.AddFriend(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get friend`, nil, nil, customErr)
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
