package tasks_service

import (
	"fmt"
	"time"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
)

func (s *TasksService) CheckTask(
	taskRequest core_domain.Task,
) (core_domain.Task, error) {
	if taskRequest.Date == "" {
		taskRequest.Date = time.Now().Format("20060102")
	}

	t, err := time.Parse("20060102", taskRequest.Date)
	if err != nil {
		return taskRequest,
			fmt.Errorf("%w:%w", err, core_errors.ErrBadRequest)
	}

	if taskRequest.Title == "" {
		return taskRequest,
			fmt.Errorf("invalid title: %w", core_errors.ErrBadRequest)
	}

	todayStr := time.Now().Format("20060102")
	today, err := time.Parse("20060102", todayStr)

	if err != nil {
		return taskRequest, err
	}

	if t.Before(today) {
		if taskRequest.Repeat == "" {
			taskRequest.Date = todayStr
		} else {
			nextDate, err := s.NextDate(
				todayStr,
				taskRequest.Date,
				taskRequest.Repeat,
			)
			if err != nil {
				return taskRequest,
					fmt.Errorf("%w:%w", err, core_errors.ErrBadRequest)
			}
			taskRequest.Date = nextDate
		}
	}

	return taskRequest, nil
}

func (s *TasksService) AddTask(
	taskRequest core_domain.Task,
) (core_domain.NewTaskResponse, error) {
	taskRequest, err := s.CheckTask(taskRequest)
	if err != nil {
		return core_domain.NewTaskResponse{}, err
	}

	resp, err := s.tasksRepository.AddTask(taskRequest)
	if err != nil {
		return core_domain.NewTaskResponse{}, err
	}
	return resp, nil
}
