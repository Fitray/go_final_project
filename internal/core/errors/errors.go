package core_errors

import (
	"errors"
	"net/http"
)

var (
	ErrBadRequest = errors.New("bad request")
	ErrNotFound   = errors.New("not found")
	ErrAuthFailed = errors.New("authentication failed")
)

func GetStatusCode(err error) int {
	switch {
	case errors.Is(err, ErrBadRequest):
		return http.StatusBadRequest
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrAuthFailed):
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
