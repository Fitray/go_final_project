package core_middleware

import (
	"net/http"

	core_auth "github.com/Fitray/go_final_project/internal/core/auth"
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
			if len(auth.Password) > 0 {
				var jwt string
				cookie, err := r.Cookie("token")
				if err == nil {
					jwt = cookie.Value
				}

				valid, err := auth.ValidateToken(jwt)
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}

				if !valid {
					http.Error(w, "Authentification required", http.StatusUnauthorized)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
