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
	AddTask(taskRequest core_domain.TaskRequest) (core_domain.TaskResponce, error)
}

func NewTasksService(tasksRepository TasksRepository) TasksService {
	return TasksService{
		tasksRepository: tasksRepository,
	}
}
