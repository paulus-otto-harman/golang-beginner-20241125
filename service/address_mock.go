package service

import (
	"github.com/stretchr/testify/mock"
	"project/domain"
)

type AddressServiceMock struct {
	mock.Mock
}

func (serviceMock *AddressServiceMock) Index(authToken string) ([]domain.Address, error) {
	args := serviceMock.Called(authToken)
	if sessionResult := args.Get(0); sessionResult != nil {
		return sessionResult.([]domain.Address), args.Error(1)
	}
	return nil, args.Error(1)
}

func (serviceMock *AddressServiceMock) Create(address *domain.Address, authToken string) error {
	args := serviceMock.Called(address, authToken)
	return args.Error(0)
}

func (serviceMock *AddressServiceMock) Update(address *domain.Address, authToken string) error {
	args := serviceMock.Called(address, authToken)
	if sessionResult := args.Get(0); sessionResult != nil {
		return args.Error(1)
	}
	return args.Error(1)
}

func (serviceMock *AddressServiceMock) SetDefault(addressId int, authToken string) error {
	args := serviceMock.Called(addressId, authToken)
	if sessionResult := args.Get(0); sessionResult != nil {
		return args.Error(1)
	}
	return args.Error(1)
}

func (serviceMock *AddressServiceMock) Delete(addressId int, authToken string) error {
	args := serviceMock.Called(addressId, authToken)
	if sessionResult := args.Get(0); sessionResult != nil {
		return args.Error(1)
	}
	return args.Error(1)
}
