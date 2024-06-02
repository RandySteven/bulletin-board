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

	endpointRouters[enums.BasicRouter] = []EndpointRouter{
		*RegisterEndpointRouter("/hello", http.MethodGet, h.DevHandler.HelloDev),
		*RegisterEndpointRouter("/check-health", http.MethodGet, h.DevHandler.HealthCheck),
		*RegisterEndpointRouter("/email", http.MethodPost, h.DevHandler.SendTestEmail),
	}

	endpointRouters[enums.AuthRouter] = []EndpointRouter{
		*RegisterEndpointRouter("/register", http.MethodPost, h.UserHandler.RegisterHandler),
		*RegisterEndpointRouter("/login", http.MethodPost, h.UserHandler.LoginHandler),
	}

	endpointRouters[enums.UserRouter] = []EndpointRouter{
		*RegisterEndpointRouter("", http.MethodGet, h.UserHandler.UserProfileHandler),
		*RegisterEndpointRouter("/{id}", http.MethodGet, h.UserHandler.UserDetailHandler),
	}

	endpointRouters[enums.TaskRouter] = []EndpointRouter{
		*RegisterEndpointRouter("", http.MethodPost, h.TaskHandler.CreateTaskHandler),
		*RegisterEndpointRouter("", http.MethodGet, h.TaskHandler.GetAllTasksHandler),
		*RegisterEndpointRouter("/{id}", http.MethodGet, h.TaskHandler.GetTaskDetailHandler),
	}

	endpointRouters[enums.RelationRouter] = []EndpointRouter{
		*RegisterEndpointRouter("", http.MethodPost, h.RelationHandler.AddFriend),
		*RegisterEndpointRouter("/followers", http.MethodGet, h.RelationHandler.GetAllFollowers),
		*RegisterEndpointRouter("/followings", http.MethodGet, h.RelationHandler.GetAllFollowings),
	}

	endpointRouters[enums.RewardRouter] = []EndpointRouter{
		*RegisterEndpointRouter("", http.MethodGet, h.RewardHandler.GetAllRewards),
		*RegisterEndpointRouter("/{id}", http.MethodGet, h.RewardHandler.GetReward),
		*RegisterEndpointRouter("/{id}", http.MethodPut, h.RewardHandler.UpdateReward),
		*RegisterEndpointRouter("/{id}", http.MethodDelete, h.RewardHandler.DeleteReward),
	}

	endpointRouters[enums.CategoryRouter] = []EndpointRouter{
		*RegisterEndpointRouter("", http.MethodGet, h.CategoryHandler.GetAllCategories),
		*RegisterEndpointRouter("/{id}", http.MethodGet, h.CategoryHandler.GetCategory),
		*RegisterEndpointRouter("", http.MethodPost, h.CategoryHandler.AddCategory),
	}

	endpointRouters[enums.CreditRouter] = []EndpointRouter{
		*RegisterEndpointRouter("", http.MethodPost, h.CreditHandler.GiveCredit),
		*RegisterEndpointRouter("", http.MethodGet, h.CreditHandler.SeeUserCredit),
	}

	return endpointRouters
}

func (h *Handlers) InitRouter(r *mux.Router) {
	mapRouters := NewEndpointRouters(h)

	basicRouter := r.PathPrefix(enums.BasicRouter.ToString()).Subrouter()
	for _, router := range mapRouters[enums.BasicRouter] {
		basicRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.BasicRouter.ToString())
	}

	authRouter := r.PathPrefix(enums.AuthRouter.ToString()).Subrouter()
	for _, router := range mapRouters[enums.AuthRouter] {
		authRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.AuthRouter.ToString())
	}

	userRouter := r.PathPrefix(enums.UserRouter.ToString()).Subrouter()
	userRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range mapRouters[enums.UserRouter] {
		userRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.UserRouter.ToString())
	}

	taskRouter := r.PathPrefix(enums.TaskRouter.ToString()).Subrouter()
	taskRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range mapRouters[enums.TaskRouter] {
		taskRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.TaskRouter.ToString())
	}

	relationRouter := r.PathPrefix(enums.RelationRouter.ToString()).Subrouter()
	relationRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range mapRouters[enums.RelationRouter] {
		relationRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.RelationRouter.ToString())
	}

	rewardRouter := r.PathPrefix(enums.RewardRouter.ToString()).Subrouter()
	rewardRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range mapRouters[enums.RewardRouter] {
		rewardRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.RewardRouter.ToString())
	}

	categoriesRouter := r.PathPrefix(enums.CategoryRouter.ToString()).Subrouter()
	categoriesRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range mapRouters[enums.CategoryRouter] {
		categoriesRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.CategoryRouter.ToString())
	}

	creditsRouter := r.PathPrefix(enums.CreditRouter.ToString()).Subrouter()
	creditsRouter.Use(middlewares.AuthenticationMiddleware)
	for _, router := range mapRouters[enums.CreditRouter] {
		creditsRouter.HandleFunc(router.path, router.handler).Methods(router.method)
		router.RouterLog(enums.CreditRouter.ToString())
	}
}

func (router *EndpointRouter) RouterLog(prefix string) {
	//blue := color.New(color.FgBlue).SprintFunc()
	//red := color.New(color.FgRed).SprintFunc()
	//green := color.New(color.FgGreen).SprintFunc()

	log.Printf("%4s | %4s/ \n", router.method, prefix+router.path)
}
