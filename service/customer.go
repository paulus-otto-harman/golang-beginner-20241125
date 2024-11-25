package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type CustomerService struct {
	Customer *repository.Customer
	Logger   *zap.Logger
}

func InitCustomerService(repo repository.Repository, log *zap.Logger) *CustomerService {
	return &CustomerService{Customer: repo.Customer, Logger: log}
}

func (repo *CustomerService) Register(customer *domain.Customer) error {
	return repo.Customer.Register(customer)
}

func (repo *CustomerService) Profile(authToken string) (*domain.Customer, error) {
	return repo.Customer.Profile(authToken)
}
