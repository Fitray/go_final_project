package tasks_service

import (
	"fmt"
	"time"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
)

func (t *TasksService) GetTasks(
	search string, limit int,
) (core_domain.GetTasksResponse, error) {
	if limit <= 0 {
		return core_domain.GetTasksResponse{
				Tasks: []core_domain.Task{},
			},
			fmt.Errorf("limit should be 1 or more: %w", core_errors.ErrBadRequest)
	}
	if search != "" {
		if time, err := time.Parse("02.01.2006", search); err == nil {
			search = time.Format("20060102")
			tasksReponse, err := t.tasksRepository.GetTasks_DateSearch(search, limit)
			if err != nil {
				return core_domain.GetTasksResponse{
					Tasks: []core_domain.Task{},
				}, err
			}
			return tasksReponse, nil
		} else {
			tasksReponse, err := t.tasksRepository.GetTasks_TextSearch(search, limit)
			if err != nil {
				return core_domain.GetTasksResponse{
					Tasks: []core_domain.Task{},
				}, err
			}
			return tasksReponse, nil
		}
	}

	tasksReponse, err := t.tasksRepository.GetTasks_NoSearch(limit)
	if err != nil {
		return core_domain.GetTasksResponse{
			Tasks: []core_domain.Task{},
		}, err
	}
	return tasksReponse, nil
}

func (t *TasksService) GetTask(
	id string,
) (core_domain.Task, error) {
	taskResponse, err := t.tasksRepository.GetTasks_FromID(id)
	if err != nil {
		return core_domain.Task{}, err
	}
	return taskResponse, nil
}
