package tasks_service

import (
	"fmt"
	"time"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
)

func (s *TasksService) CheckTask(taskRequest core_domain.Task) error {
	if taskRequest.Date == "" {
		taskRequest.Date = time.Now().Format("20060102")
	}

	t, err := time.Parse("20060102", taskRequest.Date)
	if err != nil {
		return fmt.Errorf("%w:%w", err, core_errors.ErrBadRequest)
	}

	if taskRequest.Title == "" {
		return fmt.Errorf(
			"invalid title: %w",
			core_errors.ErrBadRequest)
	}

	now := time.Now().UTC()
	t = t.UTC()

	var nextDate string
	if taskRequest.Repeat != "" {
		nextDate, err = s.NextDate(now.Format("20060102"), taskRequest.Date, taskRequest.Repeat)
		if err != nil {
			return fmt.Errorf("%w:%w", err, core_errors.ErrBadRequest)
		}
	}

	if now.After(t) {
		if nextDate == "" {
			taskRequest.Date = now.Format("20060102")
		} else {
			taskRequest.Date = nextDate
		}
	}
	return nil
}

func (s *TasksService) AddTask(
	taskRequest core_domain.Task,
) (core_domain.NewTaskResponse, error) {
	if err := s.CheckTask(taskRequest); err != nil {
		return core_domain.NewTaskResponse{}, err
	}

	resp, err := s.tasksRepository.AddTask(taskRequest)
	if err != nil {
		return core_domain.NewTaskResponse{}, err
	}
	return resp, nil
}
