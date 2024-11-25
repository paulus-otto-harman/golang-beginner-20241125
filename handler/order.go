package handler

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"project/domain"
	"project/service"
	"project/util"
)

type OrderHandler struct {
	OrderService *service.OrderService
	Logger       *zap.Logger
	Validator    util.Validation
}

func InitOrderHandler(service service.Service, log *zap.Logger, validator util.Validation) OrderHandler {
	return OrderHandler{OrderService: service.Order, Logger: log, Validator: validator}
}

func (handle OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	order := domain.Order{CustomerToken: r.Header.Get("token")}
	json.NewDecoder(r.Body).Decode(&order)

	if err := handle.Validator.ValidateStruct(order); err != nil {
		util.Response(w).ValidationFail(err)
		return
	}

	handle.OrderService.Create(&order)

	util.Response(w).Json(http.StatusCreated, "order created", struct {
		Id        string `json:"id"`
		CreatedAt string `json:"created_at"`
	}{
		Id:        fmt.Sprintf("%d", order.Id),
		CreatedAt: order.CreatedAt,
	})
}
