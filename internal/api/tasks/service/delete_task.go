package tasks_service

import (
	"fmt"

	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
)

func (s *TasksService) DeleteTask(id string) error {
	if id == "" {
		return fmt.Errorf("invalid id: %w", core_errors.ErrBadRequest)
	}

	if err := s.tasksRepository.DeleteTask(id); err != nil {
		return err
	}
	return nil
}
