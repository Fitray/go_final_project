package tasks_service

import (
	"fmt"
	"time"

	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
)

func (s *TasksService) CompleteTask(id string) error {
	if id == "" {
		return fmt.Errorf("invalid id: %w", core_errors.ErrBadRequest)
	}

	task, err := s.GetTask(id)
	if err != nil {
		return err
	}

	now_str := time.Now().UTC().Format("20060102")
	dstart := task.Date
	if dstart == "" {
		dstart = now_str
	}

	if task.Repeat == "" {
		if err := s.tasksRepository.DeleteTask(id); err != nil {
			return err
		}
	} else {
		nextDate, err := s.NextDate(now_str, dstart, task.Repeat)
		if err != nil {
			return err
		}
		err = s.tasksRepository.CompleteTask(id, nextDate)
		if err != nil {
			return err
		}
	}

	return nil
}
