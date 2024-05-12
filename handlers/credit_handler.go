package handlers

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"task_mission/entities/dtos/requests"
	"task_mission/enums"
	"task_mission/interfaces/handlers"
	"task_mission/interfaces/usecases"
	"task_mission/utils"
)

type CreditHandler struct {
	creditUsecase usecases.ICreditUseCase
}

func (c *CreditHandler) GiveCredit(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.CreditRequest{}
		dataKey = `credit`
	)
	if err := utils.BindJSON(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `invalid_json`, nil, nil, err)
		return
	}
	result, customErr := c.creditUsecase.GiveCredit(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal_error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusCreated, `success submit credit`, &dataKey, result, nil)
}

func (c *CreditHandler) SeeUserCredit(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `credits`
	)
	userIdStr := r.URL.Query().Get(`userId`)
	if userIdStr == `` {
		result, customErr := c.creditUsecase.SeeAllCredits(ctx)
		if customErr != nil {
			utils.ResponseHandler(w, customErr.ErrCode(), `internal_error`, nil, nil, customErr)
			return
		}
		utils.ResponseHandler(w, http.StatusOK, `success submit credit`, &dataKey, result, nil)
		return
	}
	uId, err := strconv.Atoi(userIdStr)
	if err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `invalid_json`, nil, nil, err)
		return
	}
	userId := uint64(uId)
	//if userId == 0 {
	//}
	result, customErr := c.creditUsecase.SeeUserCredit(ctx, userId)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `internal_error`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get user credit`, &dataKey, result, nil)
}

var _ handlers.ICreditHandler = &CreditHandler{}

func NewCreditHandler(creditUsecase usecases.ICreditUseCase) *CreditHandler {
	return &CreditHandler{
		creditUsecase: creditUsecase,
	}
}
