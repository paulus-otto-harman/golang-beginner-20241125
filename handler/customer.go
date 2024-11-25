package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"project/domain"
	"project/service"
	"project/util"
)

type CustomerHandler struct {
	CustomerService *service.CustomerService
	Logger          *zap.Logger
	Validator       util.Validation
}

func InitCustomerHandler(service service.Service, log *zap.Logger, validator util.Validation) CustomerHandler {
	return CustomerHandler{CustomerService: service.Customer, Logger: log, Validator: validator}
}

func (handle *CustomerHandler) Registration(w http.ResponseWriter, r *http.Request) {
	customer := domain.Customer{}
	json.NewDecoder(r.Body).Decode(&customer)

	body, _ := json.Marshal(customer)
	handle.Logger.Debug("validate input", zap.Any("data", json.RawMessage(body)))

	if err := handle.Validator.ValidateStruct(customer); err != nil {
		util.Response(w).ValidationFail(err)
		return
	}

	err := handle.CustomerService.Register(&customer)
	if err != nil {
		util.Response(w).Json(http.StatusBadRequest, "Fail to register customer")
		return
	}

	util.Response(w).Json(http.StatusCreated, "Customer registered successfully.", customer)
}

func (handle *CustomerHandler) Profile(w http.ResponseWriter, r *http.Request) {
	customer, err := handle.CustomerService.Profile(r.Header.Get("token"))
	if err != nil {
		util.Response(w).Json(http.StatusBadRequest, "Fail to retrieve profile")
		return
	}
	util.Response(w).Json(http.StatusCreated, "Customer Profile retrieved successfully.", customer)
}
