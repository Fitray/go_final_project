package tasks_repository

import (
	"context"
	"fmt"
	"time"

	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
)

func (r *TasksRepository) CompleteTask(id string, nextDate string) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	query := `
	UPDATE scheduler SET date = $1 WHERE id = $2
	`

	res, err := r.DB.ExecContext(ctx, query, nextDate, id)
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
