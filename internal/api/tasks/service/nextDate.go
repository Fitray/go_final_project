package tasks_service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
)

func (s *TasksService) days(rule []string, now, start time.Time) (string, error) {
	if len(rule) != 2 {
		return "", fmt.Errorf("wrong input format")
	}
	day, err := strconv.Atoi(rule[1])
	if err != nil {
		return "", err
	}
	if day > 400 {
		return "", fmt.Errorf("days amount can't be more than 400")
	}
	result, err := s.tasksRepository.NextDate(
		now, start, core_domain.NextDateParams{
			Day:  day,
			Type: rule[0],
		},
	)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *TasksService) year(rule []string, now, start time.Time) (string, error) {
	if len(rule) != 1 {
		return "", fmt.Errorf("this option doesn't require any values with it")
	}
	result, err := s.tasksRepository.NextDate(
		now, start, core_domain.NextDateParams{
			Type: rule[0],
		},
	)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *TasksService) week(rule []string, now, start time.Time) (string, error) {
	if len(rule) != 2 {
		return "", fmt.Errorf("wrong input format")
	}
	days, err := getWeekDaysFromString(rule[1])
	if err != nil {
		return "", err
	}
	result, err := s.tasksRepository.NextDate(now, start, core_domain.NextDateParams{
		Type: rule[0],
		Days: days,
	})
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *TasksService) month(rule []string, now, start time.Time) (string, error) {
	if len(rule) < 2 || len(rule) > 3 {
		return "", fmt.Errorf("wrong arguements amount")
	}

	days, err := getDaysFromString(rule[1])
	if err != nil {
		return "", err
	}

	months := make([]int, 0)

	if len(rule) == 3 {
		months, err = getMonthsFromString(rule[2])
		if err != nil {
			return "", err
		}
	}

	result, err := s.tasksRepository.NextDate(now, start, core_domain.NextDateParams{
		Type:   rule[0],
		Days:   days,
		Months: months,
	})
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *TasksService) NextDate(
	now_str string, dstart string, repeat string,
) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("repeat param can't be empty: %w", core_errors.ErrBadRequest)
	}
	now, err := time.Parse("20060102", now_str)
	if err != nil {
		return "", fmt.Errorf("%w: %v", err, core_errors.ErrBadRequest)
	}
	start, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", fmt.Errorf("%w: %v", err, core_errors.ErrBadRequest)
	}
	rule := strings.Split(repeat, " ")

	switch {
	case rule[0] == "d":
		return s.days(rule, now, start)
	case rule[0] == "y":
		return s.year(rule, now, start)
	case rule[0] == "w":
		return s.week(rule, now, start)
	case rule[0] == "m":
		return s.month(rule, now, start)
	default:
		return "", fmt.Errorf("invalid type: %w", core_errors.ErrBadRequest)
	}
}
