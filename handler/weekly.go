package handler

import (
	"go.uber.org/zap"
	"net/http"
	"project/service"
	"project/util"
)

type WeeklyHandler struct {
	WeeklyService *service.WeeklyService
	Logger        *zap.Logger
}

func InitWeeklyHandler(service service.Service, log *zap.Logger) WeeklyHandler {
	return WeeklyHandler{WeeklyService: service.Weekly, Logger: log}
}

func (handle WeeklyHandler) Get(w http.ResponseWriter, r *http.Request) {
	weeklies, _ := handle.WeeklyService.Get()
	util.Response(w).Json(http.StatusOK, "Get all weekly promos", weeklies)
}
