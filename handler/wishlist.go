package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"project/domain"
	"project/service"
	"project/util"
)

type WishlistHandler struct {
	WishlistService *service.WishlistService
	Validator       util.Validation
	Logger          *zap.Logger
}

func InitWishlistHandler(service service.Service, log *zap.Logger, validator util.Validation) WishlistHandler {
	return WishlistHandler{WishlistService: service.Wishlist, Logger: log, Validator: validator}
}

func (handle *WishlistHandler) Store(w http.ResponseWriter, r *http.Request) {
	wishlistItem := domain.WishlistItem{}

	if err := json.NewDecoder(r.Body).Decode(&wishlistItem); err != nil {
		handle.Logger.Error("invalid JSON input", zap.Error(err))
		util.Response(w).Json(http.StatusUnprocessableEntity, "invalid input")
		return
	}

	if validationError := handle.Validator.ValidateStruct(wishlistItem); validationError != nil {
		util.Response(w).Json(http.StatusUnprocessableEntity, "invalid product")
		return
	}

	if err := handle.WishlistService.Store(&wishlistItem, r.Header.Get("token")); err != nil {
		handle.Logger.Error("Add To Wishlist failed", zap.Error(err))
	}

	util.Response(w).Json(http.StatusCreated, "Added to Wishlist", wishlistItem)
}

func (handle *WishlistHandler) Destroy(w http.ResponseWriter, r *http.Request) {
	wishlistItem := domain.WishlistItem{}

	if err := json.NewDecoder(r.Body).Decode(&wishlistItem); err != nil {
		handle.Logger.Error("invalid JSON input", zap.Error(err))
		util.Response(w).Json(http.StatusUnprocessableEntity, "invalid input")
		return
	}

	if validationError := handle.Validator.ValidateStruct(wishlistItem); validationError != nil {
		util.Response(w).Json(http.StatusUnprocessableEntity, "invalid product")
		return
	}

	if err := handle.WishlistService.Destroy(&wishlistItem, r.Header.Get("token")); err != nil {
		handle.Logger.Error("Failed removing item from Wishlist", zap.Error(err))
	}

	util.Response(w).Json(http.StatusOK, "Item removed from Wishlist")
}
