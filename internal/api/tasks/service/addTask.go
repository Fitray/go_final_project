package tasks_service

import (
	"fmt"
	"time"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
)

func (s *TasksService) AddTask(
	taskRequest core_domain.TaskRequest,
) (core_domain.TaskResponce, error) {
	if taskRequest.Date == "" {
		taskRequest.Date = time.Now().Format("20060102")
	}

	t, err := time.Parse("20060102", taskRequest.Date)
	if err != nil {
		return core_domain.TaskResponce{},
			fmt.Errorf("%w:%w", err, core_errors.ErrBadRequest)
	}

	if taskRequest.Title == "" {
		return core_domain.TaskResponce{}, fmt.Errorf(
			"invalid title: %w",
			core_errors.ErrBadRequest)
	}

	now := time.Now().UTC()
	t = t.UTC()

	var nextDate string
	if taskRequest.Repeat != "" {
		nextDate, err = s.NextDate(now.Format("20060102"), taskRequest.Date, taskRequest.Repeat)
		if err != nil {
			return core_domain.TaskResponce{}, err
		}
	}

	if now.After(t) {
		if nextDate == "" {
			taskRequest.Date = now.Format("20060102")
		} else {
			taskRequest.Date = nextDate
		}
	}

	resp, err := s.tasksRepository.AddTask(taskRequest)
	if err != nil {
		return core_domain.TaskResponce{}, err
	}
	return resp, nil
}
