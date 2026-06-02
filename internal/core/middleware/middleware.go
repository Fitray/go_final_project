package core_middleware

import (
	"net/http"

	core_auth "github.com/Fitray/go_final_project/internal/core/auth"
	core_response "github.com/Fitray/go_final_project/internal/core/response"
	"github.com/go-chi/chi/v5"
)

func GetAuthChain(auth core_auth.Auth) chi.Middlewares {
	return chi.Middlewares{
		Auth(auth),
	}
}

func Auth(auth core_auth.Auth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := core_response.NewResponseHandler(w)

			if len(auth.Password) > 0 {
				var jwt string
				cookie, err := r.Cookie("token")
				if err == nil {
					jwt = cookie.Value
				}

				valid, err := auth.ValidateToken(jwt)
				if err != nil {
					rw.JSONResponce(
						map[string]string{"error": err.Error()},
						http.StatusUnauthorized,
					)
					return
				}

				if !valid {
					rw.JSONResponce(
						map[string]string{"error": "Authentification required"},
						http.StatusUnauthorized,
					)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
