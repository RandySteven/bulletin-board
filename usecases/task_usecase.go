package usecases

import (
	"context"
	"fmt"
	"log"
	"sync"
	"task_mission/apperror"
	"task_mission/entities/dtos/params"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
	"task_mission/entities/models"
	"task_mission/enums"
	"task_mission/interfaces/repositories"
	"task_mission/interfaces/usecases"
	"task_mission/pkg/firebases"
)

type taskUsecase struct {
	taskRepo           repositories.ITaskRepository
	rewardRepo         repositories.IRewardRepository
	taskRewardRepo     repositories.ITaskRewardRepository
	rewardCategoryRepo repositories.IRewardCategoryRepository
	userRepo           repositories.IUserRepository
	userProfileRepo    repositories.IUserProfileRepository
	userTaskRepo       repositories.IUserTaskRepository
	firebaseConf       firebases.Firebase
	uow                repositories.UnitOfWork
}

func (t *taskUsecase) GetTaskDetail(ctx context.Context, taskID uint64) (result *responses.TaskDetailResponse, customErr *apperror.CustomError) {
	var (
		uID         = ctx.Value(enums.UserID).(uint64)
		wg          = &sync.WaitGroup{}
		task        = &models.Task{}
		reward      = &models.Reward{}
		taskReward  = &models.TaskReward{}
		user        = &models.User{}
		userProfile = &models.UserProfile{}
		err         error
		errChan     = make(chan error, 2)
	)
	//t.firebaseConf.Store(ctx, "")
	wg.Add(2)
	go func() {
		defer wg.Done()
		user, err = t.userRepo.Find(ctx, uID)
		if err != nil {
			errChan <- err
			return
		}
		userProfile, err = t.userProfileRepo.FindByUserID(ctx, user.ID)
		if err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		task, err = t.taskRepo.Find(ctx, taskID)
		if err != nil {
			errChan <- err
			return
		}
		taskReward, err = t.taskRewardRepo.FindByTaskId(ctx, taskID)
		if err != nil {
			errChan <- err
			return
		}
		reward, err = t.rewardRepo.Find(ctx, taskReward.RewardID)
		if err != nil {
			errChan <- err
			return
		}
	}()

	wg.Wait()

	select {
	case err = <-errChan:
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get reward detail`, err)
	default:
		result = responses.NewTaskDetailResponse(task, user, userProfile, reward)
	}

	return result, nil
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
	result = responses.NewUserTaskResponse(userTask)
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
		task = request.ConvertTask(user.ID)
		log.Println(task)
		taskId, err = t.taskRepo.Save(ctx, task)
		if err != nil {
			ch <- err
			return nil
		}
		return taskId
	}()

	go func() *uint64 {
		defer wg.Done()
		reward := request.ConvertReward(user.ID)
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

func (t *taskUsecase) GetAllTasks(ctx context.Context, queryParam *params.TaskParam) (results []*responses.TaskListResponse, customErr *apperror.CustomError) {
	tasks, err := t.taskRepo.FindAll(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get all tasks`, err)
	}
	var (
		wg         sync.WaitGroup
		user       = &models.User{}
		reward     = &models.Reward{}
		taskReward = &models.TaskReward{}
		errCh      = make(chan error, len(tasks))
	)
	for _, task := range tasks {
		wg.Add(1)

		go func(task *models.Task) {
			defer wg.Done()
			user, err = t.userRepo.Find(ctx, task.UserID)
			if err != nil {
				errCh <- err
				return
			}

			taskReward, err = t.taskRewardRepo.FindByTaskId(ctx, task.ID)
			if err != nil {
				errCh <- err
				return
			}
			reward, err = t.rewardRepo.Find(ctx, taskReward.RewardID)
			if err != nil {
				errCh <- err
				return
			}

			result := responses.NewTaskListResponse(task, user, reward)
			results = append(results, result)
		}(task)
	}

	wg.Wait()

	select {
	case err = <-errCh:
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get all tasks`, err)
	default:
		return results, nil
	}
}

func (t *taskUsecase) UpdateTaskExpiryTime(ctx context.Context) (err error) {
	err = t.taskRepo.UpdateTasksExpiryDate(ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewTaskUsecase(
	taskRepo repositories.ITaskRepository,
	rewardRepo repositories.IRewardRepository,
	rewardCategoryRepo repositories.IRewardCategoryRepository,
	taskRewardRepo repositories.ITaskRewardRepository,
	userRepo repositories.IUserRepository,
	userProfileRepo repositories.IUserProfileRepository,
	userTaskRepo repositories.IUserTaskRepository,
	uow repositories.UnitOfWork,
) *taskUsecase {
	return &taskUsecase{
		taskRepo:           taskRepo,
		rewardRepo:         rewardRepo,
		rewardCategoryRepo: rewardCategoryRepo,
		taskRewardRepo:     taskRewardRepo,
		userRepo:           userRepo,
		userProfileRepo:    userProfileRepo,
		userTaskRepo:       userTaskRepo,
		uow:                uow,
	}
}

var _ usecases.ITaskUsecase = &taskUsecase{}
