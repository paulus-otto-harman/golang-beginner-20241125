package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type BannerService struct {
	Banner *repository.Banner
	Logger *zap.Logger
}

func InitBannerService(repo repository.Repository, log *zap.Logger) *BannerService {
	return &BannerService{Banner: repo.Banner, Logger: log}
}

func (repo *BannerService) Get() ([]domain.Banner, error) {
	return repo.Banner.Get()
}
