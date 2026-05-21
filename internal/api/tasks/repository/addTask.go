package tasks_repository

import (
	"context"
	"time"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
)

func (r *TasksRepository) AddTask(
	taskRequest core_domain.Task,
) (core_domain.NewTaskResponse, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	query := `
	INSERT INTO scheduler (date, title, comment, repeat)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`
	var id int
	err := r.DB.QueryRowContext(
		ctx,
		query,
		taskRequest.Date,
		taskRequest.Title,
		taskRequest.Comment,
		taskRequest.Repeat,
	).Scan(&id)
	if err != nil {
		return core_domain.NewTaskResponse{}, err
	}
	return core_domain.NewTaskResponse{Id: id}, nil
}
