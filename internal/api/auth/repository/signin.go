package auth_repository

import (
	"fmt"

	core_domain "github.com/Fitray/go_final_project/internal/core/domain"
	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
	"golang.org/x/crypto/bcrypt"
)

func (r *AuthRepository) SignIn(
	signInRequest core_domain.SignInRequest,
) (string, error) {
	if r.Auth.Password != signInRequest.Password {
		return "", fmt.Errorf("wrong password: %w", core_errors.ErrAuthFailed)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(r.Auth.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash: %w", core_errors.ErrAuthFailed)
	}

	tokenString, err := r.Auth.CreateToken(string(hash))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
