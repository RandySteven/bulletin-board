package handlers

import "net/http"

type IRelationHandler interface {
	AddFriend(w http.ResponseWriter, r *http.Request)
	GetAllFriends(w http.ResponseWriter, r *http.Request)
	GetFriendWithUserID(w http.ResponseWriter, r *http.Request)
}
