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

type CartHandler struct {
	CartService *service.CartService
	Logger      *zap.Logger
	Validator   util.Validation
}

func InitCartHandler(service service.Service, log *zap.Logger) CartHandler {
	return CartHandler{CartService: service.Cart, Logger: log}
}

func (handle CartHandler) Get(w http.ResponseWriter, r *http.Request) {
	cart, _ := handle.CartService.Get(r.Header.Get("token"))
	util.Response(w).Json(http.StatusOK, "Get customer cart", cart)
}

func (handle CartHandler) Create(w http.ResponseWriter, r *http.Request) {
	cartItem := domain.CartItem{}

	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		handle.Logger.Error("invalid JSON input", zap.Error(err))
		util.Response(w).Json(http.StatusUnprocessableEntity, "invalid input")
		return
	}

	if err := handle.CartService.Store(cartItem, r.Header.Get("token")); err != nil {
		handle.Logger.Error("failed to store cart item", zap.Error(err))
		util.Response(w).Json(http.StatusInternalServerError, "failed to add product to cart")
		return
	}
	util.Response(w).Json(http.StatusOK, "Product successfully added to cart")
}

func (handle CartHandler) Update(w http.ResponseWriter, r *http.Request) {
	cartItem := domain.CartItem{}

	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		handle.Logger.Error("invalid JSON input", zap.Error(err))
		util.Response(w).Json(http.StatusUnprocessableEntity, "invalid input")
		return
	}
	if err := handle.CartService.Update(cartItem, r.Header.Get("token")); err != nil {
		handle.Logger.Error("failed to update cart item", zap.Error(err))
		util.Response(w).Json(http.StatusInternalServerError, "failed to update cart item")
		return
	}
	util.Response(w).Json(http.StatusOK, "Cart successfully updated")
}

func (handle CartHandler) Delete(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(chi.URLParam(r, "product_id"))
	if err != nil {
		util.Response(w).Json(http.StatusBadRequest, "Invalid Product ID")
		return
	}
	if err := handle.CartService.Delete(productId, r.Header.Get("token")); err != nil {
		handle.Logger.Error("failed to delete cart item", zap.Error(err))
		util.Response(w).Json(http.StatusInternalServerError, "failed to delete cart item")
		return
	}
	util.Response(w).Json(http.StatusOK, "Product successfully removed from cart")
}
