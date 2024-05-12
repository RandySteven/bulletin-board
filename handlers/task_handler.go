package handlers

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		params  = mux.Vars(r)
		dataKey = `task`
	)
	idInt, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, err.Error(), nil, nil, err)
		return
	}
	id := uint64(idInt)
	res, controllerErr := t.taskUsecase.TakeTask(ctx, id)
	if controllerErr != nil {
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success update`, &dataKey, res, nil)
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
	var (
		rID            = uuid.NewString()
		ctx            = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey        = `task`
		taskId  uint64 = 0
	)
	param := mux.Vars(r)
	taskIdInt, err := strconv.Atoi(param["id"])
	if err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, err.Error(), nil, nil, err)
		return
	}
	taskId = uint64(taskIdInt)
	result, controllerErr := t.taskUsecase.GetTaskDetail(ctx, taskId)
	if controllerErr != nil {
		utils.ResponseHandler(w, controllerErr.ErrCode(), `failed to get all tasks`, nil, nil, controllerErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success to get task`, &dataKey, result, nil)

}

var _ handlers.ITaskHandler = &TaskHandler{}

func NewTaskHandler(taskUsecase usecases.ITaskUsecase) *TaskHandler {
	return &TaskHandler{taskUsecase: taskUsecase}
}
