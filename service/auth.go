package service

import (
	"go.uber.org/zap"
	"project/config"
	"project/domain"
	"project/repository"
)

type AuthServiceInterface interface {
	Login(user domain.User) (*domain.Session, error)
	Logout(token string) error
}

type AuthService struct {
	Auth            *repository.Auth
	Logger          *zap.Logger
	SessionLifetime int
}

func InitAuthService(repo repository.Repository, log *zap.Logger, config config.AppConfig) AuthServiceInterface {
	sessionLifetime := config.SessionLifetime
	if config.SessionLifetime == 0 {
		sessionLifetime = 2147483647
	}
	return &AuthService{Auth: repo.Auth, Logger: log, SessionLifetime: sessionLifetime}
}

func (repo *AuthService) Login(user domain.User) (*domain.Session, error) {
	return repo.Auth.Authenticate(user, repo.SessionLifetime)
}

func (repo *AuthService) Logout(token string) error {
	return repo.Auth.Logout(token)
}
