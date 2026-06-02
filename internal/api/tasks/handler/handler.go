package tasks_handler

import (
	"net/http"

	core_auth "github.com/Fitray/go_final_project/internal/core/auth"
	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_middleware "github.com/Fitray/go_final_project/internal/core/middleware"
	core_server "github.com/Fitray/go_final_project/internal/core/server"
)

type TasksHandler struct {
	tasksService TasksService
}

type TasksService interface {
	NextDate(nowStr, startStr, repeat string) (string, error)
	AddTask(taskRequest core_domain.Task) (core_domain.NewTaskResponse, error)
	GetTasks(
		search string, limit int,
	) (core_domain.GetTasksResponse, error)
	GetTask(
		id string,
	) (core_domain.Task, error)
	UpdateTask(
		taskRequest core_domain.Task,
	) error
	CompleteTask(id string) error
	DeleteTask(id string) error
}

func NewDateHandler(tasksService TasksService) TasksHandler {
	return TasksHandler{
		tasksService: tasksService,
	}
}

func (h *TasksHandler) Routes(
	routes []core_server.Route,
	auth core_auth.Auth,
) []core_server.Route {
	for _, route := range []core_server.Route{
		{
			Method:  http.MethodGet,
			Pattern: "/api/nextdate",
			Handler: h.NextDate,
		},
		{
			Method:      http.MethodPost,
			Pattern:     "/api/task",
			Handler:     h.AddTask,
			Middlewares: core_middleware.GetAuthChain(auth),
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/api/tasks",
			Handler:     h.GetTasks,
			Middlewares: core_middleware.GetAuthChain(auth),
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/api/task",
			Handler:     h.GetTask,
			Middlewares: core_middleware.GetAuthChain(auth),
		},
		{
			Method:      http.MethodPut,
			Pattern:     "/api/task",
			Handler:     h.UpdateTask,
			Middlewares: core_middleware.GetAuthChain(auth),
		},
		{
			Method:      http.MethodPost,
			Pattern:     "/api/task/done",
			Handler:     h.CompleteTask,
			Middlewares: core_middleware.GetAuthChain(auth),
		},
		{
			Method:      http.MethodDelete,
			Pattern:     "/api/task",
			Handler:     h.DeleteTask,
			Middlewares: core_middleware.GetAuthChain(auth),
		},
	} {
		routes = append(routes, route)
	}
	return routes
}
