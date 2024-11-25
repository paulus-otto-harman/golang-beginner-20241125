package handler

import (
	"go.uber.org/zap"
	"net/http"
	"project/service"
	"project/util"
)

type BannerHandler struct {
	BannerService *service.BannerService
	Logger        *zap.Logger
}

func InitBannerHandler(service service.Service, log *zap.Logger) BannerHandler {
	return BannerHandler{BannerService: service.Banner, Logger: log}
}

func (handle BannerHandler) Get(w http.ResponseWriter, r *http.Request) {
	banners, _ := handle.BannerService.Get()
	util.Response(w).Json(http.StatusOK, "Get all banners", banners)
}
