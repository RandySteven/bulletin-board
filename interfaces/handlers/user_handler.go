package handlers

import "net/http"

type IUserHandler interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	UserProfileHandler(w http.ResponseWriter, r *http.Request)
	UserDetailHandler(w http.ResponseWriter, r *http.Request)
}
