package service

import (
	"go.uber.org/zap"
	"project/repository"
	"reflect"
	"strconv"
)

type ValidationServiceInterface interface {
	IsUnique(tableName string, columnName string, value string) (bool, error)
	Exists(tableName string, columnName string, value reflect.Value) (bool, error)
	ExistsForUser(authToken string, addressId int) (bool, error)
	NotEmptyCart(authToken string) (bool, error)
}

type validationService struct {
	Validation *repository.Validation
	Logger     *zap.Logger
}

func InitValidationService(repo repository.Repository, log *zap.Logger) ValidationServiceInterface {
	return &validationService{Validation: repo.Validation, Logger: log}
}

func (repo *validationService) IsUnique(tableName string, columnName string, value string) (bool, error) {
	repo.Logger.Debug("validate :: unique", zap.String("tableName", tableName), zap.String("columnName", columnName), zap.String("value", value))
	return repo.Validation.IsUnique(tableName, columnName, value)
}

func (repo *validationService) Exists(tableName string, columnName string, value reflect.Value) (bool, error) {
	v := func(v reflect.Value) string {
		if value.Kind() == reflect.Int {
			return strconv.Itoa(int(value.Int()))
		}
		return v.String()
	}(value)

	repo.Logger.Debug("validate :: exists", zap.String("tableName", tableName), zap.String("columnName", columnName), zap.String("value", v))
	return repo.Validation.Exists(tableName, columnName, v)
}

func (repo *validationService) ExistsForUser(authToken string, addressId int) (bool, error) {
	repo.Logger.Debug("validate :: customer address must exists (existsForUser)", zap.String("authToken", authToken))
	return repo.Validation.ExistsForUser(authToken, addressId)
}

func (repo *validationService) NotEmptyCart(authToken string) (bool, error) {
	repo.Logger.Debug("validate :: notEmptyCart", zap.String("authToken", authToken))
	return repo.Validation.NotEmptyCart(authToken)
}
