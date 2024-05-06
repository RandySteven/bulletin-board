package handlers

import "net/http"

type ICreditHandler interface {
	GiveCredit(w http.ResponseWriter, r *http.Request)
	SeeUserCredit(w http.ResponseWriter, r *http.Request)
}
