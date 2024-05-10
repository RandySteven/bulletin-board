package usecases

import (
	"context"
	"log"
	"sync"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
	"task_mission/entities/models"
	"task_mission/enums"
	"task_mission/interfaces/repositories"
	"task_mission/interfaces/usecases"
	"task_mission/utils"
)

type relationUsecase struct {
	uow          repositories.UnitOfWork
	relationRepo repositories.IRelationRepository
	userRepo     repositories.IUserRepository
	creditRepo   repositories.ICreditRepository
}

func (r *relationUsecase) SeeAllFollowers(ctx context.Context) (result []*responses.FollowerResponse, customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)
	followers, err := r.relationRepo.FindUserFollowers(ctx, userId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get followers`, err)
	}
	for _, follower := range followers {
		user, err := r.userRepo.Find(ctx, follower.FriendID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
		}
		credits, err := r.creditRepo.GetUserCredits(ctx, user.ID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get credits`, err)
		}
		creditAvg := utils.CreditsAverage(credits)
		response := &responses.FollowerResponse{
			ID:        user.ID,
			Name:      user.Name,
			UserName:  user.UserName,
			UserScore: creditAvg,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		}
		result = append(result, response)
	}
	return result, nil
}

func (r *relationUsecase) SeeAllFollowings(ctx context.Context) (result []*responses.FollowingResponse, customErr *apperror.CustomError) {
	userId := ctx.Value(enums.UserID).(uint64)
	followings, err := r.relationRepo.FindUserFollowings(ctx, userId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get followings`, err)
	}
	log.Println(followings[0])
	for _, following := range followings {
		friend, err := r.userRepo.Find(ctx, following.FriendID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get friend`, err)
		}
		credits, err := r.creditRepo.GetUserCredits(ctx, friend.ID)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get credits`, err)
		}
		creditAvg := utils.CreditsAverage(credits)
		response := &responses.FollowingResponse{
			ID:        friend.ID,
			Name:      friend.Name,
			UserName:  friend.UserName,
			UserScore: creditAvg,
			CreatedAt: friend.CreatedAt,
			UpdatedAt: friend.UpdatedAt,
			DeletedAt: friend.DeletedAt,
		}
		result = append(result, response)
	}
	return result, nil
}

func (r *relationUsecase) AddFriend(ctx context.Context, request *requests.FriendRequest) (result *responses.CreateFriendResponse, customErr *apperror.CustomError) {
	var (
		wg     sync.WaitGroup
		user   = &models.User{}
		friend = &models.User{}
		userId = ctx.Value(enums.UserID).(uint64)
		errCh  = make(chan *apperror.CustomError, 2)
		err    error
	)
	log.Println(userId)
	wg.Add(2)

	go func() {
		defer wg.Done()
		user, err = r.userRepo.Find(ctx, userId)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		friend, err = r.userRepo.Find(ctx, request.FriendID)
		if err != nil {
			errCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to get friend`, err)
			return
		}
	}()

	wg.Wait()
	close(errCh)

	for customErr = range errCh {
		if &customErr != nil {
			return nil, customErr
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
	creditRepo repositories.ICreditRepository,
) *relationUsecase {
	return &relationUsecase{
		uow:          uow,
		relationRepo: relationRepo,
		userRepo:     userRepo,
		creditRepo:   creditRepo,
	}
}

var _ usecases.IRelationUsecase = &relationUsecase{}
