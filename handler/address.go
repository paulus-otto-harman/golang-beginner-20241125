package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"project/domain"
	"project/service"
	"project/util"
	"strconv"
)

type AddressHandler struct {
	AddressService service.AddressServiceInterface
	Logger         *zap.Logger
	Validator      util.Validation
}

func InitAddressHandler(service service.Service, log *zap.Logger, validator util.Validation) AddressHandler {
	return AddressHandler{AddressService: service.Address, Logger: log, Validator: validator}
}

func (handle AddressHandler) All(w http.ResponseWriter, r *http.Request) {
	addresses, err := handle.AddressService.Index(r.Header.Get("token"))
	if err != nil {
		util.Response(w).Json(http.StatusInternalServerError, "Fail to retrieve customer addresses", addresses)
		return
	}
	util.Response(w).Json(http.StatusOK, "Customer addresses retrieved successfully", addresses)
}

func (handle AddressHandler) Create(w http.ResponseWriter, r *http.Request) {
	address := domain.Address{}

	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		handle.Logger.Error("invalid json input", zap.Error(err))
		util.Response(w).Json(http.StatusUnprocessableEntity, "invalid input")
		return
	}

	if validationError := handle.Validator.ValidateStruct(address); validationError != nil {
		util.Response(w).ValidationFail(validationError)
		return
	}

	if err := handle.AddressService.Create(&address, r.Header.Get("token")); err != nil {
		handle.Logger.Error("fail to create address", zap.Error(err))
		util.Response(w).Json(http.StatusInternalServerError, "server error", err)
		return
	}
	util.Response(w).Json(http.StatusCreated, "address successfully created", address)
}

func (handle AddressHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		util.Response(w).Json(http.StatusBadRequest, "invalid address id")
		return
	}

	address := domain.Address{Id: id}
	if err = json.NewDecoder(r.Body).Decode(&address); err != nil {
		handle.Logger.Error("invalid json input", zap.Error(err))
		util.Response(w).Json(http.StatusUnprocessableEntity, "invalid input")
		return
	}

	if validationError := handle.Validator.ValidateStruct(address); validationError != nil {
		util.Response(w).Json(http.StatusUnprocessableEntity, "invalid input", validationError)
		return
	}

	if err = handle.AddressService.Update(&address, r.Header.Get("token")); err != nil {
		handle.Logger.Error("fail to update address", zap.Error(err))
		util.Response(w).Json(http.StatusInternalServerError, "server error", err)
		return
	}
	util.Response(w).Json(http.StatusOK, "address successfully updated", address)
}

func (handle AddressHandler) Default(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		util.Response(w).Json(http.StatusBadRequest, "invalid address id")
		return
	}

	if err = handle.AddressService.SetDefault(id, r.Header.Get("token")); err != nil {
		handle.Logger.Error("fail to update address", zap.Error(err))
	}
}

func (handle AddressHandler) Delete(w http.ResponseWriter, r *http.Request) {
	addressId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		util.Response(w).Json(http.StatusBadRequest, "invalid address id")
		return
	}

	if err = handle.AddressService.Delete(addressId, r.Header.Get("token")); err != nil {
		handle.Logger.Error("fail to delete address", zap.Error(err))
	}
}
