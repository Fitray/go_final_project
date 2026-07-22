package tasks_service

import (
	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
)

type TasksService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	AddTask(taskRequest core_domain.Task) (core_domain.NewTaskResponse, error)
	GetTasks(
		limit int,
	) (core_domain.GetTasksResponse, error)
	GetTasksByText(
		search string, limit int,
	) (core_domain.GetTasksResponse, error)
	GetTasksByDate(
		search string, limit int,
	) (core_domain.GetTasksResponse, error)
	GetTaskByID(
		id string,
	) (core_domain.Task, error)
	UpdateTask(taskRequest core_domain.Task) error
	CompleteTask(id string, nextDate string) error
	DeleteTask(id string) error
}

func NewTasksService(tasksRepository TasksRepository) TasksService {
	return TasksService{
		tasksRepository: tasksRepository,
	}
}
