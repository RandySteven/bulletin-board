package usecases

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"sync"
	"task_mission/apperror"
	"task_mission/entities/dtos/requests"
	"task_mission/entities/dtos/responses"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/interfaces/usecases"
	"task_mission/pkg/securities"
	"task_mission/utils"
	"time"
)

type userUsecase struct {
	uow             repositories.UnitOfWork
	userRepo        repositories.IUserRepository
	userProfileRepo repositories.IUserProfileRepository
	userRoleRepo    repositories.IUserRoleRepository
	userCreditRepo  repositories.IUserCreditRepository
}

func (u *userUsecase) VerifyUser(ctx context.Context, id uint64) (result *models.User, customErr *apperror.CustomError) {
	user, err := u.userRepo.Find(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrNotFound, `failed to find user`, err)
	}
	user.IsVerified = true
	user, err = u.userRepo.Update(ctx, user)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to update user`, err)
	}
	return user, nil
}

func (u *userUsecase) RegisterUser(ctx context.Context, register *requests.UserRegisterRequest) (result *models.User, customErr *apperror.CustomError) {
	uow, err := u.uow.Begin(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed in init transaction`, err)
	}
	u.uow = uow
	defer func() {
		if err != nil {
			u.uow.Rollback()
			return
		}
		err = u.uow.Commit()
	}()

	userDoB, err := utils.StringToDate(register.DateOfBirth)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrBadRequest, `date of birth is invalid format`, err)
	}
	user := &models.User{
		Name:        utils.UserFullName(register.FirstName, register.LastName),
		UserName:    register.UserName,
		DateOfBirth: userDoB,
		Gender:      register.Gender,
	}
	userId, err := u.userRepo.Save(ctx, user)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create user`, err)
	}

	var (
		wg sync.WaitGroup
		ch = make(chan error, 3)
	)
	wg.Add(3)

	go func() {
		defer wg.Done()
		password, err := utils.HashPassword(register.Password)
		if err != nil {
			ch <- err
			return
		}
		userProfile := &models.UserProfile{
			Email:    register.Email,
			Password: password,
			UserID:   *userId,
		}
		_, err = u.userProfileRepo.Save(ctx, userProfile)
		if err != nil {
			ch <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		userRole := &models.UserRole{
			UserID: *userId,
			RoleID: 3,
		}
		_, err = u.userRoleRepo.Save(ctx, userRole)
		if err != nil {
			ch <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		userCredit := &models.UserCredit{
			UserID: *userId,
		}
		_, err := u.userCreditRepo.Save(ctx, userCredit)
		if err != nil {
			ch <- err
			return
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for err := range ch {
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to perform user operation`, err)
		}
	}

	result, err = u.userRepo.Find(ctx, *userId)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get user`, err)
	}
	log.Println(result)

	return result, nil
}

func (u *userUsecase) LoginUser(ctx context.Context, login *requests.UserLoginRequest) (result *responses.UserLoginResponse, customErr *apperror.CustomError) {
	user, err := u.userProfileRepo.FindByEmail(ctx, login.Email)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrNotFound, `email find`, fmt.Errorf(`user is not found`))
	}
	existed := utils.ComparePassword(login.Password, user.Password)
	if !existed {
		return nil, apperror.NewCustomError(apperror.ErrNotFound, `password find`, fmt.Errorf(`user is not found`))
	}
	userRole, err := u.userRoleRepo.FindByUserID(ctx, user.UserID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrNotFound, `user role`, fmt.Errorf(`user is not found`))
	}

	claims := &securities.JWTClaim{
		UserID: user.UserID,
		RoleID: userRole.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Applications",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(securities.JWT_KEY)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to generate token`, err)
	}
	result = &responses.UserLoginResponse{
		Token: token,
	}
	return result, nil
}

func NewUserUsecase(
	uow repositories.UnitOfWork,
	userRepo repositories.IUserRepository,
	userProfileRepo repositories.IUserProfileRepository,
	userRoleRepo repositories.IUserRoleRepository,
	userCreditRepo repositories.IUserCreditRepository) *userUsecase {
	return &userUsecase{
		uow:             uow,
		userRepo:        userRepo,
		userProfileRepo: userProfileRepo,
		userRoleRepo:    userRoleRepo,
		userCreditRepo:  userCreditRepo,
	}
}

var _ usecases.IUserUsecase = &userUsecase{}
