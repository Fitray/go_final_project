package tasks_handler

import (
	"net/http"

	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
	core_response "github.com/Fitray/go_final_project/internal/core/response"
)

func (h *TasksHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {
	wr := core_response.NewResponseHandler(w)
	id := r.FormValue("id")

	err := h.tasksService.CompleteTask(id)
	if err != nil {
		wr.JSONResponce(map[string]string{"error": err.Error()},
			core_errors.GetStatusCode(err))
		return
	}

	wr.JSONResponce(map[string]string{},
		http.StatusOK)
}
