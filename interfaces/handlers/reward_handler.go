package handlers

import "net/http"

type IRewardHandler interface {
	GetAllRewards(w http.ResponseWriter, r *http.Request)
	GetReward(w http.ResponseWriter, r *http.Request)
	UpdateReward(w http.ResponseWriter, r *http.Request)
	DeleteReward(w http.ResponseWriter, r *http.Request)
}
