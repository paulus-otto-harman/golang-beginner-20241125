package service

import (
	"github.com/stretchr/testify/mock"
	"reflect"
)

type ValidationServiceMock struct {
	mock.Mock
}

func (serviceMock *ValidationServiceMock) IsUnique(tableName string, columnName string, value string) (bool, error) {
	args := serviceMock.Called(tableName, columnName, value)
	if sessionResult := args.Get(0); sessionResult != nil {
		return sessionResult.(bool), args.Error(1)
	}
	return false, args.Error(1)
}

func (serviceMock *ValidationServiceMock) Exists(tableName string, columnName string, value reflect.Value) (bool, error) {
	args := serviceMock.Called(tableName, columnName, value)
	if sessionResult := args.Get(0); sessionResult != nil {
		return sessionResult.(bool), args.Error(1)
	}
	return false, args.Error(1)
}

func (serviceMock *ValidationServiceMock) ExistsForUser(authToken string, addressId int) (bool, error) {
	args := serviceMock.Called(authToken, addressId)
	if sessionResult := args.Get(0); sessionResult != nil {
		return sessionResult.(bool), args.Error(1)
	}
	return false, args.Error(1)
}

func (serviceMock *ValidationServiceMock) NotEmptyCart(authToken string) (bool, error) {
	args := serviceMock.Called(authToken)
	if sessionResult := args.Get(0); sessionResult != nil {
		return sessionResult.(bool), args.Error(1)
	}
	return false, args.Error(1)
}
