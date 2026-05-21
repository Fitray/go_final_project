package tasks_repository

import (
	"context"
	"time"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
)

func (t *TasksRepository) getTasks_Rows(
	query string, args ...any,
) (core_domain.GetTasksResponse, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	rows, err := t.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return core_domain.GetTasksResponse{
			Tasks: []core_domain.Task{},
		}, err
	}
	defer rows.Close()

	tasksResponse := core_domain.GetTasksResponse{
		Tasks: []core_domain.Task{},
	}
	for rows.Next() {
		var task core_domain.Task
		err := rows.Scan(
			&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat,
		)
		if err != nil {
			return core_domain.GetTasksResponse{
				Tasks: []core_domain.Task{},
			}, err
		}
		tasksResponse.Tasks = append(tasksResponse.Tasks, task)
	}

	return tasksResponse, nil
}

func (t *TasksRepository) getTasks_Row(
	query string, args ...any,
) (core_domain.Task, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	row := t.DB.QueryRowContext(ctx, query, args...)

	var tasksResponse core_domain.Task
	err := row.Scan(
		&tasksResponse.Id, &tasksResponse.Date, &tasksResponse.Title, &tasksResponse.Comment, &tasksResponse.Repeat,
	)
	if err != nil {
		return core_domain.Task{}, err
	}

	return tasksResponse, nil
}

func (t *TasksRepository) GetTasks_NoSearch(
	limit int,
) (core_domain.GetTasksResponse, error) {
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	ORDER BY date
	LIMIT $1
	`
	return t.getTasks_Rows(query, limit)
}

func (t *TasksRepository) GetTasks_TextSearch(
	search string, limit int,
) (core_domain.GetTasksResponse, error) {
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	WHERE title LIKE $1 OR comment LIKE $1
	ORDER BY date
	LIMIT $2
	`
	return t.getTasks_Rows(query, "%"+search+"%", limit)
}

func (t *TasksRepository) GetTasks_DateSearch(
	search string, limit int,
) (core_domain.GetTasksResponse, error) {
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	WHERE date = $1
	LIMIT $2
	`
	return t.getTasks_Rows(query, search, limit)
}

func (t *TasksRepository) GetTasks_FromID(
	id string,
) (core_domain.Task, error) {
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	WHERE id = $1
	`
	return t.getTasks_Row(query, id)
}
