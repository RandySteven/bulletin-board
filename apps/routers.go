package apps

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"task_mission/enums"
	"task_mission/middlewares"
)

type EndpointRouter struct {
	path    string
	handler func(w http.ResponseWriter, r *http.Request)
	method  string
}

func RegisterEndpointRouter(path, method string, handler func(w http.ResponseWriter, r *http.Request)) *EndpointRouter {
	return &EndpointRouter{path: path, handler: handler, method: method}
}

func NewEndpointRouters(h *Handlers) map[enums.RouterPrefix][]EndpointRouter {
	endpointRouters := make(map[enums.RouterPrefix][]EndpointRouter)

	endpointRouters[enums.AuthRouter] = []EndpointRouter{
		*RegisterEndpointRouter("/register", http.MethodPost, h.UserHandler.RegisterHandler),
		*RegisterEndpointRouter("/login", http.MethodPost, h.UserHandler.LoginHandler),
	}

	endpointRouters[enums.TaskRouter] = []EndpointRouter{
		*RegisterEndpointRouter("", http.MethodPost, h.TaskHandler.CreateTaskHandler),
		*RegisterEndpointRouter("", http.MethodGet, h.TaskHandler.GetAllTasksHandler),
		*RegisterEndpointRouter("/{id}", http.MethodGet, h.TaskHandler.GetTaskDetailHandler),
	}

	return endpointRouters
}

func (h *Handlers) InitRouter(r *mux.Router) {
	mapRouters := NewEndpointRouters(h)

	authRouter := r.PathPrefix(string(enums.AuthRouter)).Subrouter()
	for _, router := range mapRouters[enums.AuthRouter] {
		authRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(string(enums.AuthRouter))
	}

	taskRouter := r.PathPrefix(string(enums.TaskRouter)).Subrouter()
	taskRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range mapRouters[enums.TaskRouter] {
		//if router.path != "" && router.method != http.MethodGet {
		//	taskRouter.Use(middlewares.AuthenticationMiddleware)
		//	taskRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		//} else {
		//	taskRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		//}
		taskRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(string(enums.TaskRouter))
	}
}

func (router *EndpointRouter) RouterLog(prefix string) {
	//blue := color.New(color.FgBlue).SprintFunc()
	//red := color.New(color.FgRed).SprintFunc()
	//green := color.New(color.FgGreen).SprintFunc()

	log.Printf("%4s | %4s/ \n", router.method, prefix+router.path)
}
