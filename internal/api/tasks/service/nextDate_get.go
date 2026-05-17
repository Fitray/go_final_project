package tasks_service

import (
	"fmt"
	"strconv"
	"strings"

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
