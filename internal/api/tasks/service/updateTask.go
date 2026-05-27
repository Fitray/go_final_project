package tasks_service

import core_domain "github.com/Fitray/go_final_project/internal/core/domain"

func (s *TasksService) UpdateTask(
	taskRequest core_domain.Task,
) error {
	taskRequest, err := s.CheckTask(taskRequest)
	if err != nil {
		return err
	}

	err = s.tasksRepository.UpdateTask(taskRequest)
	if err != nil {
		return err
	}

	return nil
}
