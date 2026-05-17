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
	AddTask(taskRequest core_domain.TaskRequest) (core_domain.TaskResponce, error)
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
	}
}
