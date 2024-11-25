package service

import (
	"go.uber.org/zap"
	"project/repository"
)

type SessionService struct {
	Session *repository.Session
	Logger  *zap.Logger
}

func InitSessionService(repo repository.Repository, log *zap.Logger) *SessionService {
	return &SessionService{Session: repo.Session, Logger: log}
}

func (repo *SessionService) Validate(sessionId string) error {
	return repo.Session.Validate(sessionId)
}
