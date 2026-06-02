package tasks_repository

import (
	"context"
	"fmt"

	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
)

func (r *TasksRepository) DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		r.DB.Timeout,
	)
	defer cancel()

	query := `
	DELETE FROM scheduler WHERE id=$1
	`
	res, err := r.DB.Database.ExecContext(ctx, query, id)
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
