package scheduler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
)

func getMonthsFromString(seq string) ([]int, error) {
	months_str := strings.Split(seq, ",")
	months := make([]int, len(months_str))
	for i, month_str := range months_str {
		month, err := strconv.Atoi(month_str)
		if err != nil {
			return []int{}, fmt.Errorf("%w: %v", err, core_errors.ErrBadRequest)
		}
		if month > 12 || month < 1 {
			return []int{}, fmt.Errorf("%w: %v", err, core_errors.ErrBadRequest)
		}
		months[i] = month
	}
	return months, nil
}

func getWeekDaysFromString(seq string) ([]int, error) {
	days_str := strings.Split(seq, ",")
	days := make([]int, len(days_str))
	for i, day_str := range days_str {
		day, err := strconv.Atoi(day_str)
		if err != nil {
			return []int{}, fmt.Errorf("%w: %v", err, core_errors.ErrBadRequest)
		}
		if day > 7 || day < 1 {
			return []int{}, fmt.Errorf("%w: %v", err, core_errors.ErrBadRequest)
		}
		days[i] = day
	}
	return days, nil
}

func getDaysFromString(seq string) ([]int, error) {
	days_str := strings.Split(seq, ",")
	days := make([]int, len(days_str))
	for i, day_str := range days_str {
		day, err := strconv.Atoi(day_str)
		if err != nil {
			return []int{}, fmt.Errorf("%w: %v", err, core_errors.ErrBadRequest)
		}
		if day > 31 || day < -2 || day == 0 {
			return []int{}, fmt.Errorf("%w: %v", err, core_errors.ErrBadRequest)
		}
		days[i] = day
	}
	return days, nil
}

func days(rule []string, now, start time.Time) (core_domain.NextDateParams, error) {
	if len(rule) != 2 {
		return core_domain.NextDateParams{}, fmt.Errorf("wrong input format")
	}
	day, err := strconv.Atoi(rule[1])
	if err != nil {
		return core_domain.NextDateParams{}, err
	}
	if day > 400 {
		return core_domain.NextDateParams{}, fmt.Errorf("days amount can't be more than 400")
	}
	return core_domain.NextDateParams{
		Day:  day,
		Type: rule[0],
	}, nil
}

func year(rule []string, now, start time.Time) (core_domain.NextDateParams, error) {
	if len(rule) != 1 {
		return core_domain.NextDateParams{}, fmt.Errorf("this option doesn't require any values with it")
	}
	return core_domain.NextDateParams{
		Type: rule[0],
	}, nil
}

func week(rule []string, now, start time.Time) (core_domain.NextDateParams, error) {
	if len(rule) != 2 {
		return core_domain.NextDateParams{}, fmt.Errorf("wrong input format")
	}
	days, err := getWeekDaysFromString(rule[1])
	if err != nil {
		return core_domain.NextDateParams{}, err
	}
	return core_domain.NextDateParams{
		Type: rule[0],
		Days: days,
	}, nil
}

func month(rule []string, now, start time.Time) (core_domain.NextDateParams, error) {
	if len(rule) < 2 || len(rule) > 3 {
		return core_domain.NextDateParams{}, fmt.Errorf("wrong arguements amount")
	}

	days, err := getDaysFromString(rule[1])
	if err != nil {
		return core_domain.NextDateParams{}, err
	}

	months := make([]int, 0)

	if len(rule) == 3 {
		months, err = getMonthsFromString(rule[2])
		if err != nil {
			return core_domain.NextDateParams{}, err
		}
	}

	return core_domain.NextDateParams{
		Type:   rule[0],
		Days:   days,
		Months: months,
	}, nil
}

func ParseRule(
	now_str string, dstart string, repeat string,
) (core_domain.NextDateParams, error) {
	if repeat == "" {
		return core_domain.NextDateParams{}, fmt.Errorf("repeat param can't be empty: %w", core_errors.ErrBadRequest)
	}
	now, err := time.Parse(TimeFormat, now_str)
	if err != nil {
		return core_domain.NextDateParams{}, fmt.Errorf("%w: %v", err, core_errors.ErrBadRequest)
	}
	start, err := time.Parse(TimeFormat, dstart)
	if err != nil {
		return core_domain.NextDateParams{}, fmt.Errorf("%w: %v", err, core_errors.ErrBadRequest)
	}
	rule := strings.Split(repeat, " ")

	switch {
	case rule[0] == "d":
		return days(rule, now, start)
	case rule[0] == "y":
		return year(rule, now, start)
	case rule[0] == "w":
		return week(rule, now, start)
	case rule[0] == "m":
		return month(rule, now, start)
	default:
		return core_domain.NextDateParams{}, fmt.Errorf("invalid type: %w", core_errors.ErrBadRequest)
	}
}
