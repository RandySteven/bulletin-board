package usecases

import (
	"context"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
)

type ITaskUsecase interface {
	CreateTask(ctx context.Context, request *requests.CreateTaskRequest) (result *responses.TaskResponse, customErr *apperror.CustomError)
	GetAllTasks(ctx context.Context) (results []*responses.TaskListResponse, customErr *apperror.CustomError)
	TakeTask(ctx context.Context, taskID uint64) (result *responses.UserTaskResponse, err error)
	GetTaskDetail(ctx context.Context, taskID uint64) (result *responses.TaskDetailResponse, customErr *apperror.CustomError)
	UpdateTaskExpiryTime(ctx context.Context) (err error)
}
