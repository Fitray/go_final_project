package tasks_handler

import (
	"encoding/json"
	"net/http"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
	core_response "github.com/Fitray/go_final_project/internal/core/response"
)

func (h *TasksHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	wr := core_response.NewResponseHandler(w)
	var task core_domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		wr.JSONResponce(map[string]string{"error": "Invalid JSON format"},
			http.StatusBadRequest)
		return
	}

	resp, err := h.tasksService.AddTask(task)
	if err != nil {
		wr.JSONResponce(map[string]string{"error": err.Error()},
			core_errors.GetStatusCode(err))
		return
	}
	wr.JSONResponce(resp, http.StatusCreated)
}
