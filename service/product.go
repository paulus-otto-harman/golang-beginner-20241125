package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type ProductService struct {
	Product *repository.Product
	Logger  *zap.Logger
}

func InitProductService(repo repository.Repository, log *zap.Logger) *ProductService {
	return &ProductService{Product: repo.Product, Logger: log}
}

func (repo *ProductService) Index(page int, limit int, keyword string, categoryId int, authToken string) (int, int, []domain.Product, error) {
	return repo.Product.Index(page, limit, keyword, categoryId, authToken)
}

func (repo *ProductService) Get(product *domain.ProductDetail, authToken string) error {
	return repo.Product.Get(product, authToken)
}
