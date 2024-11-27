package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type CustomerServiceInterface interface {
	Register(customer *domain.Customer) error
	Profile(authToken string) (*domain.Customer, error)
}

type customerService struct {
	Customer *repository.Customer
	Logger   *zap.Logger
}

func InitCustomerService(repo repository.Repository, log *zap.Logger) CustomerServiceInterface {
	return &customerService{Customer: repo.Customer, Logger: log}
}

func (repo *customerService) Register(customer *domain.Customer) error {
	return repo.Customer.Register(customer)
}

func (repo *customerService) Profile(authToken string) (*domain.Customer, error) {
	return repo.Customer.Profile(authToken)
}
