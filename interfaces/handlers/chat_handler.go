package handlers

import "net/http"

type IChatHandler interface {
	CreateRoom(w http.ResponseWriter, r *http.Request)
	GetAllRooms(w http.ResponseWriter, r *http.Request)
}
