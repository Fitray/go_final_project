package core_response

import (
	"encoding/json"
	"net/http"
)

type ResponseHandler struct {
	w http.ResponseWriter
}

func NewResponseHandler(w http.ResponseWriter) ResponseHandler {
	return ResponseHandler{w: w}
}

func (w *ResponseHandler) JSONResponce(resp any, status int) error {
	w.w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.w.WriteHeader(status)
	if err := json.NewEncoder(w.w).Encode(&resp); err != nil {
		return err
	}
	return nil
}
