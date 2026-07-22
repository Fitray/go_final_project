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

type Route struct {
	Method      string
	Handler     http.HandlerFunc
	Pattern     string
	Middlewares chi.Middlewares
}

type HTTPServer struct {
	Mux    *chi.Mux
	Config Config
}

func NewHTTPServer(config Config) HTTPServer {
	r := chi.NewRouter()
	return HTTPServer{
		Mux:    r,
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

func (s *HTTPServer) Init(routes []Route) {
	s.Mux.Handle("/*", http.FileServer(http.Dir(webDir)))
	for _, route := range routes {
		s.Mux.Group(func(r chi.Router) {
			if len(route.Middlewares) > 0 {
				r.Use(route.Middlewares...)
			}
			r.MethodFunc(route.Method, route.Pattern, route.Handler)
		})
	}
}
