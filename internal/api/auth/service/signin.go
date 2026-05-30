package auth_service

import core_domain "github.com/Fitray/go_final_project/internal/core/domain"

func (s *AuthService) SignIn(
	signInRequest core_domain.SignInRequest,
) (string, error) {
	token, err := s.authRepository.SignIn(signInRequest)
	if err != nil {
		return "", err
	}
	return token, nil
}
