package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type OrderService struct {
	Order  *repository.Order
	Logger *zap.Logger
}

func InitOrderService(repo repository.Repository, log *zap.Logger) *OrderService {
	return &OrderService{Order: repo.Order, Logger: log}
}

func (repo *OrderService) Create(order *domain.Order) error {
	return repo.Order.Create(order)
}
