package tasks_handler

import (
	"net/http"

	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
	core_response "github.com/Fitray/go_final_project/internal/core/response"
)

func (h *TasksHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	wr := core_response.NewResponseHandler(w)
	search := r.FormValue("search")

	tasksResponse, err := h.tasksService.GetTasks(search, 50)
	if err != nil {
		wr.JSONResponce(map[string]string{"error": err.Error()},
			core_errors.GetStatusCode(err))
		return
	}
	wr.JSONResponce(tasksResponse, http.StatusOK)
}

func (h *TasksHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	wr := core_response.NewResponseHandler(w)

	id := r.URL.Query().Get("id")
	if id == "" {
		wr.JSONResponce(map[string]string{"error": "ID is required"},
			http.StatusBadRequest)
		return
	}

	tasksResponse, err := h.tasksService.GetTask(id)
	if err != nil {
		wr.JSONResponce(map[string]string{"error": err.Error()},
			core_errors.GetStatusCode(err))
		return
	}
	wr.JSONResponce(tasksResponse, http.StatusOK)
}
