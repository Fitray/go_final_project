package tasks_handler

import (
	"net/http"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_server "github.com/Fitray/go_final_project/internal/core/server"
)

type TasksHandler struct {
	tasksService TasksService
}

type TasksService interface {
	NextDate(now, date, repeat string) (string, error)
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
}

func NewDateHandler(tasksService TasksService) TasksHandler {
	return TasksHandler{
		tasksService: tasksService,
	}
}

func (h *TasksHandler) Routes() []core_server.Route {
	return []core_server.Route{
		{
			Method:  http.MethodGet,
			Pattern: "/api/nextdate",
			Handler: h.NextDate,
		},
		{
			Method:  http.MethodPost,
			Pattern: "/api/task",
			Handler: h.AddTask,
		},
		{
			Method:  http.MethodGet,
			Pattern: "/api/tasks",
			Handler: h.GetTasks,
		},
		{
			Method:  http.MethodGet,
			Pattern: "/api/task",
			Handler: h.GetTask,
		},
		{
			Method:  http.MethodPut,
			Pattern: "/api/task",
			Handler: h.UpdateTask,
		},
	}
}
