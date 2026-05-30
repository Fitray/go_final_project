package auth_service

import core_domain "github.com/Fitray/go_final_project/internal/core/domain"

type AuthService struct {
	authRepository AuthRepository
}

type AuthRepository interface {
	SignIn(
		signInRequest core_domain.SignInRequest,
	) (string, error)
}

func NewAuthService(authRepository AuthRepository) AuthService {
	return AuthService{
		authRepository: authRepository,
	}
}
