package usecases

import (
	"context"
	"sync"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
	"task_mission/entities/models"
	"task_mission/enums"
	"task_mission/interfaces/repositories"
	"task_mission/interfaces/usecases"
)

type relationUsecase struct {
	uow          repositories.UnitOfWork
	relationRepo repositories.IRelationRepository
	userRepo     repositories.IUserRepository
}

func (r *relationUsecase) AddFriend(ctx context.Context, request *requests.FriendRequest) (result *responses.CreateFriendResponse, err error) {
	var (
		wg     sync.WaitGroup
		user   = &models.User{}
		friend = &models.User{}
		userId = ctx.Value(enums.UserID).(uint64)
		errCh  = make(chan error, 2)
	)
	wg.Add(2)

	go func() {
		defer wg.Done()
		user, err = r.userRepo.Find(ctx, userId)
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		friend, err = r.userRepo.Find(ctx, request.FriendID)
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
			return nil, apperror.NewCustomError(apperror.ErrNotFound, `failed to find user`, err)
		}
	}

	relation := &models.Relation{
		UserID:   user.ID,
		FriendID: friend.ID,
	}

	relationId, err := r.relationRepo.Save(ctx, relation)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create user relation`, err)
	}

	relation, err = r.relationRepo.Find(ctx, *relationId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get relation`, err)
	}

	result = &responses.CreateFriendResponse{
		ID:         *relationId,
		UserID:     user.ID,
		UserName:   user.UserName,
		FriendID:   friend.ID,
		FriendName: friend.UserName,
		CreatedAt:  relation.CreatedAt,
		UpdatedAt:  relation.UpdatedAt,
		DeletedAt:  relation.DeletedAt,
	}

	return
}

func (r *relationUsecase) SeeAllFriends(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func NewRelationUsecase(
	uow repositories.UnitOfWork,
	relationRepo repositories.IRelationRepository,
	userRepo repositories.IUserRepository,
) *relationUsecase {
	return &relationUsecase{
		uow:          uow,
		relationRepo: relationRepo,
		userRepo:     userRepo,
	}
}

var _ usecases.IRelationUsecase = &relationUsecase{}
