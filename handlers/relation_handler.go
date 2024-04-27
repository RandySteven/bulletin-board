package handlers

import (
	"net/http"
	"task_mission/interfaces/handlers"
)

type RelationHandler struct {
}

func (rh *RelationHandler) AddFriend(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
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
