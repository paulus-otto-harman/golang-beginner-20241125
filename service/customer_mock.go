package service

import (
	"github.com/stretchr/testify/mock"
	"project/domain"
)

type CustomerServiceMock struct {
	mock.Mock
}

func (serviceMock *CustomerServiceMock) Register(customer *domain.Customer) error {
	args := serviceMock.Called(customer)
	if sessionResult := args.Get(0); sessionResult != nil {
		return args.Error(0)
	}
	return nil
}

func (serviceMock *CustomerServiceMock) Profile(authToken string) (*domain.Customer, error) {
	args := serviceMock.Called(authToken)
	if sessionResult := args.Get(0); sessionResult != nil {
		return sessionResult.(*domain.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}
