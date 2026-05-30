package auth_handler

import (
	"encoding/json"
	"log"
	"net/http"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
	core_response "github.com/Fitray/go_final_project/internal/core/response"
)

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	wr := core_response.NewResponseHandler(w)

	var signInRequest core_domain.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&signInRequest); err != nil {
		wr.JSONResponce(map[string]string{"error": "Invalid JSON format"},
			http.StatusBadRequest)
		return
	}

	token, err := h.authService.SignIn(signInRequest)
	if err != nil {
		wr.JSONResponce(map[string]string{"error": err.Error()},
			core_errors.GetStatusCode(err))
		return
	}

	log.Printf("token: %s", token)

	wr.JSONResponce(map[string]string{"token": token}, http.StatusOK)
}
