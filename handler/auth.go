package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"project/domain"
	"project/service"
	"project/util"
)

type AuthHandler struct {
	AuthService service.AuthServiceInterface
	Validator   util.Validation
	Logger      *zap.Logger
}

func InitAuthHandler(service service.Service, log *zap.Logger, validator util.Validation) AuthHandler {
	return AuthHandler{AuthService: service.Auth, Logger: log, Validator: validator}
}

func (handle *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	user := domain.User{}
	json.NewDecoder(r.Body).Decode(&user)

	if err := handle.Validator.ValidateStruct(user); err != nil {
		util.Response(w).ValidationFail(err)
		return
	}

	result, err := handle.AuthService.Login(user)
	if err != nil {
		util.Response(w).Json(http.StatusUnauthorized, "invalid username and/or password.")
		return
	}

	util.Response(w).Json(http.StatusOK, "user authenticated", result)
}

func (handle *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")

	err := handle.AuthService.Logout(token)
	if err != nil {
		util.Response(w).Json(http.StatusNotFound, err.Error())
		return
	}
	util.Response(w).Json(http.StatusOK, "User logged out successfully.")
}
