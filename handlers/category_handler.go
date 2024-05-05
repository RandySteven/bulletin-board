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

type CategoryHandler struct {
	usecase usecases.ICategoryUsecase
}

func (c *CategoryHandler) AddCategory(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.CategoryRequest{}
		dataKey = `category`
	)
	result, customErr := c.usecase.CreateCategory(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to create category`, nil, nil, customErr)
	}
	utils.ResponseHandler(w, http.StatusCreated, `success create category`, &dataKey, result, nil)
}

func (c *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `categories`
	)
	result, customErr := c.usecase.GetAllCategories(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get categories`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get all categories`, &dataKey, result, nil)
}

func (c *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, "application/json")
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `category`
		params  = mux.Vars(r)
	)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `failed to parse category id`, nil, nil, err)
		return
	}
	categoryId := uint64(id)
	result, customErr := c.usecase.GetCategoryById(ctx, categoryId)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get category`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get category`, &dataKey, result, nil)
}

func NewCategoryHandler(usecase usecases.ICategoryUsecase) *CategoryHandler {
	return &CategoryHandler{usecase: usecase}
}

var _ handlers.ICategoryHandler = &CategoryHandler{}
