package handlers

import (
	"net/http"
	"task_mission/interfaces/handlers"
	"task_mission/interfaces/usecases"
)

type CreditHandler struct {
	creditUsecase usecases.ICreditUseCase
}

func (c *CreditHandler) GiveCredit(w http.ResponseWriter, r *http.Request) {
}

func (c *CreditHandler) SeeUserCredit(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ handlers.ICreditHandler = &CreditHandler{}

func NewCreditHandler(creditUsecase usecases.ICreditUseCase) *CreditHandler {
	return &CreditHandler{
		creditUsecase: creditUsecase,
	}
}
