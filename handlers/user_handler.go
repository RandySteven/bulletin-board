package handlers

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"sync"
	"task_mission/entities/dtos/requests"
	"task_mission/enums"
	"task_mission/interfaces/handlers"
	"task_mission/interfaces/usecases"
	email2 "task_mission/pkg/email"
	"task_mission/utils"
	"time"
)

type UserHandler struct {
	usecase usecases.IUserUsecase
}

func (u *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.UserRegisterRequest{}
		dataKey = `user`
		wg      sync.WaitGroup
	)

	if err := utils.BindJSON(r, &request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `failed to register user`, nil, nil, err)
		return
	}

	result, err := u.usecase.RegisterUser(ctx, request)
	if err != nil {
		utils.ResponseHandler(w, http.StatusCreated, `failed to register user`, nil, nil, err)
		return
	}

	metadata := map[string]interface{}{
		"name":        result.Name,
		"username":    result.UserName,
		"email":       request.Email,
		"joined_date": time.Now(),
	}

	email := email2.NewEmail(request.Email, "Register Email", metadata)

	wg.Add(1)
	go func() {
		defer wg.Done()
		email.SendEmailRegister()
		utils.ResponseHandler(w, http.StatusCreated, `success register user`, &dataKey, result, nil)
	}()
	wg.Wait()
}

func (u *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
	var (
		request = &requests.UserLoginRequest{}
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `user`
	)

	if err := utils.BindJSON(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, err.Error(), nil, nil, err)
		return
	}

	token, customErr := u.usecase.LoginUser(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), customErr.Error(), nil, nil, customErr)
		return
	}

	utils.ResponseHandler(w, http.StatusOK, `Success login`, &dataKey, token, nil)
}

func (u *UserHandler) UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserHandler) UserDetailHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ handlers.IUserHandler = &UserHandler{}

func NewUserHandler(usecase usecases.IUserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}
