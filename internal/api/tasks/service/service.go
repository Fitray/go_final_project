package tasks_service

import (
	"time"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
)

type TasksService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	NextDate(
		now, start time.Time, params core_domain.NextDateParams,
	) (string, error)
	AddTask(taskRequest core_domain.Task) (core_domain.NewTaskResponse, error)
	GetTasks_NoSearch(
		limit int,
	) (core_domain.GetTasksResponse, error)
	GetTasks_TextSearch(
		search string, limit int,
	) (core_domain.GetTasksResponse, error)
	GetTasks_DateSearch(
		search string, limit int,
	) (core_domain.GetTasksResponse, error)
	GetTasks_FromID(
		id string,
	) (core_domain.Task, error)
	UpdateTask(taskRequest core_domain.Task) error
}

func NewTasksService(tasksRepository TasksRepository) TasksService {
	return TasksService{
		tasksRepository: tasksRepository,
	}
}
