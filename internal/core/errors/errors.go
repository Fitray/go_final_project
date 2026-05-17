package core_errors

import (
	"errors"
	"net/http"
)

var (
	ErrBadRequest = errors.New("bad request")
)

func GetStatusCode(err error) int {
	switch {
	case errors.Is(err, ErrBadRequest):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
