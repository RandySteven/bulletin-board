package apps

import (
	"github.com/gorilla/mux"
	"task_mission/middlewares"
)

func RegisterMiddleware(r *mux.Router) *mux.Router {
	r.Use(
		middlewares.CorsMiddleware,
		middlewares.LoggingMiddleware,
		middlewares.TimeoutMiddleware,
	)
	return r
}
