package tasks_handler

import (
	"net/http"

	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
	core_response "github.com/Fitray/go_final_project/internal/core/response"
)

func (h *TasksHandler) NextDate(w http.ResponseWriter, r *http.Request) {
	wr := core_response.NewResponseHandler(w)

	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	result, err := h.tasksService.NextDate(now, date, repeat)
	if err != nil {
		wr.JSONResponce(
			map[string]string{"error": err.Error()},
			core_errors.GetStatusCode(err),
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(result)); err != nil {
		wr.JSONResponce(
			map[string]string{"error": err.Error()},
			core_errors.GetStatusCode(err),
		)
		return
	}
}
