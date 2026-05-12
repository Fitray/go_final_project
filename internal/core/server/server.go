package core_server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	webDir = "web"
)

type HTTPServer struct {
	Mux    *chi.Mux
	Config Config
}

func NewHTTPServer(config Config) HTTPServer {
	mux := chi.NewMux()
	return HTTPServer{
		Mux:    mux,
		Config: config,
	}
}

func (s *HTTPServer) Run() error {
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", s.Config.Addr),
		Handler: s.Mux,
	}

	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("HTTP server error: %w", err)
		}
	}
	return nil
}

func (s *HTTPServer) RegisterRoutes() {
	s.Mux.Handle("/*", http.FileServer(http.Dir(webDir)))
}
