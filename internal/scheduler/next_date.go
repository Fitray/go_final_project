package scheduler

import (
	"fmt"
	"slices"
	"time"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
)

const (
	TimeFormat     = "20060102"
	DotsTimeFormat = "02.01.2006"
)

func isValidDay(date time.Time, days []int) bool {
	lastDay := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0,
		time.UTC).Day()
	for _, day := range days {
		switch day {
		case -1:
			if date.Day() == lastDay {
				return true
			}
		case -2:
			if date.Day() == lastDay-1 {
				return true
			}
		default:
			if date.Day() == day {
				return true
			}
		}
	}
	return false
}

func nextDateFromDays(
	now, start time.Time, days int,
) (string, error) {
	for {
		start = start.AddDate(0, 0, days)
		if start.After(now) {
			return start.Format(TimeFormat), nil
		}
	}
}

func nextDateFromYear(
	now, start time.Time, year int,
) (string, error) {
	for {
		start = start.AddDate(year, 0, 0)
		if start.After(now) {
			return start.Format(TimeFormat), nil
		}
	}
}

func nextDateFromMonths(
	now, start time.Time, months, days []int,
) (string, error) {
	for {
		start = start.AddDate(0, 0, 1)
		if start.After(now) {
			if isValidDay(start, days) {
				if len(months) > 0 && !slices.Contains(months, int(start.Month())) {
					continue
				}
				return start.Format(TimeFormat), nil
			}
		}
	}
}

func nextDateFromWeeks(
	now, start time.Time, days []int,
) (string, error) {
	for {
		start = start.AddDate(0, 0, 1)
		if start.After(now) {
			dayWeek := int(start.Weekday())
			if dayWeek == 0 {
				dayWeek = 7
			}
			if slices.Contains(days, dayWeek) {
				return start.Format(TimeFormat), nil
			}
		}
	}
}

func NextDate(
	now,
	start time.Time,
	params core_domain.NextDateParams,
) (string, error) {
	switch params.Type {
	case "d":
		return nextDateFromDays(now, start, params.Day)
	case "y":
		return nextDateFromYear(now, start, 1)
	case "m":
		return nextDateFromMonths(now, start, params.Months, params.Days)
	case "w":
		return nextDateFromWeeks(now, start, params.Days)
	default:
		return "", fmt.Errorf("invalid format: %w", core_errors.ErrBadRequest)
	}
}
