package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type WeeklyService struct {
	Weekly *repository.Weekly
	Logger *zap.Logger
}

func InitWeeklyService(repo repository.Repository, log *zap.Logger) *WeeklyService {
	return &WeeklyService{Weekly: repo.Weekly, Logger: log}
}

func (repo *WeeklyService) Get() ([]domain.Weekly, error) {
	return repo.Weekly.Get()
}
