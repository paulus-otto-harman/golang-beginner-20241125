package service

import (
	"github.com/stretchr/testify/mock"
	"project/domain"
)

type AuthServiceMock struct {
	mock.Mock
}

func (serviceMock *AuthServiceMock) Login(user domain.User) (*domain.Session, error) {
	args := serviceMock.Called(user)
	if sessionResult := args.Get(0); sessionResult != nil {
		return sessionResult.(*domain.Session), args.Error(1)
	}
	return nil, args.Error(1)
}

func (serviceMock *AuthServiceMock) Logout(token string) error {
	return nil
}
