package handlers

import (
	"net/http"
	"task_mission/interfaces/handlers"
	"task_mission/interfaces/usecases"
)

type RewardHandler struct {
	usecase usecases.IRewardUsecase
}

func (rh *RewardHandler) GetAllRewards(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (rh *RewardHandler) GetReward(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
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
