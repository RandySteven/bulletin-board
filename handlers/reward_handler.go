package handlers

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"task_mission/enums"
	"task_mission/interfaces/handlers"
	"task_mission/interfaces/usecases"
	"task_mission/utils"
)

type RewardHandler struct {
	usecase usecases.IRewardUsecase
}

func (rh *RewardHandler) GetAllRewards(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `rewards`
	)
	result, customErr := rh.usecase.GetAllRewards(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get all rewards`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get all rewards`, &dataKey, result, nil)
}

func (rh *RewardHandler) GetReward(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `reward`
		params  = mux.Vars(r)
	)
	id, err := strconv.Atoi(params[`id`])
	if err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `invalid reward id`, nil, nil, err)
		return
	}
	rewardId := uint64(id)
	result, customErr := rh.usecase.GetRewardByID(ctx, rewardId)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get reward`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get reward`, &dataKey, result, customErr)
}

func (rh *RewardHandler) UpdateReward(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (rh *RewardHandler) DeleteReward(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewRewardHandler(usecase usecases.IRewardUsecase) *RewardHandler {
	return &RewardHandler{
		usecase: usecase,
	}
}

var _ handlers.IRewardHandler = &RewardHandler{}
