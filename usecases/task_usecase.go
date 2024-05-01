package usecases

import (
	"context"
	"fmt"
	"log"
	"sync"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
	"task_mission/entities/models"
	"task_mission/enums"
	"task_mission/interfaces/repositories"
	"task_mission/interfaces/usecases"
)

type taskUsecase struct {
	taskRepo           repositories.ITaskRepository
	rewardRepo         repositories.IRewardRepository
	taskRewardRepo     repositories.ITaskRewardRepository
	rewardCategoryRepo repositories.IRewardCategoryRepository
	userRepo           repositories.IUserRepository
	userTaskRepo       repositories.IUserTaskRepository
	uow                repositories.UnitOfWork
}

func (t *taskUsecase) TakeTask(ctx context.Context, taskID uint64) (result *responses.UserTaskResponse, err error) {
	var (
		wg    sync.WaitGroup
		user  = &models.User{}
		task  = &models.Task{}
		errCh = make(chan error, 2)
	)

	t.uow, err = t.uow.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err = t.uow.Rollback()
			if err != nil {
				return
			}
		}
		err = t.uow.Commit()
	}()
	userId := ctx.Value(enums.UserID).(uint64)

	wg.Add(2)

	go func() {
		defer wg.Done()
		user, err = t.userRepo.Find(ctx, userId)
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		task, err = t.taskRepo.Find(ctx, taskID)
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to process the concurrency`, err)
		}
	}

	userTask := &models.UserTask{
		UserID: user.ID,
		TaskID: task.ID,
	}
	userTaskId, err := t.userTaskRepo.Save(ctx, userTask)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create user task`, err)
	}
	userTask, err = t.userTaskRepo.Find(ctx, *userTaskId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to find user task`, err)
	}
	result = &responses.UserTaskResponse{
		ID:        userTask.ID,
		TaskID:    userTask.TaskID,
		UserID:    userTask.UserID,
		CreatedAt: userTask.CreatedAt,
		UpdatedAt: userTask.UpdatedAt,
		DeletedAt: userTask.DeletedAt,
	}
	return result, nil
}

func (t *taskUsecase) CreateTask(ctx context.Context, request *requests.CreateTaskRequest) (result *responses.TaskResponse, customErr *apperror.CustomError) {
	uow, err := t.uow.Begin(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to begin trx`, err)
	}
	t.uow = uow
	defer func() {
		if err != nil {
			err := t.uow.Rollback()
			if err != nil {
				return
			}
			return
		}
		err = t.uow.Commit()

	}()
	uId := ctx.Value(enums.UserID).(uint64)
	user, err := t.userRepo.Find(ctx, uId)
	if err != nil {
		if user == nil {
			return nil, apperror.NewCustomError(apperror.ErrNotFound, `user not found`, fmt.Errorf(`not found user with id : %d`, uId))
		}
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
	}

	var (
		wg       sync.WaitGroup
		ch       = make(chan error, 2)
		taskId   = new(uint64)
		rewardId = new(uint64)
		task     = &models.Task{}
	)
	wg.Add(2)

	go func() *uint64 {
		defer wg.Done()
		task = request.ConvertTask()
		task.UserID = user.ID
		task.Status = enums.Open
		taskId, err = t.taskRepo.Save(ctx, task)
		if err != nil {
			ch <- err
			return nil
		}
		return taskId
	}()

	go func() *uint64 {
		defer wg.Done()
		reward := request.ConvertReward()
		reward.UserID = user.ID
		rewardId, err = t.rewardRepo.Save(ctx, reward)
		if err != nil {
			ch <- err
			return nil
		}
		for _, categoryId := range request.Categories.IDs {
			rewardCategory := &models.RewardCategory{
				CategoryID: categoryId,
				RewardID:   *rewardId,
			}
			_, err := t.rewardCategoryRepo.Save(ctx, rewardCategory)
			if err != nil {
				ch <- err
				return nil
			}
		}
		return rewardId
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for errorCh := range ch {
		if errorCh != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to process task creation`, errorCh)
		}
	}
	taskReward := &models.TaskReward{
		TaskID:   *taskId,
		RewardID: *rewardId,
	}
	taskRewardId, err := t.taskRewardRepo.Save(ctx, taskReward)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to save reward repo`, err)
	}

	result = &responses.TaskResponse{
		ID:       *taskRewardId,
		TaskID:   *taskId,
		RewardID: *rewardId,
		UserID:   user.ID,
	}

	return result, nil
}

func (t *taskUsecase) GetAllTasks(ctx context.Context) (results []*responses.TaskListResponse, customErr *apperror.CustomError) {
	tasks, err := t.taskRepo.FindAll(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get all tasks`, err)
	}
	for _, task := range tasks {
		var (
			//wg         sync.WaitGroup
			user       = &models.User{}
			reward     = &models.Reward{}
			taskReward = &models.TaskReward{}
			//errCh      = make(chan error)
		)
		log.Println("user id : ", task.UserID)
		user, err = t.userRepo.Find(ctx, task.UserID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
		}

		taskReward, err = t.taskRewardRepo.FindByTaskId(ctx, task.ID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get task reward`, err)
		}
		log.Println("task reward : ", taskReward.RewardID)
		reward, err = t.rewardRepo.Find(ctx, taskReward.RewardID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get reward repo`, err)
		}
		result := &responses.TaskListResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			ExpiryDate:  task.ExpiredDate.String(),
			UserID:      user.ID,
			UserName:    user.UserName,
			RewardID:    reward.ID,
			Reward:      reward.Name,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
			DeletedAt:   task.DeletedAt,
		}
		results = append(results, result)
	}
	return results, nil
}

func NewTaskUsecase(
	taskRepo repositories.ITaskRepository,
	rewardRepo repositories.IRewardRepository,
	rewardCategoryRepo repositories.IRewardCategoryRepository,
	taskRewardRepo repositories.ITaskRewardRepository,
	userRepo repositories.IUserRepository,
	userTaskRepo repositories.IUserTaskRepository,
	uow repositories.UnitOfWork,
) *taskUsecase {
	return &taskUsecase{
		taskRepo:           taskRepo,
		rewardRepo:         rewardRepo,
		rewardCategoryRepo: rewardCategoryRepo,
		taskRewardRepo:     taskRewardRepo,
		userRepo:           userRepo,
		userTaskRepo:       userTaskRepo,
		uow:                uow,
	}
}

var _ usecases.ITaskUsecase = &taskUsecase{}
