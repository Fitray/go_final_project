package auth_handler

import (
	"net/http"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_server "github.com/Fitray/go_final_project/internal/core/server"
)

type AuthHandler struct {
	authService AuthService
}

type AuthService interface {
	SignIn(
		signInRequest core_domain.SignInRequest,
	) (string, error)
}

func NewAuthHandler(authService AuthService) AuthHandler {
	return AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Routes(
	routes []core_server.Route,
) []core_server.Route {
	for _, route := range []core_server.Route{
		{
			Method:  http.MethodPost,
			Pattern: "/api/signin",
			Handler: h.SignIn,
		},
	} {
		routes = append(routes, route)
	}
	return routes
}
