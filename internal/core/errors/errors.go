package core_errors

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Message string
	Status  int
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

var (
	ErrBadRequest = &ErrorResponse{
		Message: "bad request",
		Status:  http.StatusBadRequest}
	ErrNotFound = &ErrorResponse{
		Message: "not found",
		Status:  http.StatusNotFound}
	ErrAuthFailed = &ErrorResponse{
		Message: "authentication failed",
		Status:  http.StatusUnauthorized}
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	var res *ErrorResponse
	if ok := errors.As(err, &res); ok {
		return res.Status
	}
	return http.StatusInternalServerError
}
