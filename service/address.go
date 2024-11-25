package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type AddressServiceInterface interface {
	Index(authToken string) ([]domain.Address, error)
	Create(address *domain.Address, authToken string) error
	Update(address *domain.Address, authToken string) error
	SetDefault(addressId int, authToken string) error
	Delete(addressId int, authToken string) error
}

type AddressService struct {
	Address *repository.Address
	Logger  *zap.Logger
}

func InitAddressService(repo repository.Repository, log *zap.Logger) AddressServiceInterface {
	return &AddressService{Address: repo.Address, Logger: log}
}

func (repo *AddressService) Index(authToken string) ([]domain.Address, error) {
	addresses, err := repo.Address.Index(authToken)

	if err != nil {
		repo.Logger.Error("get all address error", zap.Error(err))
		return nil, err
	}

	return addresses, nil
}

func (repo *AddressService) Create(address *domain.Address, authToken string) error {
	return repo.Address.Store(address, authToken)
}

func (repo *AddressService) Update(address *domain.Address, authToken string) error {
	return repo.Address.Update(address, authToken)
}

func (repo *AddressService) SetDefault(addressId int, authToken string) error {
	return repo.Address.SetDefault(addressId, authToken)
}

func (repo *AddressService) Delete(addressId int, authToken string) error {
	return repo.Address.Destroy(addressId, authToken)
}
