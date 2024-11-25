package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type BestsellerService struct {
	Bestseller *repository.Bestseller
	Logger     *zap.Logger
}

func InitBestsellerService(repo repository.Repository, log *zap.Logger) *BestsellerService {
	return &BestsellerService{Bestseller: repo.Bestseller, Logger: log}
}

func (repo *BestsellerService) Index(page int, limit int) (int, int, []domain.Product, error) {
	return repo.Bestseller.Index(page, limit)
}
