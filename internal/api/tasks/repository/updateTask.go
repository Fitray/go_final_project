package tasks_repository

import (
	"context"
	"fmt"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
)

func (r *TasksRepository) UpdateTask(
	taskRequest core_domain.Task,
) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		r.DB.Timeout,
	)
	defer cancel()

	query := `
	UPDATE scheduler
	SET date = $1, title = $2, comment = $3, repeat = $4
	WHERE id = $5
	`
	res, err := r.DB.DB.ExecContext(ctx, query,
		taskRequest.Date, taskRequest.Title, taskRequest.Comment,
		taskRequest.Repeat, taskRequest.Id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task: %w`,
			core_errors.ErrBadRequest)
	}
	return nil
}
