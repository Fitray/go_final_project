package core_auth

import (
	"fmt"

	core_errors "github.com/Fitray/go_final_project/internal/core/errors"
	"github.com/golang-jwt/jwt"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Password  string `envconfig:"TODO_PASSWORD" required:"true"`
	JWTSecret string `envconfig:"TODO_JWT_SECRET" required:"true"`
}

func NewAuth() (Auth, error) {
	var authConfig Auth
	if err := envconfig.Process("", &authConfig); err != nil {
		return Auth{}, err
	}
	return authConfig, nil
}

func (a *Auth) CreateToken(hash string) (string, error) {
	claims := jwt.MapClaims{
		"hash": hash,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", core_errors.ErrAuthFailed)
	}

	return tokenString, nil
}

func (a *Auth) ValidateToken(tokenString string) (bool, error) {
	jwtToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(a.JWTSecret), nil
	})

	if err != nil {

		return false, fmt.Errorf("failed to parse token: %w", core_errors.ErrAuthFailed)
	}
	if !jwtToken.Valid {
		return false, fmt.Errorf("invalid token: %w", core_errors.ErrAuthFailed)
	}

	res, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("invalid token claims: %w", core_errors.ErrAuthFailed)
	}

	hash, ok := res["hash"].(string)
	if !ok {
		return false, fmt.Errorf("invalid token claims: %w", core_errors.ErrAuthFailed)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(a.Password)); err != nil {
		return false, fmt.Errorf("invalid token hash: %w", core_errors.ErrAuthFailed)
	}

	return true, nil
}
