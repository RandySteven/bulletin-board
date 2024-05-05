package handlers

import "net/http"

type ICategoryHandler interface {
	AddCategory(w http.ResponseWriter, r *http.Request)
	GetAllCategories(w http.ResponseWriter, r *http.Request)
	GetCategory(w http.ResponseWriter, r *http.Request)
}
