package tasks_service

import (
	"fmt"
	"time"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
	"github.com/Fitray/go_final_project/internal/scheduler"
)

func (s *TasksService) GetTasks(
	search string, limit int,
) (core_domain.GetTasksResponse, error) {
	if limit <= 0 {
		return core_domain.GetTasksResponse{
				Tasks: []core_domain.Task{},
			},
			fmt.Errorf("limit should be 1 or more: %w", core_errors.ErrBadRequest)
	}
	if search != "" {
		if time, err := time.Parse(scheduler.DotsTimeFormat, search); err == nil {
			search = time.Format(scheduler.TimeFormat)
			tasksReponse, err := s.tasksRepository.GetTasksByDate(search, limit)
			if err != nil {
				return core_domain.GetTasksResponse{
					Tasks: []core_domain.Task{},
				}, err
			}
			return tasksReponse, nil
		} else {
			tasksReponse, err := s.tasksRepository.GetTasksByText(search, limit)
			if err != nil {
				return core_domain.GetTasksResponse{
					Tasks: []core_domain.Task{},
				}, err
			}
			return tasksReponse, nil
		}
	}

	tasksReponse, err := s.tasksRepository.GetTasks(limit)
	if err != nil {
		return core_domain.GetTasksResponse{
			Tasks: []core_domain.Task{},
		}, err
	}
	return tasksReponse, nil
}

func (s *TasksService) GetTask(
	id string,
) (core_domain.Task, error) {
	taskResponse, err := s.tasksRepository.GetTaskByID(id)
	if err != nil {
		return core_domain.Task{}, err
	}
	return taskResponse, nil
}
