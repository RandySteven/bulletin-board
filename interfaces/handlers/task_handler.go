package handlers

import "net/http"

type ITaskHandler interface {
	CreateTaskHandler(w http.ResponseWriter, r *http.Request)
	GetAllTasksHandler(w http.ResponseWriter, r *http.Request)
	GetTaskDetailHandler(w http.ResponseWriter, r *http.Request)
	TakeTaskHandler(w http.ResponseWriter, r *http.Request)
}
