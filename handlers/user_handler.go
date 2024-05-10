package handlers

import (
	"context"
	"github.com/google/uuid"
	"log"
	"net/http"
	"task_mission/entities/dtos/requests"
	"task_mission/enums"
	"task_mission/interfaces/handlers"
	"task_mission/interfaces/usecases"
	"task_mission/utils"
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
	log.Println("expected result : ", result)

	utils.ResponseHandler(w, http.StatusCreated, `success register user`, &dataKey, result, nil)
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
