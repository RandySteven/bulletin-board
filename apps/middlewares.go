package apps

import (
	"github.com/gorilla/mux"
	"task_mission/middlewares"
)

func RegisterMiddleware(r *mux.Router) {
	middlewares.CorsMiddleware(r)
	middlewares.TimeoutMiddleware(r)
}
