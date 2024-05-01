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

type TaskHandler struct {
	taskUsecase usecases.ITaskUsecase
}

func (t *TaskHandler) TakeTaskHandler(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
}

func (t *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.CreateTaskRequest{}
		dataKey = `task_reward`
	)

	if err := utils.BindJSON(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, err.Error(), nil, nil, err)
		return
	}

	result, err := t.taskUsecase.CreateTask(ctx, request)
	if err != nil {
		utils.ResponseHandler(w, err.ErrCode(), `failed to create task`, nil, nil, err)
		return
	}

	utils.ResponseHandler(w, http.StatusCreated, `success to create task`, &dataKey, result, nil)
}

func (t *TaskHandler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `tasks`
	)
	result, err := t.taskUsecase.GetAllTasks(ctx)
	if err != nil {
		utils.ResponseHandler(w, err.ErrCode(), `failed to get all tasks`, nil, nil, err)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success to create task`, &dataKey, result, nil)
}

func (t *TaskHandler) GetTaskDetailHandler(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
}

var _ handlers.ITaskHandler = &TaskHandler{}

func NewTaskHandler(taskUsecase usecases.ITaskUsecase) *TaskHandler {
	return &TaskHandler{taskUsecase: taskUsecase}
}
