package apps

import (
	"github.com/gorilla/mux"
	"task_mission/middlewares"
)

func RegisterMiddleware(r *mux.Router) *mux.Router {
	//middlewares.CorsMiddleware(r)
	//middlewares.TimeoutMiddleware(r)
	//middlewares.LoggingMiddleware(r)
	r.Use(middlewares.LoggingMiddleware, middlewares.CorsMiddleware)
	return r
}
