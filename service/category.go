package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type CategoryService struct {
	Category *repository.Category
	Logger   *zap.Logger
}

func InitCategoryService(repo repository.Repository, log *zap.Logger) *CategoryService {
	return &CategoryService{Category: repo.Category, Logger: log}
}

func (repo *CategoryService) Get() ([]domain.Category, error) {
	return repo.Category.Get()
}
