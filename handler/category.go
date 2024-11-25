package handler

import (
	"go.uber.org/zap"
	"net/http"
	"project/service"
	"project/util"
)

type CategoryHandler struct {
	CategoryService *service.CategoryService
	Logger          *zap.Logger
}

func InitCategoryHandler(service service.Service, log *zap.Logger) CategoryHandler {
	return CategoryHandler{CategoryService: service.Category, Logger: log}
}

func (handle CategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	categories, _ := handle.CategoryService.Get()
	util.Response(w).Json(http.StatusOK, "Get all categories", categories)
}
