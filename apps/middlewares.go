package apps

import (
	"github.com/gorilla/mux"
	"task_mission/middlewares"
)

func RegisterMiddleware(r *mux.Router) *mux.Router {
	r.Use(
		middlewares.LoggingMiddleware,
		middlewares.CorsMiddleware,
		middlewares.TimeoutMiddleware,
	)
	return r
}
