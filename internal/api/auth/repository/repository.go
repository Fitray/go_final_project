package auth_repository

import (
	core_auth "github.com/Fitray/go_final_project/internal/core/auth"
)

type AuthRepository struct {
	Auth core_auth.Auth
}

func NewAuthRepository(
	auth core_auth.Auth,
) AuthRepository {
	return AuthRepository{
		Auth: auth,
	}
}
