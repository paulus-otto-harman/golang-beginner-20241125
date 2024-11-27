package service

import (
	"github.com/stretchr/testify/mock"
	"project/domain"
)

type CartServiceMock struct {
	mock.Mock
}

func (serviceMock *CartServiceMock) Get(authToken string) (*domain.Cart, error) {
	args := serviceMock.Called(authToken)
	if sessionResult := args.Get(0); sessionResult != nil {
		return sessionResult.(*domain.Cart), args.Error(1)
	}
	return nil, args.Error(1)
}

func (serviceMock *CartServiceMock) Store(cartItem domain.CartItem, authToken string) error {
	args := serviceMock.Called(cartItem, authToken)
	if sessionResult := args.Get(0); sessionResult != nil {
		return args.Error(0)
	}
	return nil
}

func (serviceMock *CartServiceMock) Update(cartItem domain.CartItem, authToken string) error {
	args := serviceMock.Called(cartItem, authToken)
	if sessionResult := args.Get(0); sessionResult != nil {
		return args.Error(0)
	}
	return nil
}

func (serviceMock *CartServiceMock) Delete(productId int, authToken string) error {
	args := serviceMock.Called(productId, authToken)
	if sessionResult := args.Get(0); sessionResult != nil {
		return args.Error(1)
	}
	return nil
}
