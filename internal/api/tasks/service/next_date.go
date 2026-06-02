package tasks_service

import (
	"fmt"
	"time"

	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
	"github.com/Fitray/go_final_project/internal/scheduler"
)

func (s *TasksService) NextDate(
	nowStr,
	startStr,
	repeat string,
) (string, error) {
	if repeat == "" {
		return "",
			fmt.Errorf("repeat param can't be empty: %w", core_errors.ErrBadRequest)
	}

	now, err := time.Parse(scheduler.TimeFormat, nowStr)
	if err != nil {
		return "", fmt.Errorf("%w: %v", err, core_errors.ErrBadRequest)
	}

	start, err := time.Parse(scheduler.TimeFormat, startStr)
	if err != nil {
		return "", fmt.Errorf("%w: %v", err, core_errors.ErrBadRequest)
	}

	params, err := scheduler.ParseRule(nowStr, startStr, repeat)
	if err != nil {
		return "", err
	}

	result, err := scheduler.NextDate(now, start, params)
	if err != nil {
		return "", err
	}

	return result, nil
}
